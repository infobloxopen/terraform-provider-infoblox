package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6DhcpOptiondefinitionList(t *testing.T) {
	resourceType := "infoblox_ipv6_dhcp_optiondefinition"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6DhcpOptiondefinitionExistsNIOS,
			Destroy: testAccCheckIpv6DhcpOptiondefinitionDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6DhcpOptiondefinitionExistsUDDI,
			Destroy: testAccCheckIpv6DhcpOptiondefinitionDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/ipv6_dhcp_optiondefinition/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
