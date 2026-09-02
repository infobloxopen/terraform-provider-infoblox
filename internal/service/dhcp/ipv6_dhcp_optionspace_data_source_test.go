package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6DhcpOptionspaceDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_dhcp_optionspace"
	resourceType := "infoblox_ipv6_dhcp_optionspace"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6DhcpOptionspaceExistsNIOS,
			Destroy: testAccCheckIpv6DhcpOptionspaceDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6DhcpOptionspaceExistsUDDI,
			Destroy: testAccCheckIpv6DhcpOptionspaceDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_dhcp_optionspace/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
