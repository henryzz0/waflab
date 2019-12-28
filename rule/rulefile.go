package rule

import "strings"

type Rulefile struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	No     string `json:"no"`
	Suffix string `json:"suffix"`
	Count  int    `json:"count"`

	Rules []*Rule `json:"rules"`
}

func newRulefile(id string) *Rulefile {
	rf := Rulefile{}
	rf.Id = id
	rf.parseId()
	return &rf
}

func (rf *Rulefile) parseId() {
	tokens := strings.SplitN(rf.Id, "-", 3)
	rf.Type = tokens[0]
	rf.No = tokens[1]
	rf.Suffix = tokens[2]
}
