package objectassert

import (
	"fmt"
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (s *SecurityIntegrationAssert) HasCreatedOnNotEmpty() *SecurityIntegrationAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.SecurityIntegration) error {
		t.Helper()
		if o.CreatedOn.IsZero() {
			return fmt.Errorf("expected created_on to be set, but it was zero")
		}
		return nil
	})
	return s
}
