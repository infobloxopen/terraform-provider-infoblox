package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpamHostDataSource(t *testing.T) {
	dsType := "infoblox_ipam_host"
	resourceType := "infoblox_ipam_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckIpamHostExistsUDDI,
			Destroy: testAccCheckIpamHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/ipam_host/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
