package medium

import (
	"fmt"
	"testing"
)

func TestExtraLongFactorials(t *testing.T) {
	var testCases = []struct {
		n   int32
		out string
	}{
		{n: 30, out: "265252859812191058636308480000000"},
		{n: 25, out: "15511210043330985984000000"},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test case [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			fact := extraLongFactorials(test.n)
			if fact != test.out {
				t.Errorf("invalid ecnryption ex: %s got: %s\n", test.out, fact)
			}
		})
	}
}
