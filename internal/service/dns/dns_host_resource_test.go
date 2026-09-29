package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDnsHostResource(t *testing.T) {
	resourceType := "infoblox_dns_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists: testAccCheckDnsHostExistsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/dns_host/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDnsHostExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSConfigurationAPI.HostAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DnsHost: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DnsHost not found: %s", rs.Primary.ID)
		}
		return nil
	}
}
