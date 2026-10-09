package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInfraServiceList(t *testing.T) {
	resourceType := "infoblox_infra_service"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInfraServiceExistsUDDI,
			Destroy: testAccCheckInfraServiceDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "infra/infra_service/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
