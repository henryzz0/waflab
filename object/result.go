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
	Status   int    `json:"status"`
	Response string `json:"response"`
}

func getResult(testset *Testset, testcase *Testcase) *Result {
	if testcase.Data == nil {
		return getFingerprintingResult(testset, testcase)
	} else {
		return getWafResult(testset, testcase)
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

func getFingerprintingResult(testset *Testset, testcase *Testcase) *Result {
	method := testcase.Method
	host := testset.TargetUrl
	uri := ""
	query := getQuery(testcase)
	userAgent := testcase.UserAgent

	resp, err := sendRaw(method, host, uri, query, userAgent, nil)
	if err != nil {
		//panic(err)
		res := &Result{
			Status:   -2,
			Response: "No connection",
		}
		return res
	}
	defer resp.Body.Close()

	res := &Result{}
	res.Status = resp.StatusCode
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		res.Response = string(bodyBytes)
	}

	return res
}

func GetResult(testsetId string, testcaseId string) *Result {
	testset := GetTestset(testsetId)
	testcase := GetTestcase(testcaseId)

	return getResult(testset, testcase)
}
