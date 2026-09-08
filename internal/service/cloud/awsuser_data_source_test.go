package cloud_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAwsuserDataSource(t *testing.T) {
	dsType := "infoblox_awsuser"
	resourceType := "infoblox_awsuser"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckAwsuserExistsNIOS,
			Destroy: testAccCheckAwsuserDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "cloud/awsuser/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
