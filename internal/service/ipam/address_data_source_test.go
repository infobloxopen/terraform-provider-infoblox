package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-unified/internal/acctest"
)

func TestAccAddressDataSource(t *testing.T) {
	dsType := "unified_ipam_address"
	resourceType := "unified_ipam_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAddressExistsUDDI,
			Destroy: testAccCheckAddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/address/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
