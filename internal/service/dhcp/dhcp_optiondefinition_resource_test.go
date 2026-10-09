package dhcp_test

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

func TestAccDhcpOptiondefinitionResource(t *testing.T) {
	resourceType := "infoblox_dhcp_optiondefinition"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckDhcpOptiondefinitionExistsNIOS,
			Destroy:    testAccCheckDhcpOptiondefinitionDestroyNIOS,
			Disappears: testAccCheckDhcpOptiondefinitionDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckDhcpOptiondefinitionExistsUDDI,
			Destroy:    testAccCheckDhcpOptiondefinitionDestroyUDDI,
			Disappears: testAccCheckDhcpOptiondefinitionDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dhcp/dhcp_optiondefinition/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDhcpOptiondefinitionExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DHCPAPI.DhcpoptiondefinitionAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DhcpOptiondefinition: %w", err)
		}
		if res == nil {
			return fmt.Errorf("DhcpOptiondefinition not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDhcpOptiondefinitionExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.IPAddressManagementAPI.OptionCodeAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DhcpOptiondefinition: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DhcpOptiondefinition not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDhcpOptiondefinitionDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DHCPAPI.DhcpoptiondefinitionAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DhcpOptiondefinition still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDhcpOptiondefinitionDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DHCPAPI.DhcpoptiondefinitionAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DhcpOptiondefinition: %w", err)
		}
		return nil
	}
}

func testAccCheckDhcpOptiondefinitionDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.IPAddressManagementAPI.OptionCodeAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DhcpOptiondefinition still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDhcpOptiondefinitionDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.IPAddressManagementAPI.OptionCodeAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DhcpOptiondefinition: %w", err)
		}
		return nil
	}
}
