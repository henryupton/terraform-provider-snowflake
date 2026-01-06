package resourceshowoutputassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func NetworkPoliciesDatasourceShowOutput(t *testing.T, datasourceReference string) *NetworkPolicyShowOutputAssert {
	t.Helper()

	n := NetworkPolicyShowOutputAssert{
		ResourceAssert: assert.NewDatasourceAssert(datasourceReference, "show_output", "network_policies.0."),
	}
	n.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &n
}
