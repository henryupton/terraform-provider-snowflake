package resourceparametersassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (f *FunctionResourceParametersAssert) HasAllDefaults() *FunctionResourceParametersAssert {
	return f.
		HasEnableConsoleOutput(false).
		HasLogLevel(sdk.LogLevelOff).
		HasMetricLevel(sdk.MetricLevelNone).
		HasTraceLevel(sdk.TraceLevelOff)
}
