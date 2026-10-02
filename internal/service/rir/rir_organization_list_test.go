package rir_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRirOrganizationList(t *testing.T) {
	resourceType := "infoblox_rir_organization"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRirOrganizationExistsNIOS,
			Destroy: testAccCheckRirOrganizationDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rir/rir_organization/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
