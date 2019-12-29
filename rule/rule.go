package rule

import (
	"errors"
	"regexp"
)

var reRuleId *regexp.Regexp

func init() {
	var err error
	reRuleId, err = regexp.Compile("\"id:(\\d+),")
	if err != nil {
		panic(err)
	}
}

type Rule struct {
	No   int    `json:"no"`
	Id   string `json:"id"`
	Text string `json:"text"`

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

func (r *Rule) parseText() {
	res := reRuleId.FindStringSubmatch(r.Text)
	if res == nil {
		panic(errors.New("parseText() error: cannot find rule Id in rule text: " + r.Text))
	}
	r.Id = res[1]
}

func (r *Rule) addChainRule(text string) {
	r.ChainRules = append(r.ChainRules, newChainRule(len(r.ChainRules), text))
}
