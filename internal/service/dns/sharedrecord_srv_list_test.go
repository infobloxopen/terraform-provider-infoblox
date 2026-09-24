package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordSrvList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_srv"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordSrvExistsNIOS,
			Destroy: testAccCheckSharedrecordSrvDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_srv/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
