package dtc_test

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

func TestAccDtcMonitorIcmpResource(t *testing.T) {
	resourceType := "infoblox_dtc_monitor_icmp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckDtcMonitorIcmpExistsNIOS,
			Destroy:    testAccCheckDtcMonitorIcmpDestroyNIOS,
			Disappears: testAccCheckDtcMonitorIcmpDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckDtcMonitorIcmpExistsUDDI,
			Destroy:    testAccCheckDtcMonitorIcmpDestroyUDDI,
			Disappears: testAccCheckDtcMonitorIcmpDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dtc/dtc_monitor_icmp/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDtcMonitorIcmpExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DTCAPI.DtcMonitorIcmpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorIcmp: %w", err)
		}
		if res == nil {
			return fmt.Errorf("DtcMonitorIcmp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorIcmpExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorIcmp: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DtcMonitorIcmp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorIcmpDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DTCAPI.DtcMonitorIcmpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorIcmp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorIcmpDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DTCAPI.DtcMonitorIcmpAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorIcmp: %w", err)
		}
		return nil
	}
}

func testAccCheckDtcMonitorIcmpDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorIcmp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorIcmpDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorIcmp: %w", err)
		}
		return nil
	}
}
