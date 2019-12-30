package rule

import (
	"errors"
	"regexp"

	"github.com/waflab/waflab/util"
)

const (
	RuleNormal = "normal"
	RuleControl = "control"
)

var reRuleId *regexp.Regexp
var reParanoiaLevel *regexp.Regexp

func init() {
	var err error
	reRuleId, err = regexp.Compile("id:(\\d+),")
	if err != nil {
		panic(err)
	}

	reParanoiaLevel, err = regexp.Compile("TX:EXECUTING_PARANOIA_LEVEL \"@lt (\\d)\"")
	if err != nil {
		panic(err)
	}
}

type Rule struct {
	No            int    `json:"no"`
	Id            string `json:"id"`
	Typ           string `json:"type"`
	ParanoiaLevel int    `json:"paranoiaLevel"`
	Text          string `json:"text"`

	ChainRules []*Rule `json:"chainRules"`
}

func newRule(no int, text string) *Rule {
	r := Rule{}
	r.No = no
	r.Text = text
	r.parseText()
	return &r
}

func newChainRule(no int, text string) *Rule {
	r := Rule{}
	r.No = no
	r.Text = text
	return &r
}

func parseRuleId(text string) string {
	res := reRuleId.FindStringSubmatch(text)
	if res == nil {
		panic(errors.New("parseRuleId() error: cannot find rule Id in rule text: " + text))
	}
	return res[1]
}

func parseParanoiaLevel(text string) int {
	res := reParanoiaLevel.FindStringSubmatch(text)
	if res == nil {
		return -1
	}
	return util.ParseInt(res[1])
}

func (r *Rule) parseText() {
	r.Id = parseRuleId(r.Text)
	pl := parseParanoiaLevel(r.Text)
	if pl != -1 {
		r.Typ = RuleControl
		r.ParanoiaLevel = pl
	} else {
		r.Typ = RuleNormal
		r.ParanoiaLevel = -1
	}
}

func (r *Rule) addChainRule(text string) {
	r.ChainRules = append(r.ChainRules, newChainRule(len(r.ChainRules), text))
}
