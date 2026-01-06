package model

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"
)

func (a *AccountModel) WithAdminUserTypeEnum(adminUserType sdk.UserType) *AccountModel {
	a.AdminUserType = tfconfig.StringVariable(string(adminUserType))
	return a
}
