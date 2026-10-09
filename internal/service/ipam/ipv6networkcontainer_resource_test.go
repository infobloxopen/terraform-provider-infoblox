package ipam_test

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

func TestAccIpv6networkcontainerResource(t *testing.T) {
	resourceType := "infoblox_ipv6_network_container"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckIpv6networkcontainerExistsNIOS,
			Destroy:    testAccCheckIpv6networkcontainerDestroyNIOS,
			Disappears: testAccCheckIpv6networkcontainerDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckIpv6networkcontainerExistsUDDI,
			Destroy:    testAccCheckIpv6networkcontainerDestroyUDDI,
			Disappears: testAccCheckIpv6networkcontainerDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "ipam/ipv6_network_container/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckIpv6networkcontainerExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.IPAMAPI.Ipv6networkcontainerAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Ipv6networkcontainer: %w", err)
		}
		if res == nil {
			return fmt.Errorf("Ipv6networkcontainer not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6networkcontainerExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.IPAddressManagementAPI.AddressBlockAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Ipv6networkcontainer: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("Ipv6networkcontainer not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6networkcontainerDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.IPAMAPI.Ipv6networkcontainerAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("Ipv6networkcontainer still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6networkcontainerDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.IPAMAPI.Ipv6networkcontainerAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Ipv6networkcontainer: %w", err)
		}
		return nil
	}
}

func testAccCheckIpv6networkcontainerDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.IPAddressManagementAPI.AddressBlockAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("Ipv6networkcontainer still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6networkcontainerDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.IPAddressManagementAPI.AddressBlockAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Ipv6networkcontainer: %w", err)
		}
		return nil
	}
}
