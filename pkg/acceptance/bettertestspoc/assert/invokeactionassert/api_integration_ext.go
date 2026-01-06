package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func ApiIntegrationDoesNotExist(t *testing.T, id sdk.AccountObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeApiIntegration,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.ApiIntegration, sdk.AccountObjectIdentifier] {
			return testClient.ApiIntegration.Show
		})
}
