package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpOptionspaceList(t *testing.T) {
	resourceType := "infoblox_dhcp_optionspace"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDhcpOptionspaceExistsNIOS,
			Destroy: testAccCheckDhcpOptionspaceDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDhcpOptionspaceExistsUDDI,
			Destroy: testAccCheckDhcpOptionspaceDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/dhcp_optionspace/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
