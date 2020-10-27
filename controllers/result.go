package controllers

import "github.com/waflab/waflab/object"

func (c *ApiController) GetResult() {
	testsetId := c.Input().Get("testsetId")
	testcaseId := c.Input().Get("testcaseId")

	c.Data["json"] = object.GetResult(testsetId, testcaseId)
	c.ServeJSON()
}
