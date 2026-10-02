package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDistributionscheduleResource(t *testing.T) {
	resourceType := "infoblox_distribution_schedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists: testAccCheckDistributionscheduleExistsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "grid/distribution_schedule/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDistributionscheduleExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.GridAPI.DistributionscheduleAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read Distributionschedule: %w", err)
		}
		if res == nil {
			return fmt.Errorf("Distributionschedule not found: %s", rs.Primary.ID)
		}
		return nil
	}
}
