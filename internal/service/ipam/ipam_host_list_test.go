package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpamHostList(t *testing.T) {
	resourceType := "infoblox_ipam_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckIpamHostExistsUDDI,
			Destroy: testAccCheckIpamHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/ipam_host/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
