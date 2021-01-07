package payload

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/waflab/waflab/autogen/utils"

	"github.com/waflab/waflab/test"
)

const (
	randomStringLength = 10
)

func composeCookie(payload *test.Input, key string, value string) {
	composeHeader(payload, "Cookie", fmt.Sprintf("%s=%s", key, value))
}

func composeQueryString(payload *test.Input, key string, value string) {
	payload.Uri = fmt.Sprintf("/?%s=%s", url.QueryEscape(key), url.QueryEscape(value))
}

func composeHeader(payload *test.Input, key string, value string) {
	payload.Headers[key] = value
}

func composeFile(payload *test.Input, name string, value string) {
	composeHeader(payload, "Content-Type", "multipart/form-data; boundary=----abc")
	composeHeader(payload, "Cache-Control", "no-cache")
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

func addArg(value string, payload *test.Input) error {
	key := strings.ReplaceAll(utils.RandomString(randomStringLength), "_", "")
	composeQueryString(payload, key, value)
	return nil
}

func addArgNames(value string, payload *test.Input) error {
	composeQueryString(payload, value, utils.RandomString(randomStringLength))
	return nil
}

func addFilesNames(value string, payload *test.Input) error {
	composeFile(payload, value, "1")
	return nil
}

func addFiles(value string, payload *test.Input) error {
	composeFile(payload, "files[]", value)
	return nil
}

func addQueryString(value string, payload *test.Input) error {
	composeQueryString(payload, value, "")
	return nil
}

func addRequestBody(value string, payload *test.Input) error {
	payload.Method = "POST"
	payload.Data = append(payload.Data, fmt.Sprintf("Foo_Key=%s", value))
	composeHeader(payload, "Content-Length", strconv.Itoa(len(payload.Data[0])))
	composeHeader(payload, "Content-Type", "application/x-www-form-urlencoded")
	return nil
}

func addRequestCookies(value string, payload *test.Input) error {
	composeCookie(payload, utils.RandomString(randomStringLength), value)
	return nil
}

func addRequestCookiesName(value string, payload *test.Input) error {
	composeCookie(payload, value, utils.RandomString(randomStringLength))
	return nil
}

func addRequestHeaders(value string, payload *test.Input) error {
	composeHeader(payload, utils.RandomString(randomStringLength), value)
	return nil
}

func addRequestHeadersNames(value string, payload *test.Input) error {
	composeHeader(payload, value, utils.RandomString(randomStringLength))
	return nil
}

func addRequestLine(value string, payload *test.Input) error {
	payload.RawRequest = value
	return nil
}

func addRequestMethod(value string, payload *test.Input) error {
	payload.Method = value
	return nil
}

func addRequestProtocol(value string, payload *test.Input) error {
	payload.Protocol = value
	return nil
}

func addRequestURI(value string, payload *test.Input) error {
	payload.Uri = value
	return nil
}
