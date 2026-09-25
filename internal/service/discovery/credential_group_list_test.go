package discovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCredentialGroupList(t *testing.T) {
	resourceType := "infoblox_credential_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckCredentialGroupExistsNIOS,
			Destroy: testAccCheckCredentialGroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "discovery/credential_group/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
