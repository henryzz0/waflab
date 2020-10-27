package controllers

import (
	"encoding/json"

	"github.com/waflab/waflab/object"
)

func (c *ApiController) GetTestcases() {
	c.Data["json"] = object.GetTestcases()
	c.ServeJSON()
}

func (c *ApiController) GetTestcase() {
	id := c.Input().Get("id")

	c.Data["json"] = object.GetTestcase(id)
	c.ServeJSON()
}

func (c *ApiController) UpdateTestcase() {
	id := c.Input().Get("id")

	var testcase object.Testcase
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testcase)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.UpdateTestcase(id, &testcase)
	c.ServeJSON()
}

func (c *ApiController) AddTestcase() {
	var testcase object.Testcase
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testcase)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.AddTestcase(&testcase)
	c.ServeJSON()
}

func (c *ApiController) DeleteTestcase() {
	var testcase object.Testcase
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testcase)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.DeleteTestcase(&testcase)
	c.ServeJSON()
}
