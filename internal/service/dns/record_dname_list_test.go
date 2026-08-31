package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordDnameList(t *testing.T) {
	resourceType := "infoblox_record_dname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordDnameExistsNIOS,
			Destroy: testAccCheckRecordDnameDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordDnameExistsUDDI,
			Destroy: testAccCheckRecordDnameDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_dname/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
