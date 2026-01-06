package resourceassert

import (
	"fmt"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (c *PrimaryConnectionResourceAssert) HasExactlyFailoverToAccountsInOrder(expected ...sdk.AccountIdentifier) *PrimaryConnectionResourceAssert {
	c.AddAssertion(assert.ValueSet("enable_failover_to_accounts.#", fmt.Sprintf("%d", len(expected))))
	for i, v := range expected {
		c.AddAssertion(assert.ValueSet(fmt.Sprintf("enable_failover_to_accounts.%d", i), v.Name()))
	}
	return c
}
