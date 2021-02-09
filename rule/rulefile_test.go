package rule

import (
	"testing"

	"github.com/waflab/waflab/object"
	"github.com/waflab/waflab/test"
	"github.com/waflab/waflab/util"
)

func TestSyncTestfile(t *testing.T) {
	object.InitOrmManager()

	text := util.ReadStringFromPath(util.CrsTestDir + "REQUEST-920-PROTOCOL-ENFORCEMENT/920100.yaml")

	tf := test.LoadTestfileFromString(text)
	SyncTestfile(tf, text, "regression-test")
}
