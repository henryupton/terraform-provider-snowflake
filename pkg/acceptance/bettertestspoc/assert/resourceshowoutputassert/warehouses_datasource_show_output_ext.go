package resourceshowoutputassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func WarehousesDatasourceShowOutput(t *testing.T, datasourceReference string) *WarehouseShowOutputAssert {
	t.Helper()

	s := WarehouseShowOutputAssert{
		ResourceAssert: assert.NewDatasourceAssert(datasourceReference, "show_output", "warehouses.0."),
	}
	s.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &s
}
