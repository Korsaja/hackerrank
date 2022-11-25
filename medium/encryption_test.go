package medium

import (
	"fmt"
	"testing"
)

func TestEncryption(t *testing.T) {
	var testCases = []struct {
		plainText  string
		cipherText string
	}{
		{plainText: "haveaniceday", cipherText: "hae and via ecy"},
		{plainText: "feedthedog", cipherText: "fto ehg ee dd"},
		{plainText: "chillout", cipherText: "clu hlt io"},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test case [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			cipherText := encryption(test.plainText)
			if cipherText != test.cipherText {
				t.Errorf("invalid ecnryption ex: %s got: %s\n", test.cipherText, cipherText)
			}
		})
	}
}
