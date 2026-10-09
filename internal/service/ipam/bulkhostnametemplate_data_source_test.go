package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccBulkhostnametemplateDataSource(t *testing.T) {
	dsType := "infoblox_bulk_hostname_template"
	resourceType := "infoblox_bulk_hostname_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckBulkhostnametemplateExistsNIOS,
			Destroy: testAccCheckBulkhostnametemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/bulk_hostname_template/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
