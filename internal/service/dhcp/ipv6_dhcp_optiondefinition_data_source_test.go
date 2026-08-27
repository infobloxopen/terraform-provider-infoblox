package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6DhcpOptiondefinitionDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_dhcp_optiondefinition"
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
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_dhcp_optiondefinition/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
