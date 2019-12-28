package util

import "io/ioutil"

func ReadStringFromPath(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

