package anycast_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAnycastHostDataSource(t *testing.T) {
	dsType := "infoblox_anycast_host"
	resourceType := "infoblox_anycast_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAnycastHostExistsUDDI,
			Destroy: testAccCheckAnycastHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "anycast/anycast_host/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
