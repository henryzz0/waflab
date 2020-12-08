package autogen_test

import (
	"fmt"
	"github.com/waflab/waflab/autogen"

	y "gopkg.in/yaml.v2"
)

func ExampleProcessIndepdentRule() {
	ruleString := `SecRule ARGS '@rx (?i)<script[^>]*>' \
                        "id:123,\
                         phase:2,\
                         t:lowercase,\
                         deny"`
	v := autogen.ProcessIndependentRule(ruleString)
	d, _ := y.Marshal(v)
	fmt.Printf("%s\n", string(d))
	//Output:
	//TODO
}