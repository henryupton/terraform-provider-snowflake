package resourceassert

import (
	"strconv"

	"github.com/henryupton/terraform-provider-snowflakier/pkg/acceptance/bettertestspoc/assert"
)

func (f *FunctionJavaResourceAssert) HasImportsLength(len int) *FunctionJavaResourceAssert {
	f.AddAssertion(assert.ValueSet("imports.#", strconv.FormatInt(int64(len), 10)))
	return f
}
