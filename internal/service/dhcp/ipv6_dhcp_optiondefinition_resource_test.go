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

func TestAccIpv6DhcpOptiondefinitionResource(t *testing.T) {
	resourceType := "infoblox_ipv6_dhcp_optiondefinition"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckIpv6DhcpOptiondefinitionExistsNIOS,
			Destroy:    testAccCheckIpv6DhcpOptiondefinitionDestroyNIOS,
			Disappears: testAccCheckIpv6DhcpOptiondefinitionDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckIpv6DhcpOptiondefinitionExistsUDDI,
			Destroy:    testAccCheckIpv6DhcpOptiondefinitionDestroyUDDI,
			Disappears: testAccCheckIpv6DhcpOptiondefinitionDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dhcp/ipv6_dhcp_optiondefinition/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckIpv6DhcpOptiondefinitionExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DHCPAPI.Ipv6dhcpoptiondefinitionAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Ipv6DhcpOptiondefinition: %w", err)
		}
		if res == nil {
			return fmt.Errorf("Ipv6DhcpOptiondefinition not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6DhcpOptiondefinitionExistsUDDI(resourceName string) resource.TestCheckFunc {
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
			return fmt.Errorf("failed to read Ipv6DhcpOptiondefinition: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("Ipv6DhcpOptiondefinition not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6DhcpOptiondefinitionDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DHCPAPI.Ipv6dhcpoptiondefinitionAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("Ipv6DhcpOptiondefinition still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6DhcpOptiondefinitionDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DHCPAPI.Ipv6dhcpoptiondefinitionAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Ipv6DhcpOptiondefinition: %w", err)
		}
		return nil
	}
}

func testAccCheckIpv6DhcpOptiondefinitionDestroyUDDI(resourceType string) resource.TestCheckFunc {
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
			return fmt.Errorf("Ipv6DhcpOptiondefinition still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6DhcpOptiondefinitionDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.IPAddressManagementAPI.OptionCodeAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Ipv6DhcpOptiondefinition: %w", err)
		}
		return nil
	}
}
