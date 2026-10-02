package anycast_test

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAnycastConfigResource(t *testing.T) {
	resourceType := "infoblox_anycast_config"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckAnycastConfigExistsUDDI,
			Destroy:    testAccCheckAnycastConfigDestroyUDDI,
			Disappears: testAccCheckAnycastConfigDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "anycast/anycast_config/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckAnycastConfigExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid AnycastConfig ID %q: %w", rs.Primary.ID, err)
		}
		apiRes, _, err := acctest.UDDIClient.AnycastAPI.OnPremAnycastManagerAPI.GetAnycastConfig(context.Background(), int64(uddiID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read AnycastConfig: %w", err)
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("AnycastConfig not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAnycastConfigDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid AnycastConfig ID %q: %w", rs.Primary.ID, err)
			}
			_, httpRes, err := acctest.UDDIClient.AnycastAPI.OnPremAnycastManagerAPI.GetAnycastConfig(context.Background(), int64(uddiID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("AnycastConfig still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAnycastConfigDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid AnycastConfig ID %q: %w", rs.Primary.ID, err)
		}
		_, _, err = acctest.UDDIClient.AnycastAPI.OnPremAnycastManagerAPI.DeleteAnycastConfig(context.Background(), int64(uddiID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete AnycastConfig: %w", err)
		}
		return nil
	}
}
