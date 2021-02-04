package rule

import (
	"strings"
)

func (rf *Rulefile) loadRules(text string) {
	text = FilterSecRule(text)
	ruleDataList, err := ParseRuleDataToList(text)
	if err != nil {
		panic(err)
	}

	// split rule string in to line
	var lines []string
	sep := strings.LastIndex(text, "SecRule")
	for ; sep != -1; sep = strings.LastIndex(text, "SecRule") {
		lines = append(lines, text[sep:])
		text = text[:sep]
	}
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	for i, ruleData := range ruleDataList {
		if ruleData.Actions == nil || ruleData.Actions.Id == 0 { // sub rule of chained rule
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
