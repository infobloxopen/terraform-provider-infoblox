package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInfraServiceDataSource(t *testing.T) {
	dsType := "infoblox_infra_service"
	resourceType := "infoblox_infra_service"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInfraServiceExistsUDDI,
			Destroy: testAccCheckInfraServiceDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "infra/infra_service/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
