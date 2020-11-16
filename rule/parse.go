package rule

import (
	"errors"
)

func parseRules(rf *Rulefile, text string) {
	text = removeComment(text)

	lines := parseRulesToLines(text)
	ruleDataList := parseRuleDataToList(text)
	if len(lines) != len(ruleDataList) {
		panic(errors.New("parseRules() error: len(lines) != len(ruleDataList)"))
	}

	for i, ruleData := range ruleDataList {
		if i > 0 && ruleDataList[i-1].Actions != nil && ruleDataList[i-1].Actions.Chain {
			rf.Rules[len(rf.Rules)-1].addChainRule(lines[i], ruleData)
		} else {
			r := newRule(len(rf.Rules), lines[i], ruleData)
			rf.Rules = append(rf.Rules, r)
		}
	}
}

func printRules(rf *Rulefile) {
	for _, rule := range rf.Rules {
		println(rule.Text)
	}
}
