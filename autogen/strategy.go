package autogen

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"

	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/payload"
	"github.com/waflab/waflab/autogen/transformer"
	"github.com/waflab/waflab/autogen/yaml"
	"github.com/waflab/waflab/rule"
	"github.com/waflab/waflab/test"
)

func GenerateTests(ruleString string) (YAMLs []*test.Testfile) {
	rules, err := rule.ParseRuleDataToList(ruleString)
	if err != nil {
		log.Printf("Err while parsing rule string %s\n", ruleString)
		return nil
	}

	for _, rule := range rules {
		if rule.Actions.Chain { // Chained rule
			log.Printf("Err chain rule %d not supported\n", rule.Actions.Id)
			return nil
		}
		if t := processIndependentRule(rule); t != nil {
			YAMLs = append(YAMLs, processIndependentRule(rule))
		}
	}
	return YAMLs
}

func processIndependentRule(rule *parser.RuleDirective) *test.Testfile {
	v := yaml.DefaultYAML()

	// set meta information
	v.Meta.Author = "Microsoft Research Asia"
	v.Meta.Name = fmt.Sprintf("%d.yaml", rule.Actions.Id)
	v.Meta.Description = "This YAML file is automatically generated using AutoGen"

	// reverse generate operator and transformation
	reversed, err := operator.ReverseOperator(rule.Operator)
	if err != nil {
		log.Printf("Rule %d: skip generation, %v\n", rule.Actions.Id, err)
		return nil
	}
	reversed = transformer.ReverseTransform(rule.Actions.Trans, reversed)

	// set status code
	var statusCode int
	for _, action := range rule.Actions.Action {
		switch action.Tk {
		case parser.TkActionAllow, parser.TkActionPass:
			statusCode = 200
		case parser.TkActionDeny, parser.TkActionBlock:
			statusCode = 403
		default:
		}
	}

	// process variable index exclusion
	newVariables, err := processIndexExclusion(rule.Variable)
	if err != nil {
		log.Printf("Rule %d: skip, %v", rule.Actions.Id, err)
		return nil
	}
	rule.Variable = newVariables

	current := 0 // number of test case generated
	for _, variable := range rule.Variable {
		// expand the tests slice if necessary
		if current >= len(v.Tests) {
			v.Tests = append(v.Tests, &test.Test{
				Stages: []*test.StageWrapper{
					{
						Stage: yaml.DefaultStage(),
					},
				},
			})
		}
		// add variable
		err = payload.AddVariable(variable, reversed, v.Tests[current].Stages[0].Stage.Input)
		if err != nil {
			log.Printf("Rule %d: skip %s, %v\n",
				rule.Actions.Id,
				parser.VariableNameMap[variable.Tk],
				err)
			continue
		}
		// add title and status
		v.Tests[current].TestTitle = fmt.Sprintf("%s-%s",
			strconv.Itoa(rule.Actions.Id),
			strconv.Itoa(current+1))
		v.Tests[current].Stages[0].Stage.Output.Status = []int{statusCode}

		current++
	}

	// it is possible that all variable is not supported in a rule,
	// we therefore skip it entirely
	if current == 0 {
		log.Printf("Rule %d: no avaliable variable, skip entirely", rule.Actions.Id)
		return nil
	}

	return v
}
