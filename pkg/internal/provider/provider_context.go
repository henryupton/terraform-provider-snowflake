package provider

import "github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"

type Context struct {
	Client             *sdk.Client
	EnabledFeatures    []string
	EnabledExperiments []string
}
