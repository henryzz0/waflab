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

func GetResult(testsetId string, testcaseId string) *Result {
	testset := GetTestset(testsetId)
	testcase := GetTestcase(testcaseId)

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
	//for key, value := range input.Headers {
	//	req.Header.Add(key, value)
	//}

	client := &http.Client{}
	req, err := http.NewRequest(testcase.Method, testset.TargetUrl + query, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", testcase.UserAgent)

	//for key, value := range input.Headers {
	//	req.Header.Add(key, value)
	//}

	resp, err := client.Do(req)
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
