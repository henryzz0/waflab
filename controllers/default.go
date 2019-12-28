package controllers

import (
	"github.com/astaxie/beego"
	"github.com/waflab/waflab/rule"
)

type ApiController struct {
	beego.Controller
}

func listTestsets() []*rule.Testset {
	testSets := []*rule.Testset{}

	testSets = append(testSets, &rule.Testset{
		Id:      "crs-3.3",
		Name:    "CRS",
		Version: "v3.3/dev",
		Count:   0,
	})

	return testSets
}

func (c *ApiController) ListTestsets() {
	c.Data["json"] = listTestsets()
	c.ServeJSON()
}
