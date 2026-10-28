package canonicalheader

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

type constantString struct {
	originalValue,
	nameOfConst string

	pos token.Pos
	end token.Pos
}

func newConstantKey(info *types.Info, ident *ast.Ident) (constantString, error) {
	c, ok := info.ObjectOf(ident).(*types.Const)
	if !ok {
		return constantString{}, fmt.Errorf("type %T is not support", c)
	}

	// string(r) with a rune constant r carries an integer value.
	if c.Val().Kind() != constant.String {
		return constantString{}, fmt.Errorf("constant %s is not a string", c.Name())
	}

	return constantString{
		nameOfConst:   c.Name(),
		originalValue: constant.StringVal(c.Val()),
		pos:           ident.Pos(),
		end:           ident.End(),
	}, nil
}

func (c constantString) diagnostic(canonicalHeader string) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos: c.pos,
		End: c.end,
		Message: fmt.Sprintf(
			"use %q instead of %q",
			canonicalHeader,
			c.originalValue,
		),
	}
}

func (c constantString) value() string {
	return c.originalValue
}
