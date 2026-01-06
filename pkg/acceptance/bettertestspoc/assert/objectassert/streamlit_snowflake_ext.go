package objectassert

import (
	"fmt"
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (s *StreamlitAssert) HasUrlIdNotEmpty() *StreamlitAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Streamlit) error {
		t.Helper()
		if o.UrlId == "" {
			return fmt.Errorf("expected url id to be not empty; got: %v", o.UrlId)
		}
		return nil
	})
	return s
}
