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

func TestAccHaGroupResource(t *testing.T) {
	resourceType := "infoblox_ha_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckHaGroupExistsUDDI,
			Destroy:    testAccCheckHaGroupDestroyUDDI,
			Disappears: testAccCheckHaGroupDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dhcp/ha_group/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckHaGroupExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.IPAddressManagementAPI.HaGroupAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read HaGroup: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("HaGroup not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckHaGroupDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.IPAddressManagementAPI.HaGroupAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("HaGroup still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckHaGroupDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.IPAddressManagementAPI.HaGroupAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete HaGroup: %w", err)
		}
		return nil
	}
}
