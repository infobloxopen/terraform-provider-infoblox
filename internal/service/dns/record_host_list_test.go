package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordHostList(t *testing.T) {
	resourceType := "infoblox_record_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordHostExistsNIOS,
			Destroy: testAccCheckRecordHostDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_host/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
