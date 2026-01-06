//go:build account_level_tests

package testacc

import (
	"fmt"
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/internal/tracking"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert/invokeactionassert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert/resourceassert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/model"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/providermodel"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/testprofiles"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/helpers"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/provider/experimentalfeatures"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAcc_Experimental_Warehouse_ShowImprovedPerformance(t *testing.T) {
	warehouseId := testClient().Ids.RandomAccountObjectIdentifier()

	providerModel := providermodel.SnowflakeProvider().WithProfile(testprofiles.Secondary).
		WithExperimentalFeaturesEnabled(experimentalfeatures.WarehouseShowImprovedPerformance)
	warehouseModel := model.Warehouse("test", warehouseId.Name())

	expectedWarehouseQuery := fmt.Sprintf("SHOW WAREHOUSES LIKE '%[1]s' STARTS WITH '%[1]s' LIMIT 1", warehouseId.Name())

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: providerFactoryUsingCache("TestAcc_Experimental_Warehouse_ShowImprovedPerformance"),
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_5_0),
		},
		Steps: []resource.TestStep{
			{
				Config: config.FromModels(t, providerModel, warehouseModel),
				Check: assertThat(t,
					resourceassert.WarehouseResource(t, warehouseModel.ResourceReference()).
						HasNameString(warehouseId.Name()),
					invokeactionassert.QueryHistoryEntry(t, secondaryTestClient(), expectedWarehouseQuery, tracking.CreateOperation, 100),
				),
			},
			{
				Config:       config.FromModels(t, providerModel, warehouseModel),
				ResourceName: warehouseModel.ResourceReference(),
				ImportState:  true,
				ImportStateCheck: assertThatImport(t,
					resourceassert.ImportedWarehouseResource(t, helpers.EncodeResourceIdentifier(warehouseId)).
						HasNameString(warehouseId.Name()),
					invokeactionassert.QueryHistoryEntryInImport(t, secondaryTestClient(), expectedWarehouseQuery, tracking.ImportOperation, 100),
				),
			},
		},
	})
}
