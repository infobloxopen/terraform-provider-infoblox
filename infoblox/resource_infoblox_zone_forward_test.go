package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"reflect"
	"regexp"
	"testing"
)

var testResourceZoneForwardRecord = `resource "infoblox_zone_forward" "testfz1" {
    fqdn = "test_fz.ex.org"
    comment = "test sample forward zone"
    forward_to {
        name = "test123.dz.ex.com"
        address = "10.0.0.1"
    }
    forward_to {
        name = "test245.dz.ex.com"
        address = "10.0.0.2"
    }
}`

var testResourceZoneForward = `resource "infoblox_zone_forward" "testfz1" {
    fqdn = "test_fz.ex.org"
    comment = "test sample forward zone"
    forwarding_servers {
        name = "infoblox.172_28_82_176"
        forwarders_only = true
        use_override_forwarders = true
        forward_to {
                name = "cc.fwd.com"
                address = "10.1.1.1"
        }
    }
}`

func testZoneForwardDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_zone_forward" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetZoneForwardByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testForwardZoneCompare(t *testing.T, resourceName string, expectedZF *ibclient.ZoneForward) resource.TestCheckFunc {
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
		zf, err := objMgr.SearchObjectByAltId("ZoneForward", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedZF == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.ZoneForward
		recJson, _ := json.Marshal(zf)
		err = json.Unmarshal(recJson, &rec)

		if rec.Fqdn != expectedZF.Fqdn {
			return fmt.Errorf(
				"the value of 'fqdn' field is '%s', but expected '%s'",
				rec.Fqdn, expectedZF.Fqdn)
		}
		if rec.View != nil && expectedZF.View != nil {
			if *rec.View != *expectedZF.View {
				return fmt.Errorf(
					"the value of 'view' field is '%s', but expected '%s'",
					*rec.View, *expectedZF.View)
			}
		}
		if rec.ZoneFormat != expectedZF.ZoneFormat {
			return fmt.Errorf(
				"the value of 'zone_format' field is '%s', but expected '%s'",
				rec.ZoneFormat, expectedZF.ZoneFormat)
		}
		if rec.Comment != nil && expectedZF.Comment != nil {
			if *rec.Comment != *expectedZF.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedZF.Comment)
			}
		}
		if rec.Disable != nil && expectedZF.Disable != nil {
			if *rec.Disable != *expectedZF.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *expectedZF.Disable)
			}
		}
		if rec.ForwardersOnly != nil && expectedZF.ForwardersOnly != nil {
			if *rec.ForwardersOnly != *expectedZF.ForwardersOnly {
				return fmt.Errorf(
					"the value of 'forwarders_only' field is '%t', but expected '%t'",
					*rec.ForwardersOnly, *expectedZF.ForwardersOnly)
			}
		}
		if rec.NsGroup != nil && expectedZF.NsGroup != nil {
			if *rec.NsGroup != *expectedZF.NsGroup {
				return fmt.Errorf(
					"the value of 'ns_group' field is '%s', but expected '%s'",
					*rec.NsGroup, *expectedZF.NsGroup)
			}
		}
		if rec.ForwardTo.ForwardTo != nil && expectedZF.ForwardTo.ForwardTo != nil {
			if !reflect.DeepEqual(rec.ForwardTo, expectedZF.ForwardTo) {
				return fmt.Errorf(
					"the value of 'forward_to' field is '%v', but expected '%v'",
					rec.ForwardTo, expectedZF.ForwardTo)
			}
		}
		if rec.ForwardingServers != nil && expectedZF.ForwardingServers != nil {
			if !reflect.DeepEqual(rec.ForwardingServers, expectedZF.ForwardingServers) {
				return fmt.Errorf(
					"the value of 'forwarding_servers' field is '%v', but expected '%v'",
					rec.ForwardingServers, expectedZF.ForwardingServers)
			}
		}
		return validateEAs(rec.Ea, expectedZF.Ea)
	}
}

func TestAccResourceZoneForward(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testZoneForwardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceZoneForwardRecord,
				Check: testForwardZoneCompare(t, "infoblox_zone_forward.testfz1", &ibclient.ZoneForward{
					Fqdn:       "test_fz.ex.org",
					View:       utils.StringPtr("default"),
					ZoneFormat: "FORWARD",
					Comment:    utils.StringPtr("test sample forward zone"),
					ForwardTo: ibclient.NullForwardTo{
						IsNull: false,
						ForwardTo: []ibclient.NameServer{
							{Name: "test123.dz.ex.com", Address: "10.0.0.1"},
							{Name: "test245.dz.ex.com", Address: "10.0.0.2"},
						}},
				}),
			},
			// negative test case
			{
				Config:      testResourceZoneForward,
				ExpectError: regexp.MustCompile("either external_ns_group or forward_to must be set"),
			},
		},
	})
}
