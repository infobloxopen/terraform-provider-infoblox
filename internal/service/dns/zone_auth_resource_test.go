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

func TestAccZoneAuthResource(t *testing.T) {
	resourceType := "infoblox_zone_auth"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckZoneAuthExistsNIOS,
			Destroy:    testAccCheckZoneAuthDestroyNIOS,
			Disappears: testAccCheckZoneAuthDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckZoneAuthExistsUDDI,
			Destroy:    testAccCheckZoneAuthDestroyUDDI,
			Disappears: testAccCheckZoneAuthDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/zone_auth/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckZoneAuthExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DNSAPI.ZoneAuthAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read ZoneAuth: %w", err)
		}
		if res == nil {
			return fmt.Errorf("ZoneAuth not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneAuthExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSConfigurationAPI.AuthZoneAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read ZoneAuth: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("ZoneAuth not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneAuthDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DNSAPI.ZoneAuthAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("ZoneAuth still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneAuthDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DNSAPI.ZoneAuthAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete ZoneAuth: %w", err)
		}
		return nil
	}
}

func testAccCheckZoneAuthDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSConfigurationAPI.AuthZoneAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("ZoneAuth still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckZoneAuthDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSConfigurationAPI.AuthZoneAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete ZoneAuth: %w", err)
		}
		return nil
	}
}
