package naughtystrings

import "testing"

func TestNaughtyStrings(t *testing.T) {
	for _, test := range [][]string{Base64Encoded(), Unencoded()} {
		t.Logf("Test: %v", test)

		if n := len(test); n == 0 {
			t.Error("Length: 0")
		}
	}
}
