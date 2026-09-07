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

func TestAccDtcMonitorPdpResource(t *testing.T) {
	resourceType := "infoblox_dtc_monitor_pdp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckDtcMonitorPdpExistsNIOS,
			Destroy:    testAccCheckDtcMonitorPdpDestroyNIOS,
			Disappears: testAccCheckDtcMonitorPdpDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckDtcMonitorPdpExistsUDDI,
			Destroy:    testAccCheckDtcMonitorPdpDestroyUDDI,
			Disappears: testAccCheckDtcMonitorPdpDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dtc/dtc_monitor_pdp/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDtcMonitorPdpExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DTCAPI.DtcMonitorPdpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorPdp: %w", err)
		}
		if res == nil {
			return fmt.Errorf("DtcMonitorPdp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorPdpExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckPdpAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorPdp: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DtcMonitorPdp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorPdpDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DTCAPI.DtcMonitorPdpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorPdp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorPdpDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DTCAPI.DtcMonitorPdpAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorPdp: %w", err)
		}
		return nil
	}
}

func testAccCheckDtcMonitorPdpDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckPdpAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorPdp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorPdpDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckPdpAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorPdp: %w", err)
		}
		return nil
	}
}
