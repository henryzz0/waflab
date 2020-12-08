package rule

import (
	"regexp"
	"strings"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

var reComment *regexp.Regexp
var reOtherDirective *regexp.Regexp
var reOtherDirective2 *regexp.Regexp

func init() {
	reComment, _ = regexp.Compile("(?:^|\n)#.*")
	reOtherDirective, _ = regexp.Compile("(SecMarker|SecComponentSignature).*")
	reOtherDirective2, _ = regexp.Compile("SecAction(?U:[\\s\\S]*)\n\n")
}

func removeComment(s string) string {
	s = reComment.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	s = reOtherDirective.ReplaceAllString(s, "")

	s = reOtherDirective2.ReplaceAllString(s, "")
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

func ParseRuleDataToList(text string) []*parser.RuleDirective {
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
