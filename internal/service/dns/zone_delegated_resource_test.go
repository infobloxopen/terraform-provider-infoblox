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

func TestAccZoneDelegatedResource(t *testing.T) {
	resourceType := "infoblox_zone_delegated"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckZoneDelegatedExistsNIOS,
			Destroy:    testAccCheckZoneDelegatedDestroyNIOS,
			Disappears: testAccCheckZoneDelegatedDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckZoneDelegatedExistsUDDI,
			Destroy:    testAccCheckZoneDelegatedDestroyUDDI,
			Disappears: testAccCheckZoneDelegatedDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/zone_delegated/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckZoneDelegatedExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DNSAPI.ZoneDelegatedAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read ZoneDelegated: %w", err)
		}
		if res == nil {
			return fmt.Errorf("ZoneDelegated not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneDelegatedExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSConfigurationAPI.DelegationAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read ZoneDelegated: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("ZoneDelegated not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneDelegatedDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DNSAPI.ZoneDelegatedAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("ZoneDelegated still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneDelegatedDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DNSAPI.ZoneDelegatedAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete ZoneDelegated: %w", err)
		}
		return nil
	}
}

func testAccCheckZoneDelegatedDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSConfigurationAPI.DelegationAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("ZoneDelegated still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneDelegatedDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSConfigurationAPI.DelegationAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete ZoneDelegated: %w", err)
		}
		return nil
	}
}
