package rule

import (
	"fmt"

	"github.com/waflab/waflab/util"
)

func ReadRuleset(id string) *Ruleset {
	fmt.Printf("Read ruleset for Id: [%s].\n", id)

	rs := newRuleset(id)
	if rs.Id == "crs-3.2" {
		rs.Name = "CoreRuleSet"
		rs.Version = "v3.2/master"
	}

	filenames := util.ListFileIds(util.CrsRuleDir)
	for _, filename := range filenames {
		rf := ReadRulefile(len(rs.Rulefiles), filename)
		rs.Rulefiles = append(rs.Rulefiles, rf)
		rs.RulefileMap[filename] = rf
		rs.RuleCount += rf.Count
	}
	rs.FileCount = len(rs.Rulefiles)

	return rs
}

func ReadRulefile(no int, id string) *Rulefile {
	fmt.Printf("Read rulefile for Id: [%s].\n", id)

	rf := newRulefile(no, id)

	text := util.ReadStringFromPath(util.CrsRuleDir + id + ".conf")
	parseRules2(rf, text)
	rf.loadTestsets()
	rf.syncPls()

	return rf
}
