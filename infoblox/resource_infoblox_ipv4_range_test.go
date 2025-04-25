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

func testAccRangeCompare(
	t *testing.T,
	resPath string,
	expectedRec *ibclient.Range) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
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
		recRange, err := objMgr.SearchObjectByAltId("Range", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		var rec *ibclient.Range
		recJson, _ := json.Marshal(recRange)
		err = json.Unmarshal(recJson, &rec)
		if err != nil {
			return fmt.Errorf("error unmarshalling record: %v", err)
		}
		if rec.Options != nil && expectedRec.Options != nil {
			if len(rec.Options) != len(expectedRec.Options) {
				return fmt.Errorf("the length of 'options' field is '%d' but expected '%d'", len(rec.Options), len(expectedRec.Options))
			}
			for i := range rec.Options {
				if !reflect.DeepEqual(rec.Options[i], expectedRec.Options[i]) {
					return fmt.Errorf("difference found at index %d: got '%v' but expected '%v'", i, rec.Options[i], expectedRec.Options[i])
				}
			}
		}
		if rec.Member != nil && expectedRec.Member != nil {
			if rec.Member.Name != expectedRec.Member.Name {
				return fmt.Errorf("the 'member' field is '%s' but expected '%s'", rec.Member.Name, expectedRec.Member.Name)
			}
		}
		if rec.ServerAssociationType != expectedRec.ServerAssociationType {
			return fmt.Errorf("the 'server_association_type' field is '%s' but expected '%s'", rec.ServerAssociationType, expectedRec.ServerAssociationType)
		}
		if rec.Template != expectedRec.Template {
			return fmt.Errorf("the 'template' field is '%s' but expected '%s'", rec.Template, expectedRec.Template)
		}
		if rec.NetworkView != nil && expectedRec.NetworkView != nil {
			if *rec.NetworkView != *expectedRec.NetworkView {
				return fmt.Errorf("the 'network_view' field is '%s' but expected '%s'", *rec.NetworkView, *expectedRec.NetworkView)
			}
		}
		if rec.Network != nil && expectedRec.Network != nil {
			if *rec.Network != *expectedRec.Network {
				return fmt.Errorf("the 'network' field is '%s' but expected '%s'", *rec.Network, *expectedRec.Network)
			}
		}
		if rec.StartAddr != nil && expectedRec.StartAddr != nil {
			if *rec.StartAddr != *expectedRec.StartAddr {
				return fmt.Errorf("the 'start_addr' field is '%s' but expected '%s'", *rec.StartAddr, *expectedRec.StartAddr)
			}
		}
		if rec.EndAddr != nil && expectedRec.EndAddr != nil {
			if *rec.EndAddr != *expectedRec.EndAddr {
				return fmt.Errorf("the 'end_addr' field is '%s' but expected '%s'", *rec.EndAddr, *expectedRec.EndAddr)
			}
		}
		if rec.FailoverAssociation != nil && expectedRec.FailoverAssociation != nil {
			if *rec.FailoverAssociation != *expectedRec.FailoverAssociation {
				return fmt.Errorf("the 'failover_association' field is '%s' but expected '%s'", *rec.FailoverAssociation, *expectedRec.FailoverAssociation)
			}
		}
		if rec.Comment != nil && expectedRec.Comment != nil {
			if *rec.Comment != *expectedRec.Comment {
				return fmt.Errorf("the 'comment' field is '%s' but expected '%s'", *rec.Comment, *expectedRec.Comment)
			}
		}
		if rec.Name != nil && expectedRec.Name != nil {
			if *rec.Name != *expectedRec.Name {
				return fmt.Errorf("the 'name' field is '%s' but expected '%s'", *rec.Name, *expectedRec.Name)
			}
		}
		if rec.Disable != nil && expectedRec.Disable != nil {
			if *rec.Disable != *expectedRec.Disable {
				return fmt.Errorf("the 'disable' field is '%t' but expected '%t'", *rec.Disable, *expectedRec.Disable)
			}
		}
		if rec.UseOptions != nil && expectedRec.UseOptions != nil {
			if *rec.UseOptions != *expectedRec.UseOptions {
				return fmt.Errorf("the 'use_options' field is '%t' but expected '%t'", *rec.UseOptions, *expectedRec.UseOptions)
			}
		}

		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func testAccCheckRangeDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_range" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetNetworkRangeByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func TestAccResourceRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRangeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_range" "range3" {
 							start_addr = "17.0.0.45"
 							end_addr   = "17.0.0.50"
							options {
							name         = "dhcp-lease-time"
							value        = "43200"
							vendor_class = "DHCP"
							num          = 51
							use_option   = true
						}
				}`),
				Check: resource.ComposeTestCheckFunc(
					testAccRangeCompare(t, "infoblox_ipv4_range.range3", &ibclient.Range{
						StartAddr: utils.StringPtr("17.0.0.45"),
						EndAddr:   utils.StringPtr("17.0.0.50"),
						Options: []*ibclient.Dhcpoption{{
							Name:        "dhcp-lease-time",
							Value:       "43200",
							VendorClass: "DHCP",
							Num:         51,
							UseOption:   true,
						}},
						ServerAssociationType: "NONE",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_range" "range3" {
 							start_addr = "17.0.0.20"
 							end_addr   = "17.0.0.40"
							options {
							name         = "dhcp-lease-time"
							value        = "43200"
							vendor_class = "DHCP"
							num          = 51
							use_option   = true
						}
						options {
    									name = "routers"
    									num = "3"
    									use_option = true
    									value = "17.0.0.2"
    									vendor_class = "DHCP"
  								}
						network = "17.0.0.0/24"
						comment = "test comment"		
						name = "test_range"
						disable = false
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccRangeCompare(t, "infoblox_ipv4_range.range3", &ibclient.Range{
						StartAddr: utils.StringPtr("17.0.0.20"),
						EndAddr:   utils.StringPtr("17.0.0.40"),
						Options: []*ibclient.Dhcpoption{{
							Name:        "dhcp-lease-time",
							Value:       "43200",
							VendorClass: "DHCP",
							Num:         51,
							UseOption:   true,
						},
							{
								Name:        "routers",
								Value:       "17.0.0.2",
								VendorClass: "DHCP",
								Num:         3,
								UseOption:   true,
							}},

						Name:                  utils.StringPtr("test_range"),
						Comment:               utils.StringPtr("test comment"),
						Disable:               utils.BoolPtr(false),
						ServerAssociationType: "NONE",
					})),
			},
		},
	})
}
