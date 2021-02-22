// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

import (
	"fmt"
	"testing"

	"github.com/waflab/waflab/util"
)

// go test object/testset_test.go

func testRunTestcase(t *testing.T, testset *Testset, testcase *Testcase) {
	t.Helper()

	res := getResult(testset, testcase)
	for i, statusList := range testcase.StatusLists {
		response := res.Response
		if response == "" {
			response = "(Empty)"
		}

		myStatus := res.Statuses[i]
		if util.IntListContains(statusList, myStatus) {
			fmt.Printf("(%s, %s) --> %v, response = %s\n",
				testset.Name,
				testcase.Name,
				statusList,
				response)
		} else {
			t.Errorf("(%s, %s) --> %d, supposed to be %v, response = %s\n",
				testset.Name,
				testcase.Name,
				myStatus,
				statusList,
				response)
		}
	}
}

func TestRunTestset(t *testing.T) {
	InitOrmManager()

	testsetName := "fingerprinting-test"

	testset := GetTestset(testsetName)
	testcases := getFilteredTestcases(testset)

	for _, testcase := range testcases {
		testRunTestcase(t, testset, testcase)
	}
}
