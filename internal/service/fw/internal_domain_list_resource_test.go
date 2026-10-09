package fw_test

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

func TestAccInternalDomainListResource(t *testing.T) {
	resourceType := "infoblox_internal_domain_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckInternalDomainListExistsUDDI,
			Destroy:    testAccCheckInternalDomainListDestroyUDDI,
			Disappears: testAccCheckInternalDomainListDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "fw/internal_domain_list/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckInternalDomainListExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid InternalDomainList ID %q: %w", rs.Primary.ID, err)
		}
		apiRes, _, err := acctest.UDDIClient.FWAPI.InternalDomainListsAPI.ReadInternalDomains(context.Background(), int32(uddiID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read InternalDomainList: %w", err)
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("InternalDomainList not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckInternalDomainListDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 32)
			if err != nil {
				return fmt.Errorf("invalid InternalDomainList ID %q: %w", rs.Primary.ID, err)
			}
			_, httpRes, err := acctest.UDDIClient.FWAPI.InternalDomainListsAPI.ReadInternalDomains(context.Background(), int32(uddiID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("InternalDomainList still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckInternalDomainListDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		uddiID, err := strconv.ParseInt(rs.Primary.ID, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid InternalDomainList ID %q: %w", rs.Primary.ID, err)
		}
		_, err = acctest.UDDIClient.FWAPI.InternalDomainListsAPI.DeleteSingleInternalDomains(context.Background(), int32(uddiID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete InternalDomainList: %w", err)
		}
		return nil
	}
}
