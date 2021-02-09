package autogen

import "testing"

func TestProcessIndependentRule(t *testing.T) {
	ruleString := `SecRule ARGS '@rx (?i)<script[^>]*>' \
                        "id:123,\
                         phase:2,\
                         t:lowercase,\
                         deny"`
	testfiles := GenerateTests(ruleString, 10)
	for _, testfile := range testfiles {
		println(testfile)
	}
}
