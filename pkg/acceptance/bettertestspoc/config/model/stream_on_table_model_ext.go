package model

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func StreamOnTableBase(resourceName string, id, tableId sdk.SchemaObjectIdentifier) *StreamOnTableModel {
	return StreamOnTable(resourceName, id.DatabaseName(), id.SchemaName(), id.Name(), tableId.FullyQualifiedName())
}
