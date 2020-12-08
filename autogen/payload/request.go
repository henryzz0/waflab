package payload

import (
	"fmt"
	"strings"

	"github.com/waflab/waflab/test"
)

func addArg(value string, payload *test.Input) error {
	key := strings.ReplaceAll(randomString(10), "_", "")
	payload.Uri = fmt.Sprintf("/?%s=%s", key, value)
	return nil
}
