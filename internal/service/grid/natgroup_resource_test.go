package grid_test

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

func TestAccNatgroupResource(t *testing.T) {
	resourceType := "infoblox_natgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckNatgroupExistsNIOS,
			Destroy:    testAccCheckNatgroupDestroyNIOS,
			Disappears: testAccCheckNatgroupDisappearsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "grid/natgroup/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckNatgroupExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.GridAPI.NatgroupAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Natgroup: %w", err)
		}
		if res == nil {
			return fmt.Errorf("Natgroup not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNatgroupDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.GridAPI.NatgroupAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("Natgroup still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNatgroupDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.GridAPI.NatgroupAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete Natgroup: %w", err)
		}
		return nil
	}
}
