package rule

import (
	"regexp"
	"strings"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
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

func parseRulesToLines(text string) []string {
	text = strings.ReplaceAll(text, "\\\n", "")
	lines := strings.Split(text, "\n")
	res := []string{}
	for _, line := range lines {
		if line != "" {
			res = append(res, line)
		}
	}
	return res
}

func parseRuleDataToList(text string) []*parser.RuleDirective {
	scaner := parser.NewSecLangScannerFromString(text)
	directives, err := scaner.AllDirective()
	if err != nil {
		panic(err)
	}

	res := []*parser.RuleDirective{}
	for _, directive := range directives {
		//var res parser.Directive
		d := directive
		rd, ok := (d).(*parser.RuleDirective)
		if ok {
			//println(d)
		}

		res = append(res, rd)
	}
	return res
}
