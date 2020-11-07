package test

import "gopkg.in/yaml.v2"

func LoadTestfileFromString(text string) *Testfile {
	tf := Testfile{}

	err := yaml.Unmarshal([]byte(text), &tf)
	if err != nil {
		panic(err)
	}

	return &tf
}
