package objectassert

import (
	"fmt"
	"testing"
	"time"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (a *AuthenticationPolicyAssert) HasCreatedOnNotEmpty() *AuthenticationPolicyAssert {
	a.AddAssertion(func(t *testing.T, o *sdk.AuthenticationPolicy) error {
		t.Helper()
		if o.CreatedOn == (time.Time{}) {
			return fmt.Errorf("expected created_on to be not empty")
		}
		return nil
	})
	return a
}
