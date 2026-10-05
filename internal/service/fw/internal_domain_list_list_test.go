package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInternalDomainListList(t *testing.T) {
	resourceType := "infoblox_internal_domain_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInternalDomainListExistsUDDI,
			Destroy: testAccCheckInternalDomainListDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/internal_domain_list/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
