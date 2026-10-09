package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordCnameList(t *testing.T) {
	resourceType := "infoblox_record_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordCnameExistsNIOS,
			Destroy: testAccCheckRecordCnameDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordCnameExistsUDDI,
			Destroy: testAccCheckRecordCnameDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_cname/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
