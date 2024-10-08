package base

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	// crashes identified by go-fuzz
	origs := []string{
		"&\x020000",
		"&\x020000\x9c",
		"&\x020\x9c0",
		"&\x0230j",
		"&\x02\x98000",
		"&\x02\x983\xc8j00",
		"00\x000",
		"00\x0000",
		"00\x0000000000000",
		"\x11\x030",
	}

	for _, orig := range origs {
		escaped := escape(orig)
		unescaped, err := unescape(escaped)
		if err != nil {
			t.Errorf("%s: orig: %#v, escaped: %#v, unescaped: %#v\n", err.Error(), orig, escaped, unescaped)
			continue
		}

		if unescaped != orig {
			t.Errorf("expected: %#v, got: %#v", orig, unescaped)
		}
	}
}

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		escaped := escape(orig)

		unescaped, err := unescape(escaped)
		if err != nil {
			t.Errorf("Can't unescape escaped string %q", escaped)
		}

		if unescaped != orig {
			t.Errorf("Before: %q, after: %q", orig, unescaped)
		}
	})
}
