//go:build exclude

package main

import (
	resourcemodelgen "github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/model/gen"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/datasourcemodel/gen"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/internal/genhelpers"
)

const (
	name    = "data source model builder"
	version = "0.1.0"
)

func main() {
	genhelpers.NewGenerator(
		genhelpers.NewPreambleModel(name, version).
			WithImport("encoding/json").
			WithImport("github.com/henryupton/terraform-provider-snowflakier/pkg/provider/datasources").
			WithImport("github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config").
			WithNamedImport("tfconfig", "github.com/hashicorp/terraform-plugin-testing/config"),
		gen.GetDatasourceSchemaDetails,
		resourcemodelgen.ModelFromResourceSchemaDetails,
		getFilename,
		gen.AllTemplates,
	).
		RunAndHandleOsReturn()
}

func getFilename(_ genhelpers.ResourceSchemaDetails, model resourcemodelgen.ResourceConfigBuilderModel) string {
	return genhelpers.ToSnakeCase(model.Name) + "_model" + "_gen.go"
}
