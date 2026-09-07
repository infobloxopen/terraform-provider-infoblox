package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6fixedaddressList(t *testing.T) {
	resourceType := "infoblox_ipv6_fixed_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6fixedaddressExistsNIOS,
			Destroy: testAccCheckIpv6fixedaddressDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6fixedaddressExistsUDDI,
			Destroy: testAccCheckIpv6fixedaddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/ipv6_fixed_address/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
