package rule

import (
	"github.com/hsluoyz/modsecurity-go/seclang/parser"

	proto "github.com/waflab/waflab/rule/proto"
)

//go:generate protoc -I proto --go_out=plugins=grpc:proto proto/rule.proto proto/variables.proto proto/actions.proto proto/operators.proto

func mapAction(action *parser.Action) *proto.Action {
	res := &proto.Action{
		Name:  parser.ActionNameMap[action.Tk],
		Param: action.Argument,
		Kind:  proto.Action_RUN_TIME_ONLY_IF_MATCH,
	}
	return res
}

func mapTransformer(transformer *parser.Trans) *proto.Action {
	res := &proto.Action{
		Name:  parser.TransformationNameMap[transformer.Tk],
		Param: "",
		Kind:  proto.Action_RUN_TIME_BEFORE_MATCH_ATTEMPT,
	}
	return res
}

func mapVariable(variable *parser.Variable) *proto.Variable {
	res := &proto.Variable{
		CollectionName:       "",
		Name:                 "",
		IsCount:              variable.Count,
		//KeyExclusion:         variable.Exclusion,
	}

	return res
}

func mapOperator(operator *parser.Operator) *proto.Operator {
	return nil
}

func mapRule(rd *parser.RuleDirective) *proto.Rule {
	res := &proto.Rule{}

	for _, variable := range rd.Variable {
		element := mapVariable(variable)
		res.Variables = append(res.Variables, element)
	}

	res.Op = mapOperator(rd.Operator)

	res.Id = int64(rd.Actions.Id)
	res.Phase = int32(rd.Actions.Phase)
	res.Chained = rd.Actions.Chain

	for _, action := range rd.Actions.Action {
		element := mapAction(action)
		res.ActionsRuntimePos = append(res.ActionsRuntimePos, element)
	}
	for _, transformer := range rd.Actions.Trans {
		element := mapTransformer(transformer)
		res.ActionsRuntimePre = append(res.ActionsRuntimePre, element)
	}

	return nil
}

func getProtobufRulesFromRules(rule *Rule) *proto.RuleList {
	res := &proto.RuleList{}

	element := mapRule(rule.Data)
	res.Item = append(res.Item, element)

	for _, chainRule := range rule.ChainRules {
		element := mapRule(chainRule.Data)
		res.Item = append(res.Item, element)
	}

	return res
}
