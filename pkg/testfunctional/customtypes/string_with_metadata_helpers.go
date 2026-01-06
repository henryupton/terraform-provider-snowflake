package customtypes

import "github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"

func StringWithMetadataAttributeCreate(v StringWithMetadataValue, createField **string) {
	if !v.IsNull() {
		*createField = sdk.String(v.ValueString())
	}
}
