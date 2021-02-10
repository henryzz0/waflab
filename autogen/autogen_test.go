package autogen

import (
	"fmt"
	"testing"

	"github.com/waflab/waflab/object"
	"github.com/waflab/waflab/rule"
	"gopkg.in/yaml.v2"
)

func TestSyncAutoGen(t *testing.T) {
	object.InitOrmManager()

	ruleset := rule.ReadRuleset("crs-3.2")
	for _, rulefile := range ruleset.Rulefiles {
		rules := rulefile.Rules
		for _, r := range rules {
			testfiles := GenerateTests(r.Text, 1)
			// files like REQUEST-901-INITIALIZATION.conf will be nil
			if testfiles == nil {
				fmt.Printf("[%s] [%s] testfiles == nil, Rule.Text: %s\n", rulefile.Id, r.Id, r.Text)
				continue
			}

			if len(testfiles) != 1 {
				panic(fmt.Sprintf("TestSyncAutoGen(): len(testfiles): %d != 1", len(testfiles)))
			}

			singleTestFile := testfiles[0]
			content, err := yaml.Marshal(singleTestFile)
			if err != nil {
				panic(err)
			}
			rule.SyncTestfile(singleTestFile, string(content), "autogen-test")
		}
	}
}
