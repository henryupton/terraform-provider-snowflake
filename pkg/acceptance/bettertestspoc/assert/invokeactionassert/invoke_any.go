package invokeactionassert

import (
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type genericInvocation struct {
	f func() error
}

func (s *genericInvocation) ToTerraformTestCheckFunc(t *testing.T, _ *helpers.TestClient) resource.TestCheckFunc {
	t.Helper()
	return func(_ *terraform.State) error {
		return s.f()
	}
}

func Invoke(t *testing.T, f func() error) assert.TestCheckFuncProvider {
	t.Helper()
	return &genericInvocation{
		f: f,
	}
}
