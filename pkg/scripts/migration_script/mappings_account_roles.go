package main

import (
	accconfig "github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/model"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/provider/resources"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func HandleAccountRoles(config *Config, csvInput [][]string) (string, error) {
	return HandleResources[AccountRoleCsvRow, AccountRoleRepresentation](config, csvInput, MapAccountRoleToModel)
}

func MapAccountRoleToModel(role AccountRoleRepresentation) (accconfig.ResourceModel, *ImportModel, error) {
	roleId := sdk.NewAccountObjectIdentifier(role.Name)
	resourceId := ResourceId(resources.AccountRole, roleId.FullyQualifiedName())
	resourceModel := model.AccountRole(resourceId, role.Name)

	handleIfNotEmpty(role.Comment, resourceModel.WithComment)

	importModel := NewImportModel(
		resourceModel.ResourceReference(),
		roleId.FullyQualifiedName(),
	)

	return resourceModel, importModel, nil
}
