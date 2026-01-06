package model

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func StreamOnExternalTableBase(resourceName string, id, externalTableId sdk.SchemaObjectIdentifier) *StreamOnExternalTableModel {
	return StreamOnExternalTable(resourceName, id.DatabaseName(), id.SchemaName(), id.Name(), externalTableId.FullyQualifiedName()).WithInsertOnly("true")
}
