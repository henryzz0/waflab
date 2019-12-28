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

func parseRules() []string {
	text := util.ReadStringFromPath(util.CrsRuleDir + "REQUEST-920-PROTOCOL-ENFORCEMENT.conf")
	text = removeComment(text)
	text = strings.ReplaceAll(text, "\\\n", "")

	tokens := strings.Split(text, "\n")

	res := []string{}
	for _, token := range tokens {
		if token == "" {
			continue
		}

		if strings.HasPrefix(token, "SecMarker") {
			continue
		}

		token = strings.Trim(token, " ")
		res = append(res, token)
	}

	//text = strings.ReplaceAll(text, "\n\n", "\n")
	return res
}

func printRules(rules []string) {
	for _, rule := range rules {
		println(rule)
	}
}

func parse() {
	rules := parseRules()
	printRules(rules)

	//scaner := parser.NewSecLangScannerFromString(text)
	//d, err := scaner.AllDirective()
	//if err != nil {
	//	panic(err)
	//}
	//utils.Pprint(d)
}