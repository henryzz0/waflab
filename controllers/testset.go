package controllers

import (
	"encoding/json"

	"github.com/waflab/waflab/object"
)

func (c *ApiController) GetTestsets() {
	c.Data["json"] = object.GetTestsets()
	c.ServeJSON()
}

func (c *ApiController) GetTestset() {
	id := c.Input().Get("id")

	c.Data["json"] = object.GetTestset(id)
	c.ServeJSON()
}

func (c *ApiController) UpdateTestset() {
	id := c.Input().Get("id")

	var testset object.Testset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.UpdateTestset(id, &testset)
	c.ServeJSON()
}

func (c *ApiController) AddTestset() {
	var testset object.Testset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.AddTestset(&testset)
	c.ServeJSON()
}

func (c *ApiController) DeleteTestset() {
	var testset object.Testset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &testset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.DeleteTestset(&testset)
	c.ServeJSON()
}
