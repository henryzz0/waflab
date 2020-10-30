package rule

import "testing"

func TestParseRuleExample(t *testing.T) {
	parseRuleExample()
}

func TestParseRuleText(t *testing.T) {
	s := `
SecRule TX:EXECUTING_PARANOIA_LEVEL "@lt 1" "id:920011,phase:1,pass,nolog,skipAfter:END-REQUEST-920-PROTOCOL-ENFORCEMENT"
SecRule TX:EXECUTING_PARANOIA_LEVEL "@lt 1" "id:920012,phase:2,pass,nolog,skipAfter:END-REQUEST-920-PROTOCOL-ENFORCEMENT"
`
	parseRuleText(s)
}
