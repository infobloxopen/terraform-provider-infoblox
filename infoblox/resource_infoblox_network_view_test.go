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

func testAccCheckNetworkViewRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_network_view" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetNetworkViewByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccNetworkViewCompare(t *testing.T, resPath string, expectedRec *ibclient.NetworkView) resource.TestCheckFunc {
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
		nw, err := objMgr.SearchObjectByAltId("NetworkView", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.NetworkView
		recJson, _ := json.Marshal(nw)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil {
			return fmt.Errorf("network view's 'name' field is expected to be defined but it is not")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'network name' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
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

		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func TestAccResourceNetworkView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkViewRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_network_view" "foo"{
						name = "testNetworkView"
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccNetworkViewCompare(t, "infoblox_network_view.foo", &ibclient.NetworkView{
						Name:    utils.StringPtr("testNetworkView"),
						Comment: utils.StringPtr("test comment 1"),
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

					n := &ibclient.NetworkView{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name": "testNetworkView",
						},
					)
					var res []ibclient.NetworkView
					err := conn.GetObject(n, "", qp, &res)
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
					resource "infoblox_network_view" "foo"{
						name = "testNetworkView"
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_network_view.foo", "ext_attrs",
						`{"Location":"Test loc","Tenant ID":"terraform_test_tenant"}`,
					),
					// Actual API object should have Building EA
					testAccNetworkViewCompare(t, "infoblox_network_view.foo", &ibclient.NetworkView{
						Name:    utils.StringPtr("testNetworkView"),
						Comment: utils.StringPtr("test comment 1"),
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
					resource "infoblox_network_view" "foo"{
						name = "testNetworkView"
						comment = "Updated test comment"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`,
				Check: testAccNetworkViewCompare(t, "infoblox_network_view.foo", &ibclient.NetworkView{
					Name:    utils.StringPtr("testNetworkView"),
					Comment: utils.StringPtr("Updated test comment"),
					Ea: ibclient.EA{
						"Tenant ID": "terraform_test_tenant",
						"Location":  "Test loc",
						"Site":      "Test site",
					},
				}),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
					resource "infoblox_network_view" "foo"{
						name = "testNetworkView"
						comment = "Updated test comment"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
							"Site" = "Updated test site"
						})
					}`,
				Check: testAccNetworkViewCompare(t, "infoblox_network_view.foo", &ibclient.NetworkView{
					Name:    utils.StringPtr("testNetworkView"),
					Comment: utils.StringPtr("Updated test comment"),
					Ea: ibclient.EA{
						"Tenant ID": "terraform_test_tenant",
						"Location":  "Test loc",
						"Site":      "Updated test site",
					},
				}),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
					resource "infoblox_network_view" "foo"{
						name = "testNetworkView"
						comment = "Updated test comment"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_network_view.foo", "ext_attrs",
						`{"Location":"Test loc","Tenant ID":"terraform_test_tenant"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_network_view.foo"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_network_view.foo")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						nc, err := objMgr.GetNetworkViewByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := nc.Ea["Site"]; ok {
							return fmt.Errorf("site EA should've been removed, but still present in the WAPI object")
						}

						return nil
					},
				),
			},
		},
	})
}
