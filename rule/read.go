package rule

import (
	"fmt"

	"github.com/waflab/waflab/util"
)

func ReadRuleset(id string) *Ruleset {
	fmt.Printf("Read ruleset for Id: [%s].\n", id)

	rs := newRuleset(id)
	if rs.Id == "crs-3.3" {
		rs.Name = "CoreRuleSet"
		rs.Version = "v3.3/dev"
	}

	filenames := util.ListFileIds(util.CrsRuleDir)
	for _, filename := range filenames {
		rs.Rulefiles = append(rs.Rulefiles, ReadRulefile(filename))
	}
	rs.Count = len(rs.Rulefiles)

	return rs
}

func ReadRulefile(id string) *Rulefile {
	fmt.Printf("Read rulefile for Id: [%s].\n", id)

	rf := newRulefile(id)

	text := util.ReadStringFromPath(util.CrsRuleDir + id + ".conf")
	rf.Rules = parseRules(text)
	rf.Count = len(rf.Rules)

	return rf
}
