package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAdminuserList(t *testing.T) {
	resourceType := "infoblox_admin_user"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckAdminuserExistsNIOS,
			Destroy: testAccCheckAdminuserDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "security/admin_user/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
