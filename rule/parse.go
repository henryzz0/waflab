package rule

import (
	"regexp"
	"strings"

	"github.com/waflab/waflab/util"
)

func removeComment(s string) string {
	re, _ := regexp.Compile("(?:^|\n)#.*")
	s = re.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	re, _ = regexp.Compile("(SecMarker|SecComponentSignature).*")
	s = re.ReplaceAllString(s, "")

	re, _ = regexp.Compile("SecAction(?U:[\\s\\S]*)\n\n")
	s = re.ReplaceAllString(s, "")
	return s
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

func parseRuleExample() {
	rf := newRulefile(0, "REQUEST-920-PROTOCOL-ENFORCEMENT")
	text := util.ReadStringFromPath(util.CrsRuleDir + "REQUEST-920-PROTOCOL-ENFORCEMENT.conf")

	//parseRules(rf, text)
	parseRules2(rf, text)

	rf.syncPls()
	printRules(rf)

	//scaner := parser.NewSecLangScannerFromString(text)
	//d, err := scaner.AllDirective()
	//if err != nil {
	//	panic(err)
	//}
	//utils.Pprint(d)
}

func parseRuleText(text string) {
	parseRules2(nil, text)
}
