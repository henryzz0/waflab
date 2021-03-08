// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package parse

import (
	"testing"

	"github.com/waflab/waflab/object"
)

func TestParseRule(t *testing.T) {
	object.InitOrmManager()

	parseEnabledRuleFile()
}
