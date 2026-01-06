package testacc

import (
	"fmt"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func shareConfigOneAccount(shareId sdk.AccountObjectIdentifier, comment string, account string) string {
	return fmt.Sprintf(`
resource "snowflake_share" "test" {
	name           = "%s"
	comment        = "%s"
	accounts       = ["%s"]
}
`, shareId.Name(), comment, account)
}
