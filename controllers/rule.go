// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

func (c *ApiController) ListRules() {
	rulesetId := c.Input().Get("rulesetId")
	rulefileId := c.Input().Get("rulefileId")
	rs := getOrCreateRs(rulesetId)
	rf := rs.RulefileMap[rulefileId]

	c.Data["json"] = rf
	c.ServeJSON()
}
