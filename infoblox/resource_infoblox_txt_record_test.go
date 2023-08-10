package infoblox

import (
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckTXTRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_txt_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetTXTRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccTXTRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordTXT) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		meta := testAccProvider.Meta()
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")

		rec, _ := objMgr.GetTXTRecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Name == nil {
			return fmt.Errorf("'fqdn' is expected to be defined but it is not")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
		}

		if rec.Text == nil {
			return fmt.Errorf("'text' is expected to be defined but it is not")
		}
		if *rec.Text != *expectedRec.Text {
			return fmt.Errorf(
				"'text does not match: got '%s', expected '%s'",
				*rec.Text, *expectedRec.Text)
		}

		if rec.View == nil {
			return fmt.Errorf("'dns_view' is expected to be defined but it is not")
		}
		if *rec.View != *expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				*rec.View, *expectedRec.View)
		}

		if rec.UseTtl != nil {
			if expectedRec.UseTtl == nil {
				return fmt.Errorf("'use_ttl' is expected to be undefined but it is not")
			}
			if *rec.UseTtl != *expectedRec.UseTtl {
				return fmt.Errorf(
					"'use_ttl' does not match: got '%t', expected '%t'",
					*rec.UseTtl, *expectedRec.UseTtl)
			}
			if *rec.UseTtl {
				if *rec.Ttl != *expectedRec.Ttl {
					return fmt.Errorf(
						"'TTL' usage does not match: got '%d', expected '%d'",
						rec.Ttl, expectedRec.Ttl)
				}
			}
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

func TestAccResourceTXTRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTXTRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_txt_record" "foo"{
						fqdn = "name1.test.com"
						text = "this is a sample text"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo", &ibclient.RecordTXT{
						View:   utils.StringPtr("default"),
						Name:   utils.StringPtr("name1.test.com"),
						Text:   utils.StringPtr("this is a sample text"),
						UseTtl: utils.BoolPtr(false),
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_txt_record" "foo2"{
						fqdn = "name2.test.com"
						text = "this is a sample text-2"
						ttl = 200
						dns_view = "nondefault_view"
						comment = "test comment"
						ext_attrs = jsonencode({
						  "Location" = "California"
						  "Site" = "HQ"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo2", &ibclient.RecordTXT{
						Text:    utils.StringPtr("this is a sample text-2"),
						Name:    utils.StringPtr("name2.test.com"),
						View:    utils.StringPtr("nondefault_view"),
						Ttl:     utils.Uint32Ptr(200),
						UseTtl:  utils.BoolPtr(true),
						Comment: utils.StringPtr("test comment"),
						Ea: ibclient.EA{
							"Location": "California",
							"Site":     "HQ",
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_txt_record" "foo2" {
						fqdn = "name3.test.com"
						text = "this is a text record"
						ttl = 150
						dns_view = "nondefault_view"
						comment = "test comment 2"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo2", &ibclient.RecordTXT{
						Text:    utils.StringPtr("this is a text record"),
						Name:    utils.StringPtr("name3.test.com"),
						View:    utils.StringPtr("nondefault_view"),
						Ttl:     utils.Uint32Ptr(150),
						UseTtl:  utils.BoolPtr(true),
						Comment: utils.StringPtr("test comment 2"),
					}),
				),
			},

			// negative test cases
			{
				Config: fmt.Sprintf(`
					resource "infoblox_txt_record" "foo2" {
						fqdn = "name3.test.com"
						text = "this is a text record"
						ttl = 150
						dns_view = "nondefault_view2"
						comment = "test comment 2"
					}`),
				ExpectError: regexp.MustCompile("changing the value of 'dns_view' field is not allowed"),
			},
		},
	})
}

func TestAcc_resourceTXTRecord_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTXTRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_txt_record" "foo3"{
					dns_view = "default"
					fqdn = "newtext.test.com"
					text = "this is sample text record"
					comment = "test comment on TXT record"
					ext_attrs = jsonencode({
						"Location" = "Some location"
					})
				}`,
				Check: testAccTXTRecordCompare(t, "infoblox_txt_record.foo3", &ibclient.RecordTXT{
					Text:    utils.StringPtr("this is sample text record"),
					Name:    utils.StringPtr("newtext.test.com"),
					View:    utils.StringPtr("default"),
					UseTtl:  utils.BoolPtr(false),
					Comment: utils.StringPtr("test comment on TXT record"),
					Ea: ibclient.EA{
						"Location": "Some location",
					},
				}),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.RecordTXT{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name": "newtext.test.com",
							"text": "this is sample text record",
						},
					)
					var res []ibclient.RecordTXT
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
				resource "infoblox_txt_record" "foo3"{
					dns_view = "default"
					fqdn = "newtext.test.com"
					text = "this is sample text record"
					comment = "test comment on TXT record"
					ext_attrs = jsonencode({
						"Location" = "Some location"
					})
				}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_txt_record.foo3", "ext_attrs",
						`{"Location":"Some location"}`,
					),
					// Actual API object should have Site EA
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo3", &ibclient.RecordTXT{
						Text:    utils.StringPtr("this is sample text record"),
						Name:    utils.StringPtr("newtext.test.com"),
						View:    utils.StringPtr("default"),
						UseTtl:  utils.BoolPtr(false),
						Comment: utils.StringPtr("test comment on TXT record"),
						Ea: ibclient.EA{
							"Location": "Some location",
							"Site":     "Test site",
						},
					}),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
				resource "infoblox_txt_record" "foo3"{
					dns_view = "default"
					fqdn = "newtext.test.com"
					text = "this is sample text record"
					comment = "updated comment on TXT record"
					ext_attrs = jsonencode({
						"Location" = "Some location"
					})
				}`,
				Check: testAccTXTRecordCompare(t, "infoblox_txt_record.foo3", &ibclient.RecordTXT{
					Text:    utils.StringPtr("this is sample text record"),
					Name:    utils.StringPtr("newtext.test.com"),
					View:    utils.StringPtr("default"),
					UseTtl:  utils.BoolPtr(false),
					Comment: utils.StringPtr("updated comment on TXT record"),
					Ea: ibclient.EA{
						"Location": "Some location",
						"Site":     "Test site",
					},
				}),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
				resource "infoblox_txt_record" "foo3"{
					dns_view = "default"
					fqdn = "newtext.test.com"
					text = "this is sample text record"
					comment = "test comment on TXT record"
					ext_attrs = jsonencode({
						"Location" = "Some location"
						"Site" = "Sample text site"
					})
				}`,
				Check: testAccTXTRecordCompare(t, "infoblox_txt_record.foo3", &ibclient.RecordTXT{
					Text:    utils.StringPtr("this is sample text record"),
					Name:    utils.StringPtr("newtext.test.com"),
					View:    utils.StringPtr("default"),
					UseTtl:  utils.BoolPtr(false),
					Comment: utils.StringPtr("test comment on TXT record"),
					Ea: ibclient.EA{
						"Location": "Some location",
						"Site":     "Sample text site",
					},
				}),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
				resource "infoblox_txt_record" "foo3"{
					dns_view = "default"
					fqdn = "newtext.test.com"
					text = "this is sample text record"
					comment = "test comment on TXT record"
					ext_attrs = jsonencode({
						"Location" = "Some location"
					})
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_txt_record.foo3", "ext_attrs",
						`{"Location":"Some location"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_txt_record.foo3"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_txt_record.foo3")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						trec, err := objMgr.GetTXTRecordByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}
						if _, ok := trec.Ea["Site"]; ok {
							return fmt.Errorf("Site EA should've been removed, but still present in the WAPI object")
						}
						return nil
					},
				),
			},
		},
	})
}
