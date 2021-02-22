// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package test

import (
	"regexp"

	"gopkg.in/yaml.v2"
)

var reStatus *regexp.Regexp

func init() {
	reStatus, _ = regexp.Compile("status: (\\d+)")
}

func fixStatus(s string) string {
	// status: 200 --> status: [200]
	s = reStatus.ReplaceAllString(s, "status: [$1]")
	return s
}

func LoadTestfileFromString(text string) *Testfile {
	text = fixStatus(text)

	tf := Testfile{}

	err := yaml.Unmarshal([]byte(text), &tf)
	if err != nil {
		panic(err)
	}

	return &tf
}
