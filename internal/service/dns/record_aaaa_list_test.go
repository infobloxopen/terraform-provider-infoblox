package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordAaaaList(t *testing.T) {
	resourceType := "infoblox_record_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAaaaExistsNIOS,
			Destroy: testAccCheckRecordAaaaDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordAaaaExistsUDDI,
			Destroy: testAccCheckRecordAaaaDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_aaaa/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
