package windows_test

import (
	"fmt"
	"io/fs"
	"regexp"
	"testing"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/datasourcemodel"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/config/providermodel"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/helpers/random"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/testenvs"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/testfiles"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/internal/oswrapper"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/internal/snowflakeenvs"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/manual_tests"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAcc_Provider_tomlConfigIsTooPermissive(t *testing.T) {
	_ = testenvs.GetOrSkipTest(t, testenvs.EnableManual)
	if !oswrapper.IsRunningOnWindows() {
		t.Skip("checking file permissions on other platforms is currently done in the provider package")
	}

	permissions := fs.FileMode(0o755)

	configPath := testfiles.TestFileWithCustomPermissions(t, random.AlphaN(10), random.Bytes(), permissions)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: manual_tests.ManualTestProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_5_0),
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Setenv(snowflakeenvs.ConfigPath, configPath)
				},
				Config:      config.FromModels(t, providermodel.SnowflakeProvider().WithProfile(configPath), datasourceModel()),
				ExpectError: regexp.MustCompile(fmt.Sprintf("could not load config file: config file %s has unsafe permissions - %#o", configPath, permissions)),
			},
		},
	})
}

func datasourceModel() config.DatasourceModel {
	return datasourcemodel.Database("t", "SNOWFLAKE")
}
