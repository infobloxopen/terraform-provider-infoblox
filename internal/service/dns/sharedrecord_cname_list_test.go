package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordCnameList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordCnameExistsNIOS,
			Destroy: testAccCheckSharedrecordCnameDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_cname/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
