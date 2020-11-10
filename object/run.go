package object

import (
	"fmt"
	"strings"

	"github.com/waflab/waflab/docker"
	"github.com/waflab/waflab/util"
)

func getWafBenchResult(testset *Testset, testcase *Testcase) *Result {
	url := testset.TargetUrl

	folder := docker.WriteTestcaseToFile(testcase.Name, testcase.RawData)
	statuses := docker.GetStatusFromContainer(folder, url)
	fmt.Printf("True HTTP statuses: %v\n", statuses)

	res := &Result{}
	res.Statuses = statuses

	isCorrect := true
	reasons := []string{}
	tf := testcase.Data
	for i, test := range tf.Tests {
		stage := test.Stages[0].Stage
		output := stage.Output
		expectedStatusList := output.Status

		trueStatus := res.Statuses[i]
		var reason string
		if util.IntListContains(expectedStatusList, trueStatus) {
			reason = fmt.Sprintf("%v == %d", expectedStatusList, trueStatus)
		} else {
			isCorrect = false
			reason = fmt.Sprintf("%v != %d", expectedStatusList, trueStatus)
		}
		reasons = append(reasons, reason)
	}

	if isCorrect {
		res.Response = "ok: "
	} else {
		res.Response = "wrong: "
	}
	res.Response += strings.Join(reasons, ", ")

	return res
}
