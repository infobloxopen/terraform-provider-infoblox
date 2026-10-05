package rir_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRirOrganizationDataSource(t *testing.T) {
	dsType := "infoblox_rir_organization"
	resourceType := "infoblox_rir_organization"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRirOrganizationExistsNIOS,
			Destroy: testAccCheckRirOrganizationDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rir/rir_organization/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
