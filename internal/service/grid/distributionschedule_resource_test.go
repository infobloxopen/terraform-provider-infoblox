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
	resourceType := "infoblox_distributionschedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDistributionscheduleExistsNIOS,
			Destroy: testAccCheckDistributionscheduleDestroyNIOS,
			// Disappears: testAccCheckDistributionscheduleDisappearsNIOS, // TODO: codegen bug, re-enable after fix
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "grid/distributionschedule/"+backend+"_resources.hcl", checksByBackend)
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

func testAccCheckDistributionscheduleDestroyNIOS(resourceType string) resource.TestCheckFunc {
	// Distributionschedule is a grid singleton (noDelete): destroy only drops it from state,
	// so the object always still exists afterwards and there is nothing to verify.
	return func(s *terraform.State) error {
		return nil
	}
}

// func testAccCheckDistributionscheduleDisappearsNIOS(resourceName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[resourceName]
// 		if !ok {
// 			return fmt.Errorf("not found: %s", resourceName)
// 		}
// 		conn := acctest.NIOSClient
// 		_, err := conn.GridAPI.DistributionscheduleAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
// 		if err != nil {
// 			return fmt.Errorf("failed to delete Distributionschedule: %w", err)
// 		}
// 		return nil
// 	}
// }
