package resourceassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (a *AccountResourceAssert) HasAdminUserType(expected sdk.UserType) *AccountResourceAssert {
	a.AddAssertion(assert.ValueSet("admin_user_type", string(expected)))
	return a
}
