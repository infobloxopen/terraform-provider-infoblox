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

func TestAccIpv6fixedaddresstemplateResource(t *testing.T) {
	resourceType := "infoblox_ipv6_fixed_address_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckIpv6fixedaddresstemplateExistsNIOS,
			Destroy:    testAccCheckIpv6fixedaddresstemplateDestroyNIOS,
			Disappears: testAccCheckIpv6fixedaddresstemplateDisappearsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dhcp/ipv6_fixed_address_template/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckIpv6fixedaddresstemplateExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DHCPAPI.Ipv6fixedaddresstemplateAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Ipv6fixedaddresstemplate: %w", err)
		}
		if res == nil {
			return fmt.Errorf("Ipv6fixedaddresstemplate not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6fixedaddresstemplateDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DHCPAPI.Ipv6fixedaddresstemplateAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("Ipv6fixedaddresstemplate still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpv6fixedaddresstemplateDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DHCPAPI.Ipv6fixedaddresstemplateAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Ipv6fixedaddresstemplate: %w", err)
		}
		return nil
	}
}
