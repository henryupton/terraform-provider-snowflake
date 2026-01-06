package resourceassert

import "github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"

func (s *SemanticViewResourceAssert) HasNoTables() *SemanticViewResourceAssert {
	s.AddAssertion(assert.ValueNotSet("tables"))
	return s
}
