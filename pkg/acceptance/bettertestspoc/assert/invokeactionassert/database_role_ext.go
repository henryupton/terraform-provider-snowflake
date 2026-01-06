package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func DatabaseRoleDoesNotExist(t *testing.T, id sdk.DatabaseObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeDatabaseRole,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.DatabaseRole, sdk.DatabaseObjectIdentifier] {
			return testClient.DatabaseRole.Show
		})
}
