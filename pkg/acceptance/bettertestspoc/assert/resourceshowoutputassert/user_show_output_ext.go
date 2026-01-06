package resourceshowoutputassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func (u *UserShowOutputAssert) HasCreatedOnNotEmpty() *UserShowOutputAssert {
	u.AddAssertion(assert.ResourceShowOutputValuePresent("created_on"))
	return u
}

func (u *UserShowOutputAssert) HasDaysToExpiryNotEmpty() *UserShowOutputAssert {
	u.AddAssertion(assert.ResourceShowOutputValuePresent("days_to_expiry"))
	return u
}

func (u *UserShowOutputAssert) HasMinsToUnlockNotEmpty() *UserShowOutputAssert {
	u.AddAssertion(assert.ResourceShowOutputValuePresent("mins_to_unlock"))
	return u
}

func (u *UserShowOutputAssert) HasMinsToBypassMfaNotEmpty() *UserShowOutputAssert {
	u.AddAssertion(assert.ResourceShowOutputValuePresent("mins_to_bypass_mfa"))
	return u
}

func (t *UserShowOutputAssert) HasTypeEmpty() *UserShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("type", ""))
	return t
}
