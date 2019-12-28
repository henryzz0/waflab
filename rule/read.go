package rule

import "fmt"

func ReadRuleset(id string) *Ruleset {
	fmt.Printf("Read ruleset for Id: [%s].\n", id)

	rs := newRuleset(id)
	if rs.Id == "crs-3.3" {
		rs.Name = "CoreRuleSet"
		rs.Version = "v3.3/dev"
	}

	//filenames := util.ListFileIds(util.CrsRuleDir)

	return rs
}
