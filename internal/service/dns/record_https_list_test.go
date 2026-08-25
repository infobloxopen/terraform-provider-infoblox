package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordHttpsList(t *testing.T) {
	resourceType := "infoblox_record_https"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckRecordHttpsExistsUDDI,
			Destroy: testAccCheckRecordHttpsDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_https/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
