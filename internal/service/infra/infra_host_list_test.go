package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInfraHostList(t *testing.T) {
	resourceType := "infoblox_infra_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInfraHostExistsUDDI,
			Destroy: testAccCheckInfraHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "infra/infra_host/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
