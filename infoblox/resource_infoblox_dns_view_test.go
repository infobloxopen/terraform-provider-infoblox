package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"testing"
)

func TestAcc_resourceDNSView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSViewDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_dns_view" "test_view" {
						name = "test_view"
						comment = "Acceptance test DNS view"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							Location = "Test location"
					  	})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "name", `test_view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "comment",
						`Acceptance test DNS view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "ext_attrs",
						`{"Location":"Test location","Tenant ID":"terraform_test_tenant"}`),
					validateDnsView("infoblox_dns_view.test_view", &ibclient.View{
						Name:        utils.StringPtr("test_view"),
						Comment:     utils.StringPtr("Acceptance test DNS view"),
						NetworkView: utils.StringPtr("default"),
						Ea: map[string]interface{}{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
						},
					}),
				),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					v := &ibclient.View{}
					v.SetReturnFields(append(v.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name": "test_view",
						},
					)
					var res []ibclient.View
					err := conn.GetObject(v, "", qp, &res)
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
					resource "infoblox_dns_view" "test_view" {
						name = "test_view"
						comment = "Acceptance test DNS view"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							Location = "Test location"
					  	})
					}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_dns_view.test_view", "ext_attrs",
						`{"Location":"Test location","Tenant ID":"terraform_test_tenant"}`,
					),
					// Actual API object should have Site EA
					validateDnsView("infoblox_dns_view.test_view", &ibclient.View{
						Name:        utils.StringPtr("test_view"),
						Comment:     utils.StringPtr("Acceptance test DNS view"),
						NetworkView: utils.StringPtr("default"),
						Ea: map[string]interface{}{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
							"Site":      "Test site",
						},
					}),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
					resource "infoblox_dns_view" "test_view" {
						name = "test_view"
						comment = "Acceptance test updated DNS view"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							Location = "Test location"
					  	})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "name", `test_view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "comment",
						`Acceptance test updated DNS view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "ext_attrs",
						`{"Location":"Test location","Tenant ID":"terraform_test_tenant"}`),
					// Actual API object should have Site EA
					validateDnsView("infoblox_dns_view.test_view", &ibclient.View{
						Name:        utils.StringPtr("test_view"),
						Comment:     utils.StringPtr("Acceptance test updated DNS view"),
						NetworkView: utils.StringPtr("default"),
						Ea: map[string]interface{}{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
							"Site":      "Test site",
						},
					}),
				),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
					resource "infoblox_dns_view" "test_view" {
						name = "test_view"
						comment = "Acceptance test updated DNS view"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							Location = "Test location"
							Site = "Updated test site"
					  	})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "name", `test_view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "comment",
						`Acceptance test updated DNS view`),
					resource.TestCheckResourceAttr("infoblox_dns_view.test_view", "ext_attrs",
						`{"Location":"Test location","Site":"Updated test site","Tenant ID":"terraform_test_tenant"}`),
					// Actual API object should have Site EA
					validateDnsView("infoblox_dns_view.test_view", &ibclient.View{
						Name:        utils.StringPtr("test_view"),
						Comment:     utils.StringPtr("Acceptance test updated DNS view"),
						NetworkView: utils.StringPtr("default"),
						Ea: map[string]interface{}{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
							"Site":      "Updated test site",
						},
					}),
				),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
					resource "infoblox_dns_view" "test_view" {
						name = "test_view"
						comment = "Acceptance test updated DNS view"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							Location = "Test location"
					  	})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_dns_view.test_view", "ext_attrs",
						`{"Location":"Test location","Tenant ID":"terraform_test_tenant"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_dns_view.test_view"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_dns_view.test_view")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						v := &ibclient.View{}
						v.SetReturnFields([]string{"extattrs"})

						vResult := ibclient.View{}

						err := conn.GetObject(v, id, nil, &vResult)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := vResult.Ea["Site"]; ok {
							return fmt.Errorf("site EA should've been removed, but still present in the WAPI object")
						}

						return nil
					},
				),
			},
		},
	})
}

func validateDnsView(
	resourceName string,
	expectedValue *ibclient.View) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id := res.Primary.ID
		if id == "" {
			return fmt.Errorf("ID is not set")
		}

		conn := testAccProvider.Meta().(ibclient.IBConnector)
		v := &ibclient.View{}
		v.SetReturnFields([]string{"name", "comment", "network_view", "extattrs"})

		actualView := ibclient.View{}

		err := conn.GetObject(v, id, nil, &actualView)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}

		if *actualView.Name != *expectedValue.Name {
			return fmt.Errorf(
				"the value of 'name' field is '%s', but expected '%s'",
				*actualView.Name, *expectedValue.Name)
		}

		if *actualView.Comment != *expectedValue.Comment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				*actualView.Comment, *expectedValue.Comment)
		}

		if *actualView.NetworkView != *expectedValue.NetworkView {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				*actualView.NetworkView, *expectedValue.NetworkView)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && actualView.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && actualView.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(actualView.Ea, expectedEAs)
	}
}

func testAccCheckDNSViewDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(ibclient.IBConnector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_dns_view" {
			continue
		}

		v := &ibclient.View{}
		vResult := &ibclient.View{}

		err := conn.GetObject(v, rs.Primary.ID, nil, vResult)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if vResult != nil {
			return fmt.Errorf("object with ID '%s' remains", rs.Primary.ID)
		}
	}

	return nil
}
