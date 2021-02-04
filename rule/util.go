package rule

import (
	"regexp"
	"strings"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

var commentRegexp *regexp.Regexp
var changeLineRegexp *regexp.Regexp

func init() {
	commentRegexp, _ = regexp.Compile("(?:^|\n)#.*")
	changeLineRegexp, _ = regexp.Compile(`\n{3,}`)
}

// FilterSecRule remove any non-SecRule Rule Directive (such as SecMarker and SecAction)
// from input text string
func FilterSecRule(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")   // CRLF to LF, sanity check
	text = commentRegexp.ReplaceAllString(text, "") // remove all comments
	text = strings.TrimSpace(text)
	text = changeLineRegexp.ReplaceAllString(text, "\n\n")

	var builder strings.Builder
	lines := strings.Split(text, "\n\n")
	for index, r := range lines {
		r = strings.TrimSpace(r)
		if strings.HasPrefix(r, "SecRule") {
			if index < len(lines)-1 {
				builder.WriteString(r + "\n\n")
			} else {
				builder.WriteString(r)
			}
		}
	}
	return builder.String()
}

func ParseRuleDataToList(text string) ([]*parser.RuleDirective, error) {
	scaner := parser.NewSecLangScannerFromString(text)
	directives, err := scaner.AllDirective()
	if err != nil {
		return nil, err
	}

	res := []*parser.RuleDirective{}
	for _, directive := range directives {
		//var res parser.Directive
		d := directive
		rd, ok := (d).(*parser.RuleDirective)
		if !ok {
			continue
		}

		res = append(res, rd)
	}
	return res, nil
}
