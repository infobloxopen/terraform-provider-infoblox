package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSecurityPolicyDataSource(t *testing.T) {
	dsType := "infoblox_security_policy"
	resourceType := "infoblox_security_policy"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckSecurityPolicyExistsUDDI,
			Destroy: testAccCheckSecurityPolicyDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/security_policy/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
