package discovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDiscoveryCredentialgroupDataSource(t *testing.T) {
	dsType := "infoblox_discovery_credential_group"
	resourceType := "infoblox_discovery_credential_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDiscoveryCredentialgroupExistsNIOS,
			Destroy: testAccCheckDiscoveryCredentialgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "discovery/discovery_credential_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
