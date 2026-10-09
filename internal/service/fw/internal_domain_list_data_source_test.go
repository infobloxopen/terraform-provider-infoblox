package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccInternalDomainListDataSource(t *testing.T) {
	dsType := "infoblox_internal_domain_list"
	resourceType := "infoblox_internal_domain_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckInternalDomainListExistsUDDI,
			Destroy: testAccCheckInternalDomainListDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/internal_domain_list/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
