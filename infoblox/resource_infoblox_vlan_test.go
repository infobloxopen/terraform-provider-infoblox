package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckVlanDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_vlan" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		vlan := ibclient.NewEmptyVlan()
		err := connector.GetObject(vlan, rs.Primary.ID, ibclient.NewQueryParams(false, nil), vlan)
		if err == nil {
			return fmt.Errorf("VLAN still exists")
		}
		// Suppress unused variable warning
		_ = objMgr
	}
	return nil
}

func testAccVlanCompare(t *testing.T, resPath string, expectedRec *ibclient.Vlan) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("Not found: %s", resPath)
		}

		internalId := res.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("ID is not set")
		}

		ref, found := res.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("'ref' attribute is not set")
		}

		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"test")
		obj, err := objMgr.SearchObjectByAltId("Vlan", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.Vlan
		recJson, _ := json.Marshal(obj)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil {
			return fmt.Errorf("VLAN's 'name' field is expected to be defined but it is not")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'name' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
		}

		if rec.Id == nil {
			return fmt.Errorf("VLAN's 'id' field is expected to be defined but it is not")
		}
		if *rec.Id != *expectedRec.Id {
			return fmt.Errorf(
				"'id' does not match: got '%d', expected '%d'",
				*rec.Id,
				*expectedRec.Id)
		}

		if rec.Comment != nil {
			if expectedRec.Comment == nil {
				return fmt.Errorf("'comment' is expected to be undefined but it is not")
			}
			if *rec.Comment != *expectedRec.Comment {
				return fmt.Errorf(
					"'comment' does not match: got '%s', expected '%s'",
					*rec.Comment, *expectedRec.Comment)
			}
		} else if expectedRec.Comment != nil {
			return fmt.Errorf("'comment' is expected to be defined but it is not")
		}

		if rec.Description != nil {
			if expectedRec.Description == nil {
				return fmt.Errorf("'description' is expected to be undefined but it is not")
			}
			if *rec.Description != *expectedRec.Description {
				return fmt.Errorf(
					"'description' does not match: got '%s', expected '%s'",
					*rec.Description, *expectedRec.Description)
			}
		} else if expectedRec.Description != nil {
			return fmt.Errorf("'description' is expected to be defined but it is not")
		}

		if rec.Department != nil {
			if expectedRec.Department == nil {
				return fmt.Errorf("'department' is expected to be undefined but it is not")
			}
			if *rec.Department != *expectedRec.Department {
				return fmt.Errorf(
					"'department' does not match: got '%s', expected '%s'",
					*rec.Department, *expectedRec.Department)
			}
		} else if expectedRec.Department != nil {
			return fmt.Errorf("'department' is expected to be defined but it is not")
		}

		if rec.Contact != nil {
			if expectedRec.Contact == nil {
				return fmt.Errorf("'contact' is expected to be undefined but it is not")
			}
			if *rec.Contact != *expectedRec.Contact {
				return fmt.Errorf(
					"'contact' does not match: got '%s', expected '%s'",
					*rec.Contact, *expectedRec.Contact)
			}
		} else if expectedRec.Contact != nil {
			return fmt.Errorf("'contact' is expected to be defined but it is not")
		}

		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func TestAccResourceVlan(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_vlan" "foo"{
						name = "test-vlan"
						vlan_id = 100
						comment = "test comment 1"
						description = "Test VLAN description"
						department = "IT"
						contact = "admin@example.com"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccVlanCompare(t, "infoblox_vlan.foo", &ibclient.Vlan{
						Name:        utils.StringPtr("test-vlan"),
						Id:          utils.Uint32Ptr(100),
						Comment:     utils.StringPtr("test comment 1"),
						Description: utils.StringPtr("Test VLAN description"),
						Department:  utils.StringPtr("IT"),
						Contact:     utils.StringPtr("admin@example.com"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
						},
					}),
				),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					vlan := ibclient.NewEmptyVlan()
					vlan.SetReturnFields(append(vlan.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name": "test-vlan",
						},
					)
					var res []ibclient.Vlan
					err := conn.GetObject(vlan, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].Ea["Site"] = "Test site"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
					resource "infoblox_vlan" "foo"{
						name = "test-vlan"
						vlan_id = 100
						comment = "test comment 1"
						description = "Test VLAN description"
						department = "IT"
						contact = "admin@example.com"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_vlan.foo", "ext_attrs",
						`{"Location":"Test loc","Tenant ID":"terraform_test_tenant"}`,
					),
					// Actual API object should have Site EA
					testAccVlanCompare(t, "infoblox_vlan.foo", &ibclient.Vlan{
						Name:        utils.StringPtr("test-vlan"),
						Id:          utils.Uint32Ptr(100),
						Comment:     utils.StringPtr("test comment 1"),
						Description: utils.StringPtr("Test VLAN description"),
						Department:  utils.StringPtr("IT"),
						Contact:     utils.StringPtr("admin@example.com"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
							"Site":      "Test site",
						},
					}),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
					resource "infoblox_vlan" "foo"{
						name = "test-vlan"
						vlan_id = 100
						comment = "Updated test comment"
						description = "Updated VLAN description"
						department = "Engineering"
						contact = "eng@example.com"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`,
				Check: testAccVlanCompare(t, "infoblox_vlan.foo", &ibclient.Vlan{
					Name:        utils.StringPtr("test-vlan"),
					Id:          utils.Uint32Ptr(100),
					Comment:     utils.StringPtr("Updated test comment"),
					Description: utils.StringPtr("Updated VLAN description"),
					Department:  utils.StringPtr("Engineering"),
					Contact:     utils.StringPtr("eng@example.com"),
					Ea: ibclient.EA{
						"Tenant ID": "terraform_test_tenant",
						"Location":  "Test loc",
						"Site":      "Test site",
					},
				}),
			},
		},
	})
}
