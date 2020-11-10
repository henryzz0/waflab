package util

import "fmt"

func GetTmpYamlPath(fileId string) string {
	return fmt.Sprintf("tmpFiles/%s", fileId)
}
