package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRangetemplateDataSource(t *testing.T) {
	dsType := "infoblox_range_template"
	resourceType := "infoblox_range_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRangetemplateExistsNIOS,
			Destroy: testAccCheckRangetemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/range_template/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
