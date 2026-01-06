package resourceassert

import (
	"fmt"
	"strconv"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func (s *StreamlitResourceAssert) HasExternalAccessIntegrations(expected []string) *StreamlitResourceAssert {
	s.AddAssertion(assert.ValueSet("external_access_integrations.#", strconv.FormatInt(int64(len(expected)), 10)))
	for i, val := range expected {
		s.AddAssertion(assert.ValueSet(fmt.Sprintf("external_access_integrations.%d", i), val))
	}
	return s
}
