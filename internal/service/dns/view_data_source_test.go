package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccViewDataSource(t *testing.T) {
	dsType := "infoblox_view"
	resourceType := "infoblox_view"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckViewExistsNIOS,
			Destroy: testAccCheckViewDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckViewExistsUDDI,
			Destroy: testAccCheckViewDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/view/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
