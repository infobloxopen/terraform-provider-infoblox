package anycast_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAnycastConfigDataSource(t *testing.T) {
	dsType := "infoblox_anycast_config"
	resourceType := "infoblox_anycast_config"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAnycastConfigExistsUDDI,
			Destroy: testAccCheckAnycastConfigDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "anycast/anycast_config/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
