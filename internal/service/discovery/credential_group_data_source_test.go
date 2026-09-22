package discovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCredentialGroupDataSource(t *testing.T) {
	dsType := "infoblox_credential_group"
	resourceType := "infoblox_credential_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckCredentialGroupExistsNIOS,
			Destroy: testAccCheckCredentialGroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "discovery/credential_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
