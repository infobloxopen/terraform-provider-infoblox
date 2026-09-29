package discovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDiscoveryCredentialgroupList(t *testing.T) {
	resourceType := "infoblox_discovery_credential_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDiscoveryCredentialgroupExistsNIOS,
			Destroy: testAccCheckDiscoveryCredentialgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "discovery/discovery_credential_group/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
