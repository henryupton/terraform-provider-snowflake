package datasourcemodel

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (s *StreamsModel) WithInDatabase(databaseId sdk.AccountObjectIdentifier) *StreamsModel {
	return s.WithInValue(
		tfconfig.ObjectVariable(map[string]tfconfig.Variable{
			"database": tfconfig.StringVariable(databaseId.Name()),
		}),
	)
}
