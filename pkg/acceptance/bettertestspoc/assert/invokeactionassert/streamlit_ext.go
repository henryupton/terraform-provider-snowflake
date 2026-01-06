package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func StreamlitDoesNotExist(t *testing.T, id sdk.SchemaObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeStreamlit,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.Streamlit, sdk.SchemaObjectIdentifier] {
			return testClient.Streamlit.Show
		})
}
