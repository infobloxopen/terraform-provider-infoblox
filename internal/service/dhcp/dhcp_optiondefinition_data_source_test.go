package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpOptiondefinitionDataSource(t *testing.T) {
	dsType := "infoblox_dhcp_optiondefinition"
	resourceType := "infoblox_dhcp_optiondefinition"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDhcpOptiondefinitionExistsNIOS,
			Destroy: testAccCheckDhcpOptiondefinitionDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDhcpOptiondefinitionExistsUDDI,
			Destroy: testAccCheckDhcpOptiondefinitionDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/dhcp_optiondefinition/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
