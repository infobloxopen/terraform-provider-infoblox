package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSecurityPolicyList(t *testing.T) {
	resourceType := "infoblox_security_policy"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckSecurityPolicyExistsUDDI,
			Destroy: testAccCheckSecurityPolicyDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/security_policy/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
