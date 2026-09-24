package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccApplicationFilterDataSource(t *testing.T) {
	dsType := "infoblox_application_filter"
	resourceType := "infoblox_application_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckApplicationFilterExistsUDDI,
			Destroy: testAccCheckApplicationFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/application_filter/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
