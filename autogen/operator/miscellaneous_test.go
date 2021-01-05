package operator

import (
	"strings"
	"testing"
)

func TestReverseIPMatchSimple(t *testing.T) {
	argument := "127.0.0.1,::1" // from CRS:9005100

	res, err := reverseIPMatch(argument, false)
	if err != nil {
		panic(err)
	}
	isMatched := true
	for _, str := range strings.Split(argument, ",") {
		if str == res {
			isMatched = true
		}
	}
	if !isMatched {
		t.Errorf("Mismatch: %s not in the %s", res, argument)
	}
}
