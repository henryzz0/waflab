package controllers

import (
	"github.com/astaxie/beego"
	"github.com/waflab/waflab/rule"
	"github.com/waflab/waflab/util"
)

type ApiController struct {
	beego.Controller
}


var rsm map[string]*rule.Ruleset

func init() {
	rsm = map[string]*rule.Ruleset{}
}

func getOrCreateRs(id string) *rule.Ruleset {
	var rs *rule.Ruleset

	if _, ok := rsm[id]; ok {
		rs = rsm[id]
	} else {
		rs = rule.ReadRuleset(id)
		rsm[id] = rs
	}

	return rs
}

func listRulesetIds() []string {
	res := []string{}
	res = append(res, "crs-3.3")
	return res
}

func listRulesets() []*rule.Ruleset {
	res := []*rule.Ruleset{}

	for _, id := range listRulesetIds() {
		getOrCreateRs(id)
	}

	m := map[string]interface{}{}
	for k, v := range rsm {
		m[k] = v
	}
	kv := util.SortMapsByKey(&m)
	for _, v := range *kv {
		res = append(res, v.Key.(*rule.Ruleset))
	}

	return res
}

func (c *ApiController) ListRulesets() {
	c.Data["json"] = listRulesets()
	c.ServeJSON()
}
