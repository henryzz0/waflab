package controllers

func (c *ApiController) ListRulefiles() {
	rulesetId := c.Input().Get("rulesetId")
	rs := getOrCreateRs(rulesetId)

	c.Data["json"] = rs
	c.ServeJSON()
}
