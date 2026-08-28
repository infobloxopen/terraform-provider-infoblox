package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordPtrList(t *testing.T) {
	resourceType := "infoblox_record_ptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordPtrExistsNIOS,
			Destroy: testAccCheckRecordPtrDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordPtrExistsUDDI,
			Destroy: testAccCheckRecordPtrDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_ptr/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
