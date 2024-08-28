package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var testResourceZoneDelegatedRecord = `resource "infoblox_zone_auth" "zone1" {
  fqdn = "test3.com"
  view = "default"
  zone_format = "FORWARD"
  ns_group = ""
  restart_if_needed = true
  soa_default_ttl = 36000
  soa_expire = 72000
  soa_negative_ttl = 600
  soa_refresh = 1800
  soa_retry = 900
  comment = "Zone Auth created newly"
  ext_attrs = jsonencode({
    Location = "AcceptanceTerraform"
  })
}

resource "infoblox_zone_delegated" "testzd1" {
    fqdn = "test_zd.test3.com"
    delegate_to {
        name = "ns2.infoblox.com"
        address = "10.0.0.1"
    }
    depends_on = [infoblox_zone_auth.zone1]
}`

func testAccCheckZoneDelegatedDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_zone_delegated" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetZoneDelegatedByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("zone delegation record found after destroy")
		}
	}
	return nil
}

func testAccZoneDelegatedCompare(t *testing.T, resPath string, expectedRec *ibclient.ZoneDelegated) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
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
		zd, err := objMgr.SearchObjectByAltId("ZoneDelegated", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}

		// Assertion of object type and error handling
		var rec *ibclient.ZoneDelegated
		recJson, _ := json.Marshal(zd)
		err = json.Unmarshal(recJson, &rec)

		if zd == nil {
			return fmt.Errorf("zone delegated record not found")
		}

		if expectedRec == nil {
			return fmt.Errorf("expected record is nil")
		}

		if rec.Fqdn != expectedRec.Fqdn {
			return fmt.Errorf(
				"the value of 'fqdn' field is '%s', but expected '%s'",
				rec.Fqdn, expectedRec.Fqdn)
		}
		if rec.View != nil && expectedRec.View != nil {
			if *rec.View != *expectedRec.View {
				return fmt.Errorf(
					"the value of 'view' field is '%s', but expected '%s'",
					*rec.View, *expectedRec.View)
			}
		}

		if rec.ZoneFormat != expectedRec.ZoneFormat {
			return fmt.Errorf(
				"the value of 'zone_format' field is '%s', but expected '%s'",
				rec.ZoneFormat, expectedRec.ZoneFormat)
		}
		if rec.Comment != nil && expectedRec.Comment != nil {
			if *rec.Comment != *expectedRec.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedRec.Comment)
			}
		}
		if rec.Disable != nil && expectedRec.Disable != nil {
			if *rec.Disable != *expectedRec.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *expectedRec.Disable)
			}
		}
		if rec.Locked != nil && expectedRec.Locked != nil {
			if *rec.Locked != *expectedRec.Locked {
				return fmt.Errorf(
					"the value of 'locked' field is '%t', but expected '%t'",
					*rec.Locked, *expectedRec.Locked)
			}
		}
		if rec.DelegatedTtl != nil && expectedRec.DelegatedTtl != nil {
			if *rec.DelegatedTtl != *expectedRec.DelegatedTtl {
				return fmt.Errorf(
					"the value of 'delegated_ttl' field is '%d', but expected '%d'",
					*rec.DelegatedTtl, *expectedRec.DelegatedTtl)
			}
		}

		if rec.NsGroup != nil && expectedRec.NsGroup != nil {
			if *rec.NsGroup != *expectedRec.NsGroup {
				return fmt.Errorf(
					"the value of 'ns_group' field is '%s', but expected '%s'",
					*rec.NsGroup, *expectedRec.NsGroup)
			}
		}
		if rec.DelegateTo.NameServers != nil && expectedRec.DelegateTo.NameServers != nil {
			if !reflect.DeepEqual(rec.DelegateTo, expectedRec.DelegateTo) {
				return fmt.Errorf(
					"the value of 'delegate_to' field is '%v', but expected '%v'",
					rec.DelegateTo, expectedRec.DelegateTo)
			}
		}

		return validateEAs(rec.Ea, expectedRec.Ea)

	}
}

func TestAccResourceZoneDelegated(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDelegatedDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceZoneDelegatedRecord,
				Check: testAccZoneDelegatedCompare(t, "infoblox_zone_delegated.testzd1", &ibclient.ZoneDelegated{
					Fqdn: "test_zd.test3.com",
					DelegateTo: ibclient.NullableNameServers{
						IsNull: false,
						NameServers: []ibclient.NameServer{
							{Name: "ns2.infoblox.com", Address: "10.0.0.1"},
						}},
					ZoneFormat: "FORWARD",
					View:       utils.StringPtr("default"),
				}),
			}},
	})
}
