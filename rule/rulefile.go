package rule

import (
	"fmt"
	"strings"

	"github.com/waflab/waflab/util"
	"gopkg.in/yaml.v2"
)

type Rulefile struct {
	No           int    `json:"no"`
	Id           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Count        int    `json:"count"`
	Pl1Count     int    `json:"pl1Count"`
	Pl2Count     int    `json:"pl2Count"`
	Pl3Count     int    `json:"pl3Count"`
	Pl4Count     int    `json:"pl4Count"`
	TestCount    int    `json:"testCount"`
	Pl1TestCount int    `json:"pl1TestCount"`
	Pl2TestCount int    `json:"pl2TestCount"`
	Pl3TestCount int    `json:"pl3TestCount"`
	Pl4TestCount int    `json:"pl4TestCount"`

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

	for _, r := range rf.Rules {
		if r.ParanoiaLevel == -1 {
			r.ParanoiaLevel = 1
		}

		if r.ParanoiaLevel == 1 {
			rf.Pl1Count += 1
			rf.Pl1TestCount += r.TestCount
		} else if r.ParanoiaLevel == 2 {
			rf.Pl2Count += 1
			rf.Pl2TestCount += r.TestCount

		} else if r.ParanoiaLevel == 3 {
			rf.Pl3Count += 1
			rf.Pl3TestCount += r.TestCount

		} else if r.ParanoiaLevel == 4 {
			rf.Pl4Count += 1
			rf.Pl4TestCount += r.TestCount
		} else {
			println(r.Id)
		}
	}
	rf.Count = len(rf.Rules)
	rf.TestCount = rf.Pl1TestCount + rf.Pl2TestCount + rf.Pl3TestCount + rf.Pl4TestCount
}

func (rf *Rulefile) loadTestsets() {
	for _, r := range rf.Rules {
		// util.CrsTestDir + "REQUEST-920-PROTOCOL-ENFORCEMENT/920100.yaml"
		path := fmt.Sprintf("%s%s/%s.yaml", util.CrsTestDir, rf.Id, r.Id)
		if !util.FileExist(path) {
			continue
		}
		text := util.ReadStringFromPath(path)

		ts := Testset{}
		err := yaml.Unmarshal([]byte(text), &ts)
		if err != nil {
			panic(err)
		}
		r.Testset = &ts
		r.TestCount = len(ts.Tests)
	}
}
