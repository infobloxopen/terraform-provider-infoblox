package cloud_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAwsuserDataSource(t *testing.T) {
	dsType := "infoblox_aws_user"
	resourceType := "infoblox_aws_user"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckAwsuserExistsNIOS,
			Destroy: testAccCheckAwsuserDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "cloud/aws_user/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
