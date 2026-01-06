package resourceshowoutputassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func (s *SecurityIntegrationShowOutputAssert) HasCreatedOnNotEmpty() *SecurityIntegrationShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValuePresent("created_on"))
	return s
}
