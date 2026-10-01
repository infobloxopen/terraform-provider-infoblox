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

func TestAccIpamHostResource(t *testing.T) {
	resourceType := "infoblox_ipam_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckIpamHostExistsUDDI,
			Destroy:    testAccCheckIpamHostDestroyUDDI,
			Disappears: testAccCheckIpamHostDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "ipam/ipam_host/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckIpamHostExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.IPAddressManagementAPI.IpamHostAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read IpamHost: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("IpamHost not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpamHostDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.IPAddressManagementAPI.IpamHostAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("IpamHost still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIpamHostDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.IPAddressManagementAPI.IpamHostAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete IpamHost: %w", err)
		}
		return nil
	}
}
