package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordCaaList(t *testing.T) {
	resourceType := "infoblox_record_caa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordCaaExistsNIOS,
			Destroy: testAccCheckRecordCaaDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordCaaExistsUDDI,
			Destroy: testAccCheckRecordCaaDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_caa/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
