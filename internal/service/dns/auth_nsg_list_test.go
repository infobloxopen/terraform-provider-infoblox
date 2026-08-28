package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAuthNsgList(t *testing.T) {
	resourceType := "infoblox_auth_nsg"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAuthNsgExistsUDDI,
			Destroy: testAccCheckAuthNsgDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/auth_nsg/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
