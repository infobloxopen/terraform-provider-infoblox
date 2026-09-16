package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFilteroptionDataSource(t *testing.T) {
	dsType := "infoblox_filteroption"
	resourceType := "infoblox_filteroption"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFilteroptionExistsNIOS,
			Destroy: testAccCheckFilteroptionDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckFilteroptionExistsUDDI,
			Destroy: testAccCheckFilteroptionDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/filteroption/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
