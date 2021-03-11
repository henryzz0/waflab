// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/waflab/waflab/docker"
	"github.com/waflab/waflab/util"
)

var master *docker.Master

func InitMaster() {
	master = docker.MakeMaster(1)
}

func getWafBenchResult(testset *Testset, testcase *Testcase) *Result {
	url := testset.TargetUrl

	statuses := make([]int, 0)
	hitrules := make([]string, 0)
	responses, err := master.InsertTask(url, testcase.RawData)
	if err != nil {
		panic(err)
	}
	for _, resp := range responses {
		status, err := strconv.Atoi(resp.Status)
		if err != nil {
			status = 0
		}
		statuses = append(statuses, status)
		hitrules = append(hitrules, resp.HitRule)
	}
	fmt.Printf("True HTTP statuses: %v\n", statuses)

	res := &Result{}
	res.Statuses = statuses
	res.HitRules = hitrules

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
		res.Response += strings.Join(reasons, ", ")
	}

	return res
}
