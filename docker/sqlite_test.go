// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package docker

import "testing"

func TestReadTraffic(t *testing.T) {
	path := "C:/wafbench/regression.db"
	ormManager := newOrmManager(path)
	traffics := getTraffics(ormManager)
	println(traffics)

	//testcase := object.GetTestcase("920290.yaml")
	//saveTestcaseIntoDb(testcase, "C:/wafbench/regression.db")
}
