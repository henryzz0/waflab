package autogen

import (
	"fmt"
	"strconv"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"

	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/payload"
	"github.com/waflab/waflab/autogen/transformer"
	"github.com/waflab/waflab/autogen/yaml"
	"github.com/waflab/waflab/rule"
	"github.com/waflab/waflab/test"
)

func ProcessIndependentRule(ruleString string) (YAMLs []*test.Testfile) {
	rules := rule.ParseRuleDataToList(ruleString)
	for _, rule := range rules {
		v := yaml.DefaultYAML()
		reversed, _ := operator.ReverseOperator(rule.Operator)
		reversed = transformer.ReverseTransform(rule.Actions.Trans, reversed)

		var statusCode int
		for _, action := range rule.Actions.Action {
			switch action.Tk {
			case parser.TkActionAllow, parser.TkActionPass:
				statusCode = 200
			case parser.TkActionDeny:
				statusCode = 403
			default:
			}
		}

		// process variable index exclusion
		newVariables, err := processIndexExclusion(rule.Variable)
		if err != nil {
			panic(err)
		}
		rule.Variable = newVariables

		for i, variable := range rule.Variable {
			if i >= len(v.Tests) {
				v.Tests = append(v.Tests, &test.Test{
					Stages: []*test.StageWrapper{
						{
							Stage: yaml.DefaultStage(),
						},
					},
				})
			}
			v.Tests[i].TestTitle = fmt.Sprintf("%s-%s", strconv.Itoa(rule.Actions.Id), strconv.Itoa(i))
			payload.AddVariable(variable, reversed, v.Tests[i].Stages[0].Stage.Input)
			v.Tests[i].Stages[0].Stage.Output.Status = []int{statusCode}
		}

		YAMLs = append(YAMLs, v)
	}

	return YAMLs
}
