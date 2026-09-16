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

func TestAccNsgroupStubmemberResource(t *testing.T) {
	resourceType := "infoblox_nsgroup_stubmember"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckNsgroupStubmemberExistsNIOS,
			Destroy:    testAccCheckNsgroupStubmemberDestroyNIOS,
			Disappears: testAccCheckNsgroupStubmemberDisappearsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/nsgroup_stubmember/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckNsgroupStubmemberExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DNSAPI.NsgroupStubmemberAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read NsgroupStubmember: %w", err)
		}
		if res == nil {
			return fmt.Errorf("NsgroupStubmember not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNsgroupStubmemberDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DNSAPI.NsgroupStubmemberAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("NsgroupStubmember still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNsgroupStubmemberDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DNSAPI.NsgroupStubmemberAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete NsgroupStubmember: %w", err)
		}
		return nil
	}
}
