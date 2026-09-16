package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordMxList(t *testing.T) {
	resourceType := "infoblox_record_mx"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordMxExistsNIOS,
			Destroy: testAccCheckRecordMxDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordMxExistsUDDI,
			Destroy: testAccCheckRecordMxDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_mx/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
