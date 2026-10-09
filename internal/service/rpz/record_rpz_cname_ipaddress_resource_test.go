package rpz_test

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

func TestAccRecordRpzCnameIpaddressResource(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname_ipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckRecordRpzCnameIpaddressExistsNIOS,
			Destroy:    testAccCheckRecordRpzCnameIpaddressDestroyNIOS,
			Disappears: testAccCheckRecordRpzCnameIpaddressDisappearsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "rpz/record_rpz_cname_ipaddress/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckRecordRpzCnameIpaddressExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.RPZAPI.RecordRpzCnameIpaddressAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read RecordRpzCnameIpaddress: %w", err)
		}
		if res == nil {
			return fmt.Errorf("RecordRpzCnameIpaddress not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckRecordRpzCnameIpaddressDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.RPZAPI.RecordRpzCnameIpaddressAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("RecordRpzCnameIpaddress still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckRecordRpzCnameIpaddressDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.RPZAPI.RecordRpzCnameIpaddressAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete RecordRpzCnameIpaddress: %w", err)
		}
		return nil
	}
}
