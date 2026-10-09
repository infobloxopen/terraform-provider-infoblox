package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6fixedaddresstemplateList(t *testing.T) {
	resourceType := "infoblox_ipv6_fixed_address_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6fixedaddresstemplateExistsNIOS,
			Destroy: testAccCheckIpv6fixedaddresstemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/ipv6_fixed_address_template/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
