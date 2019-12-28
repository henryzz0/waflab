package controllers

import "github.com/waflab/waflab/rule"

var tsm map[string]*rule.Testset

func init() {
	tsm = map[string]*rule.Testset{}
}

func getOrCreateTs(id string) *rule.Testset {
	var ts *rule.Testset

	if _, ok := tsm[id]; ok {
		ts = tsm[id]
	} else {
		ts = rule.ReadTestset(id)
		tsm[id] = ts
	}

	return ts
}
