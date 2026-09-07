package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRangetemplateDataSource(t *testing.T) {
	dsType := "infoblox_rangetemplate"
	resourceType := "infoblox_rangetemplate"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRangetemplateExistsNIOS,
			Destroy: testAccCheckRangetemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/rangetemplate/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
