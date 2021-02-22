// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetPath(path string) string {
	return filepath.Dir(path)
}

func EnsureFileFolderExists(path string) {
	p := GetPath(path)
	if !FileExist(p) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func GetAbsolutePath(path string) string {
	res, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return res
}

func RemovePath(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func filterFile(name string) bool {
	return strings.HasSuffix(name, ".conf")
}

func ListFileIds(path string) []string {
	res := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.IsDir() && filterFile(f.Name()) {
			fileId := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			res = append(res, fileId)
		}
	}

	return res
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
