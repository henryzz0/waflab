// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

import (
	"testing"

	"github.com/waflab/waflab/docker"
)

func TestLoadTraffic(t *testing.T) {
	InitOrmManager()

	testcase := GetTestcase("920290.yaml")
	folder := docker.WriteTestcaseToFile(testcase.Name, testcase.RawData)
	println(folder)
}

func TestRunContainer(t *testing.T) {
	InitOrmManager()

	testcase := GetTestcase("920280.yaml")

	folder := docker.WriteTestcaseToFile(testcase.Name, testcase.RawData)
	//folder := "I:/github_repos/waflab/tmpFiles/920290.yaml"
	url := "http://test.waflab.org:7080"
	statuses := docker.GetStatusFromContainer(folder, url)
	println(statuses)
}
