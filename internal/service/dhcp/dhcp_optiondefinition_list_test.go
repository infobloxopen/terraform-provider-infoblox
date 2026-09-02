package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpOptiondefinitionList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "dhcp/dhcp_optiondefinition/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
