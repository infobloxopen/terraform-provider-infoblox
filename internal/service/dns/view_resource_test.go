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

func TestAccViewResource(t *testing.T) {
	resourceType := "infoblox_view"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckViewExistsNIOS,
			Destroy:    testAccCheckViewDestroyNIOS,
			Disappears: testAccCheckViewDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckViewExistsUDDI,
			Destroy:    testAccCheckViewDestroyUDDI,
			Disappears: testAccCheckViewDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/view/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckViewExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DNSAPI.ViewAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read View: %w", err)
		}
		if res == nil {
			return fmt.Errorf("View not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckViewExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSConfigurationAPI.ViewAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read View: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("View not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckViewDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DNSAPI.ViewAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("View still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckViewDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DNSAPI.ViewAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete View: %w", err)
		}
		return nil
	}
}

func testAccCheckViewDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSConfigurationAPI.ViewAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("View still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckViewDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSConfigurationAPI.ViewAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete View: %w", err)
		}
		return nil
	}
}
