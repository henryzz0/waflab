package object

import (
	"path/filepath"

	"github.com/waflab/waflab/object"
	"github.com/waflab/waflab/util"
)

func writeTestcaseToFile(testcaseName string) string {
	testcase := object.GetTestcase(testcaseName)
	path := util.GetTmpYamlPath("../../tmpFiles/" + testcase.Name + "/" +testcase.Name)
	util.EnsureFileFolderExists(path)
	util.WriteStringToPath(testcase.RawData, path)

	res := util.GetAbsolutePath(path)
	res = util.GetPath(res)
	res = filepath.ToSlash(res)
	return res
}
