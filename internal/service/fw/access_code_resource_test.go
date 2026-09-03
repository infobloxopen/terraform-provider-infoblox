package fw_test

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

func TestAccAccessCodeResource(t *testing.T) {
	resourceType := "infoblox_access_code"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckAccessCodeExistsUDDI,
			Destroy:    testAccCheckAccessCodeDestroyUDDI,
			Disappears: testAccCheckAccessCodeDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "fw/access_code/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckAccessCodeExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.FWAPI.AccessCodesAPI.ReadAccessCode(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read AccessCode: %w", err)
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("AccessCode not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAccessCodeDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.FWAPI.AccessCodesAPI.ReadAccessCode(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("AccessCode still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAccessCodeDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.FWAPI.AccessCodesAPI.DeleteSingleAccessCodes(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete AccessCode: %w", err)
		}
		return nil
	}
}
