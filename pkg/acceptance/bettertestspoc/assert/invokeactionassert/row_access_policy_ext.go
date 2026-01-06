package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func RowAccessPolicyDoesNotExist(t *testing.T, id sdk.SchemaObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeRowAccessPolicy,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.RowAccessPolicy, sdk.SchemaObjectIdentifier] {
			return testClient.RowAccessPolicy.Show
		})
}
