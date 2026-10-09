package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSuperhostList(t *testing.T) {
	resourceType := "infoblox_superhost"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSuperhostExistsNIOS,
			Destroy: testAccCheckSuperhostDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/superhost/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
