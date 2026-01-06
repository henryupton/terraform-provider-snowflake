package resourceassert

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (s *SecondaryConnectionResourceAssert) HasAsReplicaOfIdentifier(expected sdk.ExternalObjectIdentifier) *SecondaryConnectionResourceAssert {
	s.AddAssertion(assert.ValueSet("as_replica_of", expected.FullyQualifiedName()))
	return s
}
