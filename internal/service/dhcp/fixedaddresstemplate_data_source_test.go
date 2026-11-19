package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFixedaddresstemplateDataSource(t *testing.T) {
	dsType := "infoblox_fixed_address_template"
	resourceType := "infoblox_fixed_address_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFixedaddresstemplateExistsNIOS,
			Destroy: testAccCheckFixedaddresstemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/fixed_address_template/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
