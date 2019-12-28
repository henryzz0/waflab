package rule

type Rule struct {
	No   int    `json:"no"`
	Text string `json:"text"`
}

func newRule(no int, text string) *Rule {
	r := Rule{}
	r.No = no
	r.Text = text
	return &r
}
