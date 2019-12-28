package rule

type Ruleset struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Count   int    `json:"count"`

	Rulefiles   []*Rulefile          `json:"rulefiles"`
	RulefileMap map[string]*Rulefile `json:"-"`
}

func newRuleset(id string) *Ruleset {
	rs := Ruleset{}
	rs.Id = id
	rs.RulefileMap = map[string]*Rulefile{}
	return &rs
}
