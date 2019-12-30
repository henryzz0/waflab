package rule

import "strings"

type Rulefile struct {
	No      int    `json:"no"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Count   int    `json:"count"`
	PlCount int    `json:"plCount"`

	Rules []*Rule `json:"rules"`
}

func newRulefile(no int, id string) *Rulefile {
	rf := Rulefile{}
	rf.No = no
	rf.Id = id
	rf.parseId()
	return &rf
}

func (rf *Rulefile) parseId() {
	tokens := strings.SplitN(rf.Id, "-", 3)
	rf.Type = tokens[0]
	rf.Name = tokens[1]
	rf.Desc = tokens[2]
}

func (rf *Rulefile) syncPls() {
	pl := -1
	for _, r := range rf.Rules {
		if r.Typ == RuleControl {
			pl = r.ParanoiaLevel
			r.ParanoiaLevel = -1
		} else if r.Typ == RuleNormal {
			r.ParanoiaLevel = pl
		}
	}

	newRules := []*Rule{}
	for _, r := range rf.Rules {
		if r.Typ == RuleNormal {
			newRules = append(newRules, r)
		}
	}
	rf.Rules = newRules
}
