package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFixedaddresstemplateList(t *testing.T) {
	resourceType := "infoblox_fixed_address_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFixedaddresstemplateExistsNIOS,
			Destroy: testAccCheckFixedaddresstemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/fixed_address_template/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
