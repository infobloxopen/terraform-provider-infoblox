package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDnsServerList(t *testing.T) {
	resourceType := "infoblox_dns_server"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDnsServerExistsUDDI,
			Destroy: testAccCheckDnsServerDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/dns_server/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
