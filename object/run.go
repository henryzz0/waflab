package object

import (
	"fmt"

	"github.com/waflab/waflab/docker"
	"github.com/waflab/waflab/util"
)

func getWafBenchResult(testset *Testset, testcase *Testcase) *Result {
	url := testset.TargetUrl

	folder := docker.WriteTestcaseToFile(testcase.Name, testcase.RawData)
	status := docker.GetStatusFromContainer(folder, url)
	fmt.Printf("True HTTP status: [%d].\n", status)

	tf := testcase.Data
	test := tf.Tests[0]
	stage := test.Stages[0].Stage
	output := stage.Output
	expectedStatusList := output.Status

	res := &Result{}
	res.Status = status
	if util.IntListContains(expectedStatusList, res.Status) {
		res.Response = fmt.Sprintf("ok: %v == %d", expectedStatusList, res.Status)
	} else {
		res.Response = fmt.Sprintf("wrong: %v != %d", expectedStatusList, res.Status)
	}
	return res
}
