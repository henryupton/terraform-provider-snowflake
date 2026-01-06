//go:build exclude

package main

import (
	"strings"

	objectassertgen "github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert/objectassert/gen"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert/resourceshowoutputassert/gen"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/internal/genhelpers"
)

const (
	name    = "resource show output assertions"
	version = "0.1.0"
)

func main() {
	genhelpers.NewGenerator(
		genhelpers.NewPreambleModel(name, version).
			WithImport("testing").
			WithImport("github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert").
			WithImport("github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"),
		objectassertgen.GetSdkObjectDetails,
		gen.ModelFromSdkObjectDetails,
		getFilename,
		gen.AllTemplates,
	).
		RunAndHandleOsReturn()
}

func getFilename(_ genhelpers.SdkObjectDetails, model gen.ResourceShowOutputAssertionsModel) string {
	if model.IsDescribeOutput {
		return strings.TrimSuffix(genhelpers.ToSnakeCase(model.Name), "_details") + "_desc_output" + "_gen.go"
	}
	return genhelpers.ToSnakeCase(model.Name) + "_show_output" + "_gen.go"
}
