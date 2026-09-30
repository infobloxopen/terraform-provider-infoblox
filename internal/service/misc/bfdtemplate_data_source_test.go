package misc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccBfdtemplateDataSource(t *testing.T) {
	dsType := "infoblox_bfd_template"
	resourceType := "infoblox_bfd_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckBfdtemplateExistsNIOS,
			Destroy: testAccCheckBfdtemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "misc/bfd_template/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
