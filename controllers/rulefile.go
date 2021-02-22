// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

func (c *ApiController) ListRulefiles() {
	rulesetId := c.Input().Get("rulesetId")
	rs := getOrCreateRs(rulesetId)

	c.Data["json"] = rs
	c.ServeJSON()
}
