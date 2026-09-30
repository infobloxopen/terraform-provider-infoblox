package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAdminuserDataSource(t *testing.T) {
	dsType := "infoblox_admin_user"
	resourceType := "infoblox_admin_user"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckAdminuserExistsNIOS,
			Destroy: testAccCheckAdminuserDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "security/admin_user/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
