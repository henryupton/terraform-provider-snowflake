package model

import (
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/provider/resources"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"
)

func ListingWithStagedManifestWithLocation(
	resourceName string,
	name string,
	stageId sdk.SchemaObjectIdentifier,
	location string,
) *ListingModel {
	l := &ListingModel{ResourceModelMeta: config.Meta(resourceName, resources.Listing)}
	l.WithName(name)
	l.WithManifestValue(tfconfig.ListVariable(
		tfconfig.MapVariable(map[string]tfconfig.Variable{
			"from_stage": tfconfig.ListVariable(
				tfconfig.MapVariable(map[string]tfconfig.Variable{
					"stage":    tfconfig.StringVariable(stageId.FullyQualifiedName()),
					"location": tfconfig.StringVariable(location),
				}),
			),
		}),
	))
	return l
}

func ListingWithStagedManifestWithOptionals(
	resourceName string,
	name string,
	stageId sdk.SchemaObjectIdentifier,
	versionName string,
	versionComment string,
	location string,
) *ListingModel {
	l := &ListingModel{ResourceModelMeta: config.Meta(resourceName, resources.Listing)}
	l.WithName(name)
	l.WithManifestValue(tfconfig.ListVariable(
		tfconfig.MapVariable(map[string]tfconfig.Variable{
			"from_stage": tfconfig.ListVariable(
				tfconfig.MapVariable(map[string]tfconfig.Variable{
					"stage":           tfconfig.StringVariable(stageId.FullyQualifiedName()),
					"version_name":    tfconfig.StringVariable(versionName),
					"version_comment": tfconfig.StringVariable(versionComment),
					"location":        tfconfig.StringVariable(location),
				}),
			),
		}),
	))
	return l
}
