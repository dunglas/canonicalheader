package qualified

import (
	"net/http"

	"qualified/keys"
)

const suffix = "test-header"

func _() {
	h := http.Header{}

	h.Get(keys.NoCanonical)         // want `use "Tt" instead of "TT"`
	h.Get(string(keys.NoCanonical)) // want `use "Tt" instead of "TT"`
	h.Get(keys.Canonical)
	h.Get(keys.Prefix + suffix) // want `use "X-Test-Header" instead of "X-test-header"`
	h.Get(("TT"))               // want `use "Tt" instead of "TT"`

	_ = h[keys.NoCanonical]       // want `use "Tt" instead of "TT"`
	delete(h, keys.Prefix+suffix) // want `use "X-Test-Header" instead of "X-test-header"`
	_ = http.Header{
		keys.NoCanonical: nil, // want `use "Tt" instead of "TT"`
	}
}
