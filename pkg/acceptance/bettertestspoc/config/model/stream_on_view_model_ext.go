package model

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func StreamOnViewBase(resourceName string, id sdk.SchemaObjectIdentifier, viewId sdk.SchemaObjectIdentifier) *StreamOnViewModel {
	return StreamOnView(resourceName, id.DatabaseName(), id.SchemaName(), id.Name(), viewId.FullyQualifiedName())
}
