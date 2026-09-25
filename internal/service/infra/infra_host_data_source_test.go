package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInfraHostDataSource(t *testing.T) {
	dsType := "infoblox_infra_host"
	resourceType := "infoblox_infra_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInfraHostExistsUDDI,
			Destroy: testAccCheckInfraHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "infra/infra_host/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
