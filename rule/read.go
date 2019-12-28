package rule

import "fmt"

func ReadTestset(id string) *Testset {
	fmt.Printf("Read testset for Id: [%s].\n", id)

	ts := newTestset(id)
	if ts.Id == "crs-3.3" {
		ts.Name = "CoreRuleSet"
		ts.Version = "v3.3/dev"
	}

	//filenames := util.ListFileIds(util.CrsRuleDir)

	return ts
}
