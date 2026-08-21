package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpHostList(t *testing.T) {
	resourceType := "infoblox_dhcp_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDhcpHostExistsUDDI,
			Destroy: testAccCheckDhcpHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/dhcp_host/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
