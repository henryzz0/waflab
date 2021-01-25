package autogen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/waflab/waflab/test"
	y "gopkg.in/yaml.v2"
)

// GenerateTestFromDirectory read ModSecurity rule from dirPath and
// then write the generated test as yaml file into the ouput directory.
// It will create a folder with the name of config to hold all generated .yaml test file.
// Notice that it only read rules file ends with .conf
func GenerateTestFromDirectory(dirPath, output string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error %v at %q\n", err, path)
			return nil
		}

		// only process file with conf suffix
		if info.Mode().IsRegular() && filepath.Ext(path) == ".conf" {
			ruleStrings, err := readRuleStringFromConf(path)
			if err != nil {
				fmt.Printf("error %v when read %q\n", err, path)
				return nil
			}

			// generate testfiles from rules
			var tests []*test.Testfile
			for _, rs := range ruleStrings {
				tests = append(tests, GenerateTests(rs, 10)...)
			}

			// write generated tests into files
			testDir := filepath.Join(filepath.Dir(output), info.Name())
			os.MkdirAll(testDir, os.ModePerm) // make a directory with the name of config
			for _, test := range tests {
				out, err := y.Marshal(test)
				if err != nil {
					fmt.Printf("error %v when marshal %s\n", err, test.Meta.Name)
					continue
				}
				err = ioutil.WriteFile(filepath.Join(testDir, test.Meta.Name), out, os.ModePerm)
				if err != nil {
					fmt.Printf("error %v when write %s\n", err, test.Meta.Name)
				}
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// readRuleStringFromConf read rule string from config and
// remove any additional comments
func readRuleStringFromConf(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// replace all comments
	commentRegexp, _ := regexp.Compile("(?:^|\n)#.*")
	str := commentRegexp.ReplaceAllString(string(content), "")

	// Match SecRule only
	var rules []string
	for _, s := range strings.Split(str, "\r\n\r\n") {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "SecRule") {
			s = strings.ReplaceAll(s, "\r\n", "\n") // CRLF to LF
			rules = append(rules, s)
		}
	}
	return rules, nil
}
