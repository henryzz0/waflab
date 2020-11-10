package rule

import (
	"fmt"
	"strings"

	"github.com/waflab/waflab/object"
	"github.com/waflab/waflab/test"
	"github.com/waflab/waflab/util"
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
			rf.Pl1TestCount += r.RegressionTestCount
		} else if r.ParanoiaLevel == 2 {
			rf.Pl2Count += 1
			rf.Pl2TestCount += r.RegressionTestCount

		} else if r.ParanoiaLevel == 3 {
			rf.Pl3Count += 1
			rf.Pl3TestCount += r.RegressionTestCount

		} else if r.ParanoiaLevel == 4 {
			rf.Pl4Count += 1
			rf.Pl4TestCount += r.RegressionTestCount
		} else {
			println(r.Id)
		}
	}
	rf.Count = len(rf.Rules)
	rf.TestCount = rf.Pl1TestCount + rf.Pl2TestCount + rf.Pl3TestCount + rf.Pl4TestCount
}

func getUserAgent(tf *test.Testfile) string {
	headers := tf.Tests[0].Stages[0].Stage.Input.Headers
	if userAgent, ok := headers["User-Agent"]; ok {
		return userAgent
	} else {
		return ""
	}
}

func getStatus(tf *test.Testfile) int {
	status := tf.Tests[0].Stages[0].Stage.Output.Status
	if len(status) > 0 {
		return status[0]
	} else {
		return -1
	}
}

func syncTestfile(tf *test.Testfile, text string) {
	testcase := object.Testcase{
		Name:        tf.Meta.Name,
		CreatedTime: util.GetCurrentTime(),
		Desc:        tf.Meta.Description,
		Author:      tf.Meta.Author,
		Enabled:     tf.Meta.Enabled,
		TestCount:   len(tf.Tests),
		Method:      tf.Tests[0].Stages[0].Stage.Input.Method,
		UserAgent:   getUserAgent(tf),
		Status:      getStatus(tf),
		Data:        tf,
		RawData:     text,
	}

	if object.GetTestcase(testcase.Name) != nil {
		object.DeleteTestcase(&testcase)
	}
	object.AddTestcase(&testcase)

	testset := object.GetTestset("regression-test")
	if !util.StringListContains(testset.Testcases, tf.Meta.Name) {
		testset.Testcases = append(testset.Testcases, tf.Meta.Name)
	}
	object.UpdateTestset(testset.Name, testset)
}

func (rf *Rulefile) loadTestsets() {
	for _, r := range rf.Rules {
		// util.CrsTestDir + "REQUEST-920-PROTOCOL-ENFORCEMENT/920100.yaml"
		path := fmt.Sprintf("%s%s/%s.yaml", util.CrsTestDir, rf.Id, r.Id)
		if !util.FileExist(path) {
			continue
		}

		text := util.ReadStringFromPath(path)
		tf := test.LoadTestfileFromString(text)
		//syncTestfile(tf, text)

		r.RegressionTestCount = len(tf.Tests)
	}
}
