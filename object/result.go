// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Result struct {
	Statuses []int    `json:"statuses"`
	HitRules []string `json:"hitRules"`
	Response string   `json:"response"`
}

func getResult(testset *Testset, testcase *Testcase, typ string) *Result {
	if testcase.Data == nil {
		return getFingerprintingResult(testset, testcase, typ)
	} else {
		//return getWafResult(testset, testcase)
		return getWafBenchResult(testset, testcase, typ)
	}
}

func getQuery(testcase *Testcase) string {
	tokens := []string{}

	for _, pair := range testcase.QueryStrings {
		value := pair.Value
		if pair.Key == "data" {
			value = base64.StdEncoding.EncodeToString([]byte(value))
		}
		value = url.QueryEscape(value)
		tokens = append(tokens, fmt.Sprintf("%s=%s", pair.Key, value))
	}

	query := fmt.Sprintf("?%s", strings.Join(tokens, "&"))
	return query
}

func getFingerprintingResult(testset *Testset, testcase *Testcase, typ string) *Result {
	method := testcase.Method
	host := getTestUrl(testset, typ)
	uri := ""
	query := getQuery(testcase)
	userAgent := testcase.UserAgent

	resp, err := sendRaw(method, host, uri, query, userAgent, nil)
	if err != nil {
		//panic(err)
		res := &Result{
			Statuses: []int{-2},
			Response: "No connection",
		}
		return res
	}
	defer resp.Body.Close()

	res := &Result{}
	res.Statuses = []int{resp.StatusCode}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		res.Response = string(bodyBytes)
	}

	return res
}

func GetResult(testsetId string, testcaseId string, typ string) *Result {
	testset := GetTestset(testsetId)
	testcase := GetTestcase(testcaseId)

	result := getResult(testset, testcase, typ)

	if typ == "baseline" {
		testcase.BaselineStatuses = result.Statuses
	} else if typ == "target" {
		testcase.TrueStatuses = result.Statuses
		testcase.Result = result.Response
		testcase.HitRules = result.HitRules
	}

	UpdateTestcase(testcaseId, testcase)
	return result
}
