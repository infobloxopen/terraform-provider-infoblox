package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAccessCodeDataSource(t *testing.T) {
	dsType := "infoblox_access_code"
	resourceType := "infoblox_access_code"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAccessCodeExistsUDDI,
			Destroy: testAccCheckAccessCodeDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/access_code/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
