package clouddiscovery_test

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

func TestAccCloudDiscoveryProviderResource(t *testing.T) {
	resourceType := "infoblox_cloud_discovery_provider"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:     testAccCheckCloudDiscoveryProviderExistsUDDI,
			Destroy:    testAccCheckCloudDiscoveryProviderDestroyUDDI,
			Disappears: testAccCheckCloudDiscoveryProviderDisappearsUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "clouddiscovery/cloud_discovery_provider/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckCloudDiscoveryProviderExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read CloudDiscoveryProvider: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("CloudDiscoveryProvider not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckCloudDiscoveryProviderDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("CloudDiscoveryProvider still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckCloudDiscoveryProviderDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete CloudDiscoveryProvider: %w", err)
		}
		return nil
	}
}
