package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"reflect"
	"testing"
)

func testNSRecordCompare(t *testing.T, resourceName string, expectedNSRecord *ibclient.RecordNS) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if res.Primary.ID == "" {
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

		RecordNS, err := objMgr.GetNSRecordByRef(ref)
		if err != nil {
			if isNotFoundError(err) {
				if expectedNSRecord == nil {
					return nil
				}
				return fmt.Errorf("object with ref %s not found", ref)
			}
		}
		var rec *ibclient.RecordNS
		recJson, _ := json.Marshal(RecordNS)
		err = json.Unmarshal(recJson, &rec)
		if rec.Name == "" {
			return fmt.Errorf("'name' is expected to be defined but it is not")
		}
		if rec.Name != expectedNSRecord.Name {
			return fmt.Errorf(
				"'name' does not match: got '%s', expected '%s'",
				rec.Name,
				expectedNSRecord.Name)
		}
		if rec.Nameserver != nil {
			if *rec.Nameserver != *expectedNSRecord.Nameserver {
				return fmt.Errorf("'nameserver' does not match: got '%s', expected '%s'",
					*rec.Nameserver, *expectedNSRecord.Nameserver)
			}
		}
		if rec.Addresses != nil && expectedNSRecord.Addresses != nil {
			if len(rec.Addresses) != len(expectedNSRecord.Addresses) {
				return fmt.Errorf("the length of 'addresses' field is '%d' but expected '%d'", len(rec.Addresses), len(expectedNSRecord.Addresses))
			}

			for i := range rec.Addresses {
				if !reflect.DeepEqual(rec.Addresses[i], expectedNSRecord.Addresses[i]) {
					return fmt.Errorf("difference found at index %d: got '%v' but expected '%v'", i, rec.Addresses[i], expectedNSRecord.Addresses[i])
				}
			}
		}
		if rec.Creator != expectedNSRecord.Creator {
			return fmt.Errorf("'creator' does not match: got '%s', expected '%s'", rec.Creator, expectedNSRecord.Creator)
		}
		if rec.View != expectedNSRecord.View {
			return fmt.Errorf("'view' does not match: got '%s', expected '%s'", rec.View, expectedNSRecord.View)
		}
		return nil
	}
}
func testAccCheckNSRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_ns_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetNSRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func TestAccResourceNSRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
							resource "infoblox_ns_record" "rec1" {
							name = "test.com"
							nameserver = "name6.test.com"
 							addresses{
   							address = "2.3.4.5"
							auto_create_ptr=true
 							}
							dns_view="default"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testNSRecordCompare(t, "infoblox_ns_record.rec1", &ibclient.RecordNS{
						Name:       "test.com",
						Nameserver: utils.StringPtr("name6.test.com"),
						Addresses: []*ibclient.ZoneNameServer{{
							Address:       "2.3.4.5",
							AutoCreatePtr: true,
						}},
						Creator: "STATIC",
						View:    "default",
					})),
			},
		},
	})
}
