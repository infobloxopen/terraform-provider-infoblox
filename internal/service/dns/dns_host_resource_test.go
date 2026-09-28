package dns_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDnsHostResource(t *testing.T) {
	resourceType := "infoblox_dns_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckDnsHostExistsUDDI,
			Destroy: testAccCheckDnsHostDestroyUDDI,
			//Disappears: testAccCheckDnsHostDisappearsUDDI,
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

func testAccCheckDnsHostDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSConfigurationAPI.HostAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DnsHost still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

//func testAccCheckDnsHostDisappearsUDDI(resourceName string) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[resourceName]
//		if !ok {
//			return fmt.Errorf("not found: %s", resourceName)
//		}
//_, err := acctest.UDDIClient.DNSConfigurationAPI.HostAPI.(context.Background(), rs.Primary.ID).Execute()
//		if err != nil {
//			return fmt.Errorf("failed to delete DnsHost: %w", err)
//		}
//		return nil
//	}
//}
