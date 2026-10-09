package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAuthNsgDataSource(t *testing.T) {
	dsType := "infoblox_auth_nsg"
	resourceType := "infoblox_auth_nsg"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAuthNsgExistsUDDI,
			Destroy: testAccCheckAuthNsgDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/auth_nsg/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
