// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package util

import (
	"io/ioutil"
	"strconv"
)

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ReadStringFromPath(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func WriteStringToPath(s string, path string) {
	err := ioutil.WriteFile(path, []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}

func StringListContains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func IntListContains(list []int, i int) bool {
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return false
}
