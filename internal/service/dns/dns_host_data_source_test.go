package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDnsHostDataSource(t *testing.T) {
	dsType := "infoblox_dns_host"
	resourceType := "infoblox_dns_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDnsHostExistsUDDI,
			Destroy: testAccCheckDnsHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/dns_host/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
