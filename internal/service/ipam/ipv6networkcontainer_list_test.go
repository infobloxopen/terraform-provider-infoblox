package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6networkcontainerList(t *testing.T) {
	resourceType := "infoblox_ipv6_network_container"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6networkcontainerExistsNIOS,
			Destroy: testAccCheckIpv6networkcontainerDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6networkcontainerExistsUDDI,
			Destroy: testAccCheckIpv6networkcontainerDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/ipv6_network_container/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
