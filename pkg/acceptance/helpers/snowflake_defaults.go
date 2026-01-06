package helpers

import (
	"fmt"
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/testenvs"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

type SnowflakeDefaultsClient struct {
	context *TestClientContext
	ids     *IdsGenerator
}

func NewSnowflakeDefaultsClient(context *TestClientContext) *SnowflakeDefaultsClient {
	return &SnowflakeDefaultsClient{
		context: context,
	}
}

func (c *SnowflakeDefaultsClient) EnabledForSnowflakeOauthSecurityIntegration(t *testing.T) bool {
	t.Helper()
	if c.context.snowflakeEnvironment == testenvs.SnowflakeNonProdEnvironment {
		return true
	}
	return false
}

func (c *SnowflakeDefaultsClient) StageIdentifierOutputFormatForStreamOnDirectoryTable(t *testing.T, id sdk.SchemaObjectIdentifier) string {
	t.Helper()
	if c.context.snowflakeEnvironment == testenvs.SnowflakeNonProdEnvironment {
		return fmt.Sprintf(`"%s"."%s".%s`, id.DatabaseName(), id.SchemaName(), id.Name())
	}
	return id.Name()
}
