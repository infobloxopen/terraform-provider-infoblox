package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCategoryFilterList(t *testing.T) {
	resourceType := "infoblox_category_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCategoryFilterExistsUDDI,
			Destroy: testAccCheckCategoryFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/category_filter/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
