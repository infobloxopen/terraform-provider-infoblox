package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDhcpHostResource(t *testing.T) {
	resourceType := "infoblox_dhcp_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckDhcpHostExistsUDDI,
			Destroy:    testAccCheckDhcpHostDestroyUDDI,
			Disappears: testAccCheckDhcpHostDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dhcp/dhcp_host/"+backend+"_resources.tfvars", checksByBackend)
		})
	}
}

func testAccCheckDhcpHostExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.IPAddressManagementAPI.DhcpHostAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DhcpHost: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DhcpHost not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDhcpHostDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// DhcpHost is a system-managed object that cannot be deleted via the API.
		// "destroy" only removes it from Terraform state; the host continues to exist on the platform.
		return nil
	}
}

func testAccCheckDhcpHostDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// DhcpHost cannot be deleted via the API; "disappears" testing is not applicable.
		return nil
	}
}
