package object

import (
	"testing"

	"github.com/waflab/waflab/object"
)

func TestLoadTraffic(t *testing.T) {
	object.InitOrmManager()

	folder := writeTestcaseToFile("920290.yaml")
	println(folder)
}

func TestRunContainer(t *testing.T) {
	//object.InitOrmManager()
	//folder := writeTestcaseToFile("920290.yaml")

	folder := "I:/github_repos/waflab/tmpFiles/920290.yaml"
	runContainer(folder)
}
