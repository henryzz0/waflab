// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package test

import (
	"testing"

	"github.com/waflab/waflab/util"
)

func TestParseTestset(t *testing.T) {
	//rf := newRulefile(0, "REQUEST-920-PROTOCOL-ENFORCEMENT")
	text := util.ReadStringFromPath(util.CrsTestDir + "REQUEST-920-PROTOCOL-ENFORCEMENT/920100.yaml")
	//parseRules(rf, text)
	//rf.syncPls()
	//printRules(rf)

	LoadTestfileFromString(text)
}
