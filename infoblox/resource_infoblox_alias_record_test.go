package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
	"testing"
)

var testResourceAliasRecord = `resource "infoblox_alias_record" "record1" {
    name = "alias2222.test.com"
    comment = "test alias record"
    target_name = "aa.gg.com"
	target_type = "PTR"
	view = "default"
	ttl = 3600
	disable = false
	ext_attrs = jsonencode({
    	"Location" = "65.8665701230204, -37.00791763398113"
  	})
}`

var testResourceAliasRecord2 = `resource "infoblox_alias_record" "record2" {
    name = "alias5678.test.com"
    target_name = "aa.bb.com"
	target_type = "A"
}`

var testResourceAliasRecord3 = `resource "infoblox_alias_record" "record3" {
    name = "alias999.test.com"
    target_name = "aa.bb.com"
	target_type = ""
}`

func testAliasRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_alias_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetAliasRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testAliasRecordCompare(t *testing.T, resourceName string, aliasRecord *ibclient.RecordAlias) resource.TestCheckFunc {
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
		zf, err := objMgr.SearchObjectByAltId("AliasRecord", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if aliasRecord == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.RecordAlias
		recJson, _ := json.Marshal(zf)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil && aliasRecord.Name != nil {
			if *rec.Name != *aliasRecord.Name {
				return fmt.Errorf(
					"the value of 'name' field is '%s', but expected '%s'",
					*rec.Name, *aliasRecord.Name)
			}
		}
		if rec.TargetName != nil && aliasRecord.TargetName != nil {
			if *rec.TargetName != *aliasRecord.TargetName {
				return fmt.Errorf(
					"the value of 'target_name' field is '%s', but expected '%s'",
					*rec.TargetName, *aliasRecord.TargetName)
			}
		}
		if rec.TargetType != aliasRecord.TargetType {
			return fmt.Errorf("the value of 'target_type' field is '%s', but expected '%s'", rec.TargetType, aliasRecord.TargetType)
		}
		if rec.View != nil && aliasRecord.View != nil {
			if *rec.View != *aliasRecord.View {
				return fmt.Errorf(
					"the value of 'view' field is '%s', but expected '%s'",
					*rec.View, *aliasRecord.View)
			}
		}
		if rec.Comment != nil && aliasRecord.Comment != nil {
			if *rec.Comment != *aliasRecord.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *aliasRecord.Comment)
			}
		}
		if rec.Disable != nil && aliasRecord.Disable != nil {
			if *rec.Disable != *aliasRecord.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *aliasRecord.Disable)
			}
		}
		if rec.Ttl != nil && aliasRecord.Ttl != nil {
			if *rec.Ttl != *aliasRecord.Ttl {
				return fmt.Errorf(
					"the value of 'ttl' field is '%d', but expected '%d'",
					*rec.Ttl, *aliasRecord.Ttl)
			}
		}
		return validateEAs(rec.Ea, aliasRecord.Ea)
	}
}

func TestAccResourceAliasRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAliasRecordDestroy,
		Steps: []resource.TestStep{
			// maximum params
			{
				Config: testResourceAliasRecord,
				Check: testAliasRecordCompare(t, "infoblox_alias_record.record1", &ibclient.RecordAlias{
					Name:       utils.StringPtr("alias2222.test.com"),
					TargetName: utils.StringPtr("aa.gg.com"),
					TargetType: "PTR",
					View:       utils.StringPtr("default"),
					Ttl:        utils.Uint32Ptr(3600),
					Disable:    utils.BoolPtr(false),
					Ea:         map[string]interface{}{"Location": "65.8665701230204, -37.00791763398113"},
					Comment:    utils.StringPtr("test alias record"),
				}),
			},
			// minimum params
			{
				Config: testResourceAliasRecord2,
				Check: testAliasRecordCompare(t, "infoblox_alias_record.record2", &ibclient.RecordAlias{
					Name:       utils.StringPtr("alias5678.test.com"),
					TargetName: utils.StringPtr("aa.bb.com"),
					TargetType: "A",
					//View:       utils.StringPtr("default"),
					//Ttl:       utils.Uint32Ptr(3600),
					//Disable:   utils.BoolPtr(false),
					//Ea:        map[string]interface{}{"Location": "65.8665701230204, -37.00791763398113"},
					//Comment: utils.StringPtr("test alias record"),
				}),
			},
			// negative test case
			{
				Config:      testResourceAliasRecord3,
				ExpectError: regexp.MustCompile("name, targetName and targetType are required to create an Alias Record"),
			},
		},
	})
}
