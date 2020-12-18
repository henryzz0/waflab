package payload

import (
	"fmt"
	"strings"

	"github.com/waflab/waflab/autogen/utils"

	"github.com/waflab/waflab/test"
)

func composeUrl(key string, value string) string {
	return fmt.Sprintf("/?%s=%s", key, value)
}

func addArg(value string, payload *test.Input) error {
	key := strings.ReplaceAll(utils.RandomString(10), "_", "")
	payload.Uri = composeUrl(key, value)
	return nil
}

func addArgNames(value string, payload *test.Input) error {
	payload.Uri = composeUrl(value, utils.RandomString(10))
	return nil
}

func composeFile(name string, value string, payload *test.Input) {
	payload.Headers["Content-Type"] = "multipart/form-data; boundary=----abc"
	payload.Headers["Cache-Control"] = "no-cache"
	payload.Method = "POST"
	payload.Data = []string{
		"------abc",
		fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"", name, value),
		"Content-Type: text/plain",
		"",
		"Content ",
		"",
		"------abc--",
	}
}

func addFilesNames(value string, payload *test.Input) error {
	composeFile("files[]", value, payload)
	return nil
}

func addFiles(value string, payload *test.Input) error {
	composeFile(value, "1", payload)
	return nil
}
