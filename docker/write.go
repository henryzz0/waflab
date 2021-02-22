// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package docker

import (
	"path/filepath"
	"regexp"

	"github.com/waflab/waflab/util"
)

var reResponse *regexp.Regexp

func init() {
	reResponse, _ = regexp.Compile("HTTP/\\d\\.\\d (\\d+)")
}

func getStatus(text string) int {
	res := reResponse.FindStringSubmatch(text)
	if res == nil {
		return -1
	}

	return util.ParseInt(res[1])
}

func WriteTestcaseToFile(testcaseName string, data string) string {
	path := util.GetTmpYamlPath("../../tmpFiles/" + testcaseName + "/" + testcaseName)
	util.EnsureFileFolderExists(path)
	util.WriteStringToPath(data, path)

	res := util.GetAbsolutePath(path)
	res = util.GetPath(res)
	res = filepath.ToSlash(res)
	return res
}

func readDbResult(path string) []int {
	ormManager := newOrmManager(path)
	traffics := getTraffics(ormManager)
	err := ormManager.Close()
	if err != nil {
		panic(err)
	}

	res := []int{}
	for _, traffic := range traffics {
		resp := traffic.RawResponse
		status := getStatus(resp)
		res = append(res, status)
	}
	return res
}
