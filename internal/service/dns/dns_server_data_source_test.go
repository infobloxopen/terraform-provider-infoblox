package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDnsServerDataSource(t *testing.T) {
	dsType := "infoblox_dns_server"
	resourceType := "infoblox_dns_server"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDnsServerExistsUDDI,
			Destroy: testAccCheckDnsServerDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/dns_server/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
