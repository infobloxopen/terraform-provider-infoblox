package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpHostDataSource(t *testing.T) {
	dsType := "infoblox_dhcp_host"
	resourceType := "infoblox_dhcp_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDhcpHostExistsUDDI,
			Destroy: testAccCheckDhcpHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/dhcp_host/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
