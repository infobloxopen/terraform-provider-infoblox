package infoblox

import (
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var testAccresourceCNAMERecordCreate = fmt.Sprintf(`
resource "infoblox_cname_record" "foo"{
	dns_view="default"
	canonical="test-canonicalName.test.com"
	alias="test-aliasname.test.com"
	comment="CNAME record created"
	ext_attrs = jsonencode({
		"Tenant ID" = "terraform_test_tenant"
		"Location" = "Test loc"
		"Site" = "Test site"
		"TestEA1"=["text1","text2"]
	  })

}`)

var testAccresourceCNAMERecordUpdate = fmt.Sprintf(`
resource "infoblox_cname_record" "foo"{
	dns_view="default"
	canonical="test-canonicalName.test.com"
	alias="test-aliasname.test.com"
	comment="CNAME record updated"
	ext_attrs = jsonencode({
		"Tenant ID" = "terraform_test_tenant"
		"Location" = "Test loc 2"
		"Site" = "Test site 2"
		"TestEA1"="text3"
	  })

}`)

func validateRecordCNAME(
	resourceName string,
	expectedValue *ibclient.RecordCNAME) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id := res.Primary.ID
		if id == "" {
			return fmt.Errorf("ID is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"terraform_test_tenant")
		recCNAME, err := objMgr.GetCNAMERecordByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}

		expCanonical := *expectedValue.Canonical
		if *recCNAME.Canonical != expCanonical {
			return fmt.Errorf(
				"the value of 'canonical' field is '%s', but expected '%s'",
				*recCNAME.Canonical, expCanonical)
		}

		expName := *expectedValue.Name
		if *recCNAME.Name != expName {
			return fmt.Errorf(
				"the value of 'alias Name' field is '%s', but expected '%s'",
				*recCNAME.Name, expName)
		}

		expComment := *expectedValue.Comment
		if *recCNAME.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				*recCNAME.Comment, expComment)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && recCNAME.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && recCNAME.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(recCNAME.Ea, expectedEAs)
	}
}

func TestAccResourceCNAMERecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCNAMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceCNAMERecordCreate,
				Check: validateRecordCNAME(
					"infoblox_cname_record.foo",
					&ibclient.RecordCNAME{
						View:      utils.StringPtr("default"),
						Canonical: utils.StringPtr("test-canonicalName.test.com"),
						Name:      utils.StringPtr("test-aliasname.test.com"),
						Zone:      "test.com",
						Comment:   utils.StringPtr("CNAME record created"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
			{
				Config: testAccresourceCNAMERecordUpdate,
				Check: validateRecordCNAME(
					"infoblox_cname_record.foo",
					&ibclient.RecordCNAME{
						View:      utils.StringPtr("default"),
						Canonical: utils.StringPtr("test-canonicalName.test.com"),
						Name:      utils.StringPtr("test-aliasname.test.com"),
						Zone:      "test.com",
						Comment:   utils.StringPtr("CNAME record updated"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc 2",
							"Site":      "Test site 2",
							"TestEA1":   "text3",
						},
					},
				),
			},
		},
	})
}

func testAccCheckCNAMERecordDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_cname_record" {
			continue
		}
		res, err := objMgr.GetCNAMERecordByRef(rs.Primary.ID)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if res != nil {
			return fmt.Errorf("object with ID '%s' remains", rs.Primary.ID)
		}
	}
	return nil
}

func TestAcc_resourceCNAMERecord_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCNAMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_cname_record" "foo1"{
					dns_view="default"
					canonical="test.somewhere.net"
					alias="sample1.test.com"
					comment="CNAME record for sample"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"Site" = "Marking site"
					})
				}`,
				Check: validateRecordCNAME(
					"infoblox_cname_record.foo1",
					&ibclient.RecordCNAME{
						View:      utils.StringPtr("default"),
						Canonical: utils.StringPtr("test.somewhere.net"),
						Name:      utils.StringPtr("sample1.test.com"),
						Zone:      "test.com",
						Comment:   utils.StringPtr("CNAME record for sample"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Marking site",
						},
					},
				),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EA
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.RecordCNAME{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name":      "sample1.test.com",
							"canonical": "test.somewhere.net",
						},
					)
					var res []ibclient.RecordCNAME
					err := conn.GetObject(n, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].Ea["Location"] = "California"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
				resource "infoblox_cname_record" "foo1"{
					dns_view="default"
					canonical="test.somewhere.net"
					alias="sample1.test.com"
					comment="CNAME record for sample"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"Site" = "Marking site"
					})
				}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Location EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_cname_record.foo1", "ext_attrs",
						`{"Site":"Marking site","Tenant ID":"terraform_test_tenant"}`,
					),
					// Actual API object should have Location EA
					validateRecordCNAME(
						"infoblox_cname_record.foo1",
						&ibclient.RecordCNAME{
							View:      utils.StringPtr("default"),
							Canonical: utils.StringPtr("test.somewhere.net"),
							Name:      utils.StringPtr("sample1.test.com"),
							Zone:      "test.com",
							Comment:   utils.StringPtr("CNAME record for sample"),
							Ea: ibclient.EA{
								"Tenant ID": "terraform_test_tenant",
								"Site":      "Marking site",
								"Location":  "California",
							},
						},
					),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
				resource "infoblox_cname_record" "foo1"{
					dns_view="default"
					canonical="testhidden.somewhere.net"
					alias="sample1.test.com"
					comment="CNAME record for sample"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"Site" = "Marking site"
					})
				}`,
				Check: validateRecordCNAME(
					"infoblox_cname_record.foo1",
					&ibclient.RecordCNAME{
						View:      utils.StringPtr("default"),
						Canonical: utils.StringPtr("testhidden.somewhere.net"),
						Name:      utils.StringPtr("sample1.test.com"),
						Zone:      "test.com",
						Comment:   utils.StringPtr("CNAME record for sample"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Marking site",
							"Location":  "California",
						},
					},
				),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
				resource "infoblox_cname_record" "foo1"{
					dns_view="default"
					canonical="test.somewhere.net"
					alias="sample1.test.com"
					comment="CNAME record for sample"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"Site" = "Marking site"
						"Location" = "New California"
					})
				}`,
				Check: validateRecordCNAME(
					"infoblox_cname_record.foo1",
					&ibclient.RecordCNAME{
						View:      utils.StringPtr("default"),
						Canonical: utils.StringPtr("test.somewhere.net"),
						Name:      utils.StringPtr("sample1.test.com"),
						Zone:      "test.com",
						Comment:   utils.StringPtr("CNAME record for sample"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Marking site",
							"Location":  "New California",
						},
					},
				),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
				resource "infoblox_cname_record" "foo1"{
					dns_view="default"
					canonical="test.somewhere.net"
					alias="sample1.test.com"
					comment="CNAME record for sample"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"Site" = "Marking site"
					})
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_cname_record.foo1", "ext_attrs",
						`{"Site":"Marking site","Tenant ID":"terraform_test_tenant"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_cname_record.foo1"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_cname_network.foo1")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						crec, err := objMgr.GetCNAMERecordByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := crec.Ea["Location"]; ok {
							return fmt.Errorf("Location EA should've been removed, but still present in the WAPI object")
						}

						return nil
					},
				),
			},
		},
	})
}
