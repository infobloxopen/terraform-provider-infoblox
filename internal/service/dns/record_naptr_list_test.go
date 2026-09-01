package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordNaptrList(t *testing.T) {
	resourceType := "infoblox_record_naptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordNaptrExistsNIOS,
			Destroy: testAccCheckRecordNaptrDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordNaptrExistsUDDI,
			Destroy: testAccCheckRecordNaptrDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_naptr/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
