package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func ViewDoesNotExist(t *testing.T, id sdk.SchemaObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeView,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.View, sdk.SchemaObjectIdentifier] {
			return testClient.View.Show
		})
}
