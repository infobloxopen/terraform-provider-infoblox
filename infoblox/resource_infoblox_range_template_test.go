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

var testResourceRangeTemplate1 = `resource "infoblox_ipv4_range_template" "template1" {
  name = "rangetemplate11"
  number_of_addresses = 10
  offset = 70
}`

var testResourceRangeTemplate2 = `resource "infoblox_ipv4_range_template" "template2" {
  name = "range-template22"
  number_of_addresses = 40
  offset = 30
  comment = "Temporary Range Template"
  use_options = true
  ext_attrs = jsonencode({
    "Site" = "Kobe"
  })
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
}`

func testAccCheckRangeTemplateDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ipv4_range_template" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetRangeTemplateByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("range template record found after destroy")
		}
	}
	return nil
}

func testAccRangeTemplateCompare(t *testing.T, resPath string, expectedRec *ibclient.Rangetemplate) resource.TestCheckFunc {
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
		rangeTemplate, err := objMgr.SearchObjectByAltId("RangeTemplate", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}

		// Assertion of object type and error handling
		var rec *ibclient.Rangetemplate
		recJson, _ := json.Marshal(rangeTemplate)
		err = json.Unmarshal(recJson, &rec)

		if rangeTemplate == nil {
			return fmt.Errorf("range template record not found")
		}

		if expectedRec == nil {
			return fmt.Errorf("expected record is nil")
		}

		if rec.Name != nil && expectedRec.Name != nil {
			if *rec.Name != *expectedRec.Name {
				return fmt.Errorf(
					"the value of 'name' field is '%s', but expected '%s'",
					*rec.Name, *expectedRec.Name)
			}
		}
		if rec.Comment != nil && expectedRec.Comment != nil {
			if *rec.Comment != *expectedRec.Comment {
				return fmt.Errorf(
					"the value of 'coment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedRec.Comment)
			}
		}
		if rec.NumberOfAddresses != nil && expectedRec.NumberOfAddresses != nil {
			if *rec.NumberOfAddresses != *expectedRec.NumberOfAddresses {
				return fmt.Errorf(
					"the value of 'number_of_addresses' field is '%d', but expected '%d'",
					*rec.NumberOfAddresses, *expectedRec.NumberOfAddresses)
			}
		}
		if rec.Offset != nil && expectedRec.Offset != nil {
			if *rec.Offset != *expectedRec.Offset {
				return fmt.Errorf(
					"the value of 'offset' field is '%d', but expected '%d'",
					*rec.Offset, *expectedRec.Offset)
			}
		}

		if rec.UseOptions != nil && expectedRec.UseOptions != nil {
			if *rec.UseOptions != *expectedRec.UseOptions {
				return fmt.Errorf(
					"the value of 'use_options' field is '%t', but expected '%t'",
					*rec.UseOptions, *expectedRec.UseOptions)
			}
		}
		if rec.FailoverAssociation != nil && expectedRec.FailoverAssociation != nil {
			if *rec.FailoverAssociation != *expectedRec.FailoverAssociation {
				return fmt.Errorf(
					"the value of 'failover_association' field is '%s', but expected '%s'",
					*rec.FailoverAssociation, *expectedRec.FailoverAssociation)
			}
		}
		if rec.ServerAssociationType != expectedRec.ServerAssociationType {
			return fmt.Errorf(
				"the value of 'server_association_type' field is '%s', but expected '%s'",
				rec.ServerAssociationType, expectedRec.ServerAssociationType)
		}

		if rec.Options != nil && expectedRec.Options != nil {
			if !compareDHCPOptions(rec.Options, expectedRec.Options) {
				return fmt.Errorf("the value of 'options' field is '%v', but expected '%v'", rec.Options, expectedRec.Options)
			}
		}
		if rec.Member != nil && expectedRec.Member != nil {
			if !reflect.DeepEqual(rec.Member, expectedRec.Member) {
				return fmt.Errorf(
					"the value of 'member' field is '%v', but expected '%v'",
					rec.Member, expectedRec.Member)
			}
		}

		return validateEAs(rec.Ea, expectedRec.Ea)

	}
}

func TestAccResourceRangeTemplate(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRangeTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceRangeTemplate1,
				Check: testAccRangeTemplateCompare(t, "infoblox_ipv4_range_template.template1", &ibclient.Rangetemplate{
					Name:              utils.StringPtr("rangetemplate11"),
					NumberOfAddresses: utils.Uint32Ptr(10),
					Offset:            utils.Uint32Ptr(70),
					Options: []*ibclient.Dhcpoption{{
						Name:        "dhcp-lease-time",
						Value:       "43200",
						VendorClass: "DHCP",
						Num:         51,
						UseOption:   true,
					}},
					ServerAssociationType: "NONE",
				}),
			},
			{
				Config: testResourceRangeTemplate2,
				Check: testAccRangeTemplateCompare(t, "infoblox_ipv4_range_template.template2", &ibclient.Rangetemplate{
					Name:              utils.StringPtr("range-template22"),
					NumberOfAddresses: utils.Uint32Ptr(40),
					Offset:            utils.Uint32Ptr(30),
					Options: []*ibclient.Dhcpoption{
						{
							Name:        "domain-name-servers",
							Value:       "11.22.33.44",
							VendorClass: "DHCP",
							Num:         6,
							UseOption:   true,
						},
						{
							Name:        "dhcp-lease-time",
							Value:       "43200",
							VendorClass: "DHCP",
							Num:         51,
							UseOption:   true,
						},
					},
					ServerAssociationType: "NONE",
					Ea: map[string]interface{}{
						"Site": "Kobe",
					},
				}),
			},
		},
	})
}

func compareDHCPOptions(options1, options2 []*ibclient.Dhcpoption) bool {
	if len(options1) != len(options2) {
		return false
	}
	for i := range options1 {
		if options1[i].Name != options2[i].Name || options1[i].Value != options2[i].Value {
			return false
		}
	}
	return true
}
