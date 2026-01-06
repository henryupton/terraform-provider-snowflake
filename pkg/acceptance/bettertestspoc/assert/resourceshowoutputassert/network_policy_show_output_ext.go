package resourceshowoutputassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func (n *NetworkPolicyShowOutputAssert) HasCreatedOnNotEmpty() *NetworkPolicyShowOutputAssert {
	n.AddAssertion(assert.ResourceShowOutputValuePresent("created_on"))
	return n
}
