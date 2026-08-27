package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordNsList(t *testing.T) {
	resourceType := "infoblox_record_ns"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordNsExistsNIOS,
			Destroy: testAccCheckRecordNsDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordNsExistsUDDI,
			Destroy: testAccCheckRecordNsDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_ns/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
