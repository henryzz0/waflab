package rule

import "github.com/senghoo/modsecurity-go/seclang/parser"

func parseRules2(rf *Rulefile, text string) {
	text = removeComment(text)

	scaner := parser.NewSecLangScannerFromString(text)
	directives, err := scaner.AllDirective()
	if err != nil {
		panic(err)
	}

	for _, directive := range directives {
		//var res parser.Directive
		res := directive
		d, ok := (res).(*parser.RuleDirective)
		if ok {
			println(d)
		}

		print(d)
	}
}
