package acl_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNamedaclList(t *testing.T) {
	resourceType := "infoblox_namedacl"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNamedaclExistsNIOS,
			Destroy: testAccCheckNamedaclDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckNamedaclExistsUDDI,
			Destroy: testAccCheckNamedaclDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "acl/namedacl/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
