// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package parse

import (
	"encoding/json"
	"fmt"

	"github.com/waflab/waflab/util"
)

type WafRulesFile struct {
	RuleSets []WafRuleSet `json:"value"`
}

type WafRuleSet struct {
	Name       string        `json:"name"`
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Properties WafProperties `json:"properties"`
}

type WafProperties struct {
	ProvisioningState string         `json:"provisioningState"`
	RuleSetId         string         `json:"ruleSetId"`
	RuleSetType       string         `json:"ruleSetType"`
	RuleSetVersion    string         `json:"ruleSetVersion"`
	RuleGroups        []WafRuleGroup `json:"ruleGroups"`
}

type WafRuleGroup struct {
	RuleGroupName string    `json:"ruleGroupName"`
	Description   string    `json:"description"`
	Rules         []WafRule `json:"rules"`
}

type WafRule struct {
	RuleId        string `json:"ruleId"`
	Description   string `json:"description"`
	DefaultAction string `json:"defaultAction"`
	DefaultState  string `json:"defaultState"`
}

func parseEnabledRuleFile() {
	s := util.ReadStringFromPath(util.EnabledRulePath)
	rulesFile := WafRulesFile{}
	err := json.Unmarshal([]byte(s), &rulesFile)
	if err != nil {
		panic(err)
	}

	for _, ruleSet := range rulesFile.RuleSets {
		ruleGroups := ruleSet.Properties.RuleGroups
		for _, ruleGroup := range ruleGroups {
			for _, rule := range ruleGroup.Rules {
				fmt.Printf("ruleId: %s, defaultAction: %s, defaultState: %s\n", rule.RuleId, rule.DefaultAction, rule.DefaultState)
			}
		}
	}
}
