package rule

import (
	"regexp"
	"strings"

	"github.com/waflab/waflab/util"
)

func removeComment(s string) string {
	re, _ := regexp.Compile("#.*\n")
	return re.ReplaceAllString(s, "")
}

func parseRules(rf *Rulefile, text string) {
	text = removeComment(text)
	text = strings.ReplaceAll(text, "\\\n", "")

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "SecMarker") || strings.HasPrefix(line, "SecComponentSignature") {
			continue
		}

		if strings.HasPrefix(line, "SecRule") {
			rf.Rules = append(rf.Rules, newRule(len(rf.Rules), line))
		} else if strings.HasPrefix(line, "    SecRule") {
			line = strings.Trim(line, " ")
			rf.Rules[len(rf.Rules) - 1].addChainRule(line)
		}
	}
}

func printRules(rf *Rulefile) {
	for _, rule := range rf.Rules {
		println(rule.Text)
	}
}

func parse() {
	rf := newRulefile(0, "REQUEST-920-PROTOCOL-ENFORCEMENT")
	text := util.ReadStringFromPath(util.CrsRuleDir + "REQUEST-920-PROTOCOL-ENFORCEMENT.conf")
	parseRules(rf, text)
	printRules(rf)

	//scaner := parser.NewSecLangScannerFromString(text)
	//d, err := scaner.AllDirective()
	//if err != nil {
	//	panic(err)
	//}
	//utils.Pprint(d)
}