package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6fixedaddresstemplateDataSource(t *testing.T) {
	dsType := "infoblox_ipv6fixedaddresstemplate"
	resourceType := "infoblox_ipv6fixedaddresstemplate"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6fixedaddresstemplateExistsNIOS,
			Destroy: testAccCheckIpv6fixedaddresstemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6fixedaddresstemplate/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
