package cloud_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAwsuserList(t *testing.T) {
	resourceType := "infoblox_aws_user"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckAwsuserExistsNIOS,
			Destroy: testAccCheckAwsuserDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "cloud/aws_user/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
