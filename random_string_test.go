package randomstring

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		charset  string
		generate func(n int) string
	}{
		{"Alpha", Alpha, RandomAlpha},
		{"AlphaNumber", AlphaNumber, RandomAlphaNumber},
		{"AlphaLower", AlphaLower, RandomAlphaLower},
		{"AlphaUpper", AlphaUpper, RandomAlphaUpper},
		{"Number", Number, RandomNumber},
	} {
		l := 16
		s1 := testCase.generate(l)
		s2 := testCase.generate(l)
		if s1 == s2 {
			t.Errorf(
				"%s random string generator should not generate two same string '%s'",
				testCase.name, s1,
			)
		}
		if len(s1) != l || len(s2) != l {
			t.Errorf(
				"%s random string generator should generate string with length %d, but got (%d, %d)",
				testCase.name, l, len(s1), len(s2),
			)
		}
		var invalidChars []rune
		for _, c := range s1 + s2 {
			if !strings.ContainsRune(testCase.charset, c) {
				invalidChars = append(invalidChars, c)
			}
		}
		if len(invalidChars) > 0 {
			t.Errorf(
				"%s random string generator should generate characters only in '%s', but got invalid ones '%s'",
				testCase.name, testCase.charset, string(invalidChars),
			)
		}
	}
}
