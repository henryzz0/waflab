package object

import (
	"fmt"
	"testing"
)

// go test object/testset_test.go

func testRunTestcase(t *testing.T, testset *Testset, testcase *Testcase) {
	t.Helper()

	status := testcase.Status
	res := getResult(testset, testcase)
	myStatus := res.Status
	response := res.Response
	if response == "" {
		response = "(Empty)"
	}

	if status == myStatus {
		fmt.Printf("(%s, %s) --> %d, response = %s\n",
			testset.Name,
			testcase.Name,
			status,
			response)
	} else {
		t.Errorf("(%s, %s) --> %d, supposed to be %d, response = %s\n",
			testset.Name,
			testcase.Name,
			myStatus,
			status,
			response)
	}
}

func TestRunTestset(t *testing.T) {
	InitOrmManager()

	testsetName := "fingerprinting"

	testset := GetTestset(testsetName)
	testcases := getFilteredTestcases(testset)

	for _, testcase := range testcases {
		testRunTestcase(t, testset, testcase)
	}
}
