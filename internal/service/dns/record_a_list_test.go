package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordAList(t *testing.T) {
	resourceType := "infoblox_record_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAExistsNIOS,
			Destroy: testAccCheckRecordADestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordAExistsUDDI,
			Destroy: testAccCheckRecordADestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_a/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
