package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccOptionGroupDataSource(t *testing.T) {
	dsType := "infoblox_option_group"
	resourceType := "infoblox_option_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckOptionGroupExistsUDDI,
			Destroy: testAccCheckOptionGroupDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/option_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
