package infoblox

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
)

var testResourceDtcMonitorIcmp1 = `
	resource "infoblox_dtc_monitor_icmp" "test-monitor_icmp1" {
		name = "test-monitor_icmp1"
	}
`
var testResourceDtcMonitorIcmp1_Rename = `
	resource "infoblox_dtc_monitor_icmp" "test-monitor_icmp1" {
		name = "test-monitor_icmp1_rename"
	}
`

var testResourceDtcMonitorIcmp2 = `
	resource "infoblox_dtc_monitor_icmp" "test-monitor_icmp2" {
		name       = "test-monitor_icmp2"
		comment    = "test dtc monitor_icmp with max params"
		ext_attrs  = jsonencode({ "Site" = "India" })
		interval   = 7
		retry_down = 2
		retry_up   = 2
		timeout    = 5
	}
`

func testDtcMonitorIcmpCompare(t *testing.T, resourceName string, expectedMonitorIcmp *ibclient.DtcMonitorIcmp) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		internalId := res.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("ID is not set")
		}

		ref, found := res.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("'ref' attribute is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"test")
		dtcMonitorIcmp, err := objMgr.SearchObjectByAltId("DtcMonitorIcmp", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedMonitorIcmp == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcMonitorIcmp
		recJson, _ := json.Marshal(dtcMonitorIcmp)
		err = json.Unmarshal(recJson, &rec)

		if rec.Comment != nil && expectedMonitorIcmp.Comment != nil {
			if *rec.Comment != *expectedMonitorIcmp.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedMonitorIcmp.Comment)
			}
		}

		if rec.Interval != nil && expectedMonitorIcmp.Interval != nil {
			if *rec.Interval != *expectedMonitorIcmp.Interval {
				return fmt.Errorf(
					"the value of 'interval' field is '%d', but expected '%d'",
					*rec.Interval, *expectedMonitorIcmp.Interval)
			}
		}

		if rec.Name != nil && expectedMonitorIcmp.Name != nil {
			if *rec.Name != *expectedMonitorIcmp.Name {
				return fmt.Errorf(
					"the value of 'name' field is '%s', but expected '%s'",
					*rec.Name, *expectedMonitorIcmp.Name)
			}
		}

		if rec.RetryDown != nil && expectedMonitorIcmp.RetryDown != nil {
			if *rec.RetryDown != *expectedMonitorIcmp.RetryDown {
				return fmt.Errorf(
					"the value of 'retry_down' field is '%d', but expected '%d'",
					*rec.RetryDown, *expectedMonitorIcmp.RetryDown)
			}
		}

		if rec.RetryUp != nil && expectedMonitorIcmp.RetryUp != nil {
			if *rec.RetryUp != *expectedMonitorIcmp.RetryUp {
				return fmt.Errorf(
					"the value of 'retry_up' field is '%d', but expected '%d'",
					*rec.RetryUp, *expectedMonitorIcmp.RetryUp)
			}
		}

		if rec.Timeout != nil && expectedMonitorIcmp.Timeout != nil {
			if *rec.Timeout != *expectedMonitorIcmp.Timeout {
				return fmt.Errorf(
					"the value of 'timeout' field is '%d', but expected '%d'",
					*rec.Timeout, *expectedMonitorIcmp.Timeout)
			}
		}

		return validateEAs(rec.Ea, expectedMonitorIcmp.Ea)
	}
}

func testDtcdtcMonitorIcmpDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_monitor_icmp" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcMonitorIcmpByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func TestAccResourceDtcMonitorIcmp(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcdtcMonitorIcmpDestroy,
		Steps: []resource.TestStep{
			// minimum params
			{
				Config: testResourceDtcMonitorIcmp1,
				Check: testDtcMonitorIcmpCompare(t,
					"infoblox_dtc_monitor_icmp.test-monitor_icmp1",
					&ibclient.DtcMonitorIcmp{
						Name: utils.StringPtr("test-monitor_icmp1"),
					}),
			},
			// rename
			{
				Config: testResourceDtcMonitorIcmp1_Rename,
				Check: testDtcMonitorIcmpCompare(t,
					"infoblox_dtc_monitor_icmp.test-monitor_icmp1",
					&ibclient.DtcMonitorIcmp{
						Name: utils.StringPtr("test-monitor_icmp1_rename"),
					}),
			},
			// maximum params
			{
				Config: testResourceDtcMonitorIcmp2,
				Check: testDtcMonitorIcmpCompare(t,
					"infoblox_dtc_monitor_icmp.test-monitor_icmp2",
					&ibclient.DtcMonitorIcmp{
						Name:      utils.StringPtr("test-monitor_icmp2"),
						Comment:   utils.StringPtr("test dtc monitor_icmp with max params"),
						Ea:        map[string]interface{}{"Site": "India"},
						Interval:  utils.Uint32Ptr(7),
						RetryDown: utils.Uint32Ptr(2),
						RetryUp:   utils.Uint32Ptr(2),
						Timeout:   utils.Uint32Ptr(5),
					}),
			},
		},
	})
}
