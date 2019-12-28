package controllers

import (
	"github.com/astaxie/beego"
	"github.com/waflab/waflab/rule"
	"github.com/waflab/waflab/util"
)

type ApiController struct {
	beego.Controller
}

func listTestsetIds() []string {
	res := []string{}
	res = append(res, "crs-3.3")
	return res
}

func listTestsets() []*rule.Testset {
	res := []*rule.Testset{}

	for _, id := range listTestsetIds() {
		getOrCreateTs(id)
	}

	m := map[string]interface{}{}
	for k, v := range tsm {
		m[k] = v
	}
	kv := util.SortMapsByKey(&m)
	for _, v := range *kv {
		res = append(res, v.Key.(*rule.Testset))
	}

	return res
}

func (c *ApiController) ListTestsets() {
	c.Data["json"] = listTestsets()
	c.ServeJSON()
}
