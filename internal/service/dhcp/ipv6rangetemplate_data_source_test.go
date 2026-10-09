package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6rangetemplateDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_range_template"
	resourceType := "infoblox_ipv6_range_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6rangetemplateExistsNIOS,
			Destroy: testAccCheckIpv6rangetemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_range_template/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
