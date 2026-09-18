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

func TestAccDtcMonitorSnmpResource(t *testing.T) {
	resourceType := "infoblox_dtc_monitor_snmp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckDtcMonitorSnmpExistsNIOS,
			Destroy:    testAccCheckDtcMonitorSnmpDestroyNIOS,
			Disappears: testAccCheckDtcMonitorSnmpDisappearsNIOS,
		},
		"uddi": {
			Exists:     testAccCheckDtcMonitorSnmpExistsUDDI,
			Destroy:    testAccCheckDtcMonitorSnmpDestroyUDDI,
			Disappears: testAccCheckDtcMonitorSnmpDisappearsUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dtc/dtc_monitor_snmp/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckDtcMonitorSnmpExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.DTCAPI.DtcMonitorSnmpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorSnmp: %w", err)
		}
		if res == nil {
			return fmt.Errorf("DtcMonitorSnmp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorSnmpExistsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		apiRes, _, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckSnmpAPI.Read(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to read DtcMonitorSnmp: %w", err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("DtcMonitorSnmp not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorSnmpDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.DTCAPI.DtcMonitorSnmpAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorSnmp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorSnmpDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.DTCAPI.DtcMonitorSnmpAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorSnmp: %w", err)
		}
		return nil
	}
}

func testAccCheckDtcMonitorSnmpDestroyUDDI(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckSnmpAPI.Read(context.Background(), rs.Primary.ID).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("DtcMonitorSnmp still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDtcMonitorSnmpDisappearsUDDI(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		_, err := acctest.UDDIClient.DNSTrafficControlAPI.HealthCheckSnmpAPI.Delete(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete DtcMonitorSnmp: %w", err)
		}
		return nil
	}
}
