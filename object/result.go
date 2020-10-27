package object

import (
	"io/ioutil"
	"net/http"
)

type Result struct {
	Status   int    `json:"status"`
	Response string `json:"response"`
}

func GetResult(testsetId string, testcaseId string) *Result {
	testset := GetTestset(testsetId)
	testcase := GetTestcase(testcaseId)

	client := &http.Client{}
	req, err := http.NewRequest(testcase.Method, testset.TargetUrl, nil)
	if err != nil {
		panic(err)
	}

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
