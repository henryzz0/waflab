package rule

type Ruleset struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Count   int    `json:"count"`

	Rulefiles []*Rulefile `json:"rulefiles"`
}

func newRuleset(id string) *Ruleset {
	rs := Ruleset{}
	rs.Id = id
	return &rs
}
