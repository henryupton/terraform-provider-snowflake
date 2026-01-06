package resourceassert

import (
	"fmt"
	"strconv"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
	"github.com/henryupton/terraform-provider-snowflakier/pkg/sdk"
)

func (r *RowAccessPolicyResourceAssert) HasArguments(args []sdk.TableColumnSignature) *RowAccessPolicyResourceAssert {
	r.AddAssertion(assert.ValueSet("argument.#", strconv.FormatInt(int64(len(args)), 10)))
	for i, v := range args {
		r.AddAssertion(assert.ValueSet(fmt.Sprintf("argument.%d.name", i), v.Name))
		r.AddAssertion(assert.ValueSet(fmt.Sprintf("argument.%d.type", i), v.Type.ToSql()))
	}
	return r
}
