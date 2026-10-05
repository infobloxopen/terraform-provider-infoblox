package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCategoryFilterDataSource(t *testing.T) {
	dsType := "infoblox_category_filter"
	resourceType := "infoblox_category_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCategoryFilterExistsUDDI,
			Destroy: testAccCheckCategoryFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/category_filter/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
