package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6networkList(t *testing.T) {
	resourceType := "infoblox_ipv6_network"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6networkExistsNIOS,
			Destroy: testAccCheckIpv6networkDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6networkExistsUDDI,
			Destroy: testAccCheckIpv6networkDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/ipv6_network/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
