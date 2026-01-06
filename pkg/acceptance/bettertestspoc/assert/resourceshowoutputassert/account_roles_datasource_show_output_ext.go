package resourceshowoutputassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func AccountRolesDatasourceShowOutput(t *testing.T, datasourceReference string) *RoleShowOutputAssert {
	t.Helper()

	s := RoleShowOutputAssert{
		ResourceAssert: assert.NewDatasourceAssert(datasourceReference, "show_output", "account_roles.0."),
	}
	s.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &s
}
