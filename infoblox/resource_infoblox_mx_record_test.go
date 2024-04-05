package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckMXRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_mx_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetMXRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccMXRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordMX) resource.TestCheckFunc {
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

		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"test")
		recMx, err := objMgr.SearchObjectByAltId("MX", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.RecordMX
		recJson, _ := json.Marshal(recMx)
		err = json.Unmarshal(recJson, &rec)

		//rec, _ := objMgr.GetMXRecordByRef(res.Primary.ID)
		//if rec == nil {
		//	return fmt.Errorf("record not found")
		//}

		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
		}

		if *rec.View != *expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				*rec.View, *expectedRec.View)
		}

		if *rec.MailExchanger != *expectedRec.MailExchanger {
			return fmt.Errorf(
				"'mail_exchanger' does not match: got '%s', expected '%s'",
				*rec.MailExchanger, *expectedRec.MailExchanger)
		}

		if *rec.Preference != *expectedRec.Preference {
			return fmt.Errorf(
				"'priority' does not match: got '%d', expected '%d'",
				*rec.Preference, *expectedRec.Preference)
		}

		if expectedRec.Ttl != nil {
			if *rec.UseTtl != *expectedRec.UseTtl {
				return fmt.Errorf(
					"TTL usage does not match: got '%t', expected '%t'",
					*rec.UseTtl, *expectedRec.UseTtl)
			}
		}
		if *rec.UseTtl {
			if *rec.Ttl != *expectedRec.Ttl {
				return fmt.Errorf(
					"'Ttl' usage does not match: got '%d', expected '%d'",
					*rec.Ttl, *expectedRec.Ttl)
			}
		}

		expComment := *expectedRec.Comment
		if expComment != "" {
			if *rec.Comment != expComment {
				return fmt.Errorf(
					"'comment' does not match: got '%s', expected '%s'",
					*rec.Comment, *expectedRec.Comment)
			}
		}

		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func TestAccResourceMXRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMXRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone" {
						fqdn = "test.com"
					}
					resource "infoblox_mx_record" "foo"{
						fqdn = "name1.test.com"
						mail_exchanger = "sample.mx1.com"
						preference = 25
						depends_on = [infoblox_zone_auth.zone]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo", &ibclient.RecordMX{
						View:          utils.StringPtr("default"),
						Name:          utils.StringPtr("name1.test.com"),
						MailExchanger: utils.StringPtr("sample.mx1.com"),
						Preference:    utils.Uint32Ptr(25),
						Ttl:           utils.Uint32Ptr(0),
						UseTtl:        utils.BoolPtr(false),
						Comment:       utils.StringPtr(""),
						Ea:            nil,
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo1"{
						fqdn = "name2.test.com"
						mail_exchanger = "sample.mx2.com"
						preference = 30
						ttl = 300
						dns_view = "nondefault_view"
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Location" = "Los Angeles"
							"Site" = "HQ"
						})
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo1", &ibclient.RecordMX{
						Name:          utils.StringPtr("name2.test.com"),
						MailExchanger: utils.StringPtr("sample.mx2.com"),
						Preference:    utils.Uint32Ptr(30),
						View:          utils.StringPtr("nondefault_view"),
						Ttl:           utils.Uint32Ptr(300),
						UseTtl:        utils.BoolPtr(true),
						Comment:       utils.StringPtr("test comment 1"),
						Ea: ibclient.EA{
							"Location": "Los Angeles",
							"Site":     "HQ",
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo2"{
						fqdn = "name3.test.com"
						mail_exchanger = "sample.mx3.com"
						preference = 35
						ttl = 150
						dns_view = "nondefault_view"
						comment = "test comment 2"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo2", &ibclient.RecordMX{
						Name:          utils.StringPtr("name3.test.com"),
						MailExchanger: utils.StringPtr("sample.mx3.com"),
						Preference:    utils.Uint32Ptr(35),
						View:          utils.StringPtr("nondefault_view"),
						Ttl:           utils.Uint32Ptr(150),
						UseTtl:        utils.BoolPtr(true),
						Comment:       utils.StringPtr("test comment 2"),
						Ea:            nil,
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo3"{
						fqdn = "name4.test.com"
						dns_view = "nondefault_view"
						mail_exchanger = "sample.mx3.com"
						preference = 35
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo3", &ibclient.RecordMX{
						Name:          utils.StringPtr("name4.test.com"),
						View:          utils.StringPtr("nondefault_view"),
						MailExchanger: utils.StringPtr("sample.mx3.com"),
						Preference:    utils.Uint32Ptr(35),
						Ttl:           utils.Uint32Ptr(0),
						UseTtl:        utils.BoolPtr(false),
						Comment:       utils.StringPtr(""),
						Ea:            nil,
					}),
				),
			},

			// negative test cases
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo3"{
						fqdn = "name4.test.com"
						dns_view = "nondefault_view"
						mail_exchanger = "sample.mx3.com"
						preference = 350000
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("'preference' must be integer and must be in the range from 0 to 65535 inclusively"),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo3"{
						fqdn = "name4.test.com"
						dns_view = "nondefault_view"
						mail_exchanger = "sample.mx3.com"
						preference = 35
						ttl = -1
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("TTL value must be 0 or higher"),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view.name
					}
					resource "infoblox_mx_record" "foo3"{
						fqdn = "name4.test.com"
						dns_view = "nondefault_view2"
						mail_exchanger = "sample.mx3.com"
						preference = 35
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("changing the value of 'dns_view' field is not allowed"),
			},
		},
	})
}

func TestAcc_resourceMXRecord_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMXRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_mx_record" "foo4"{
					dns_view = "default"
					fqdn = "testdemo.test.com"
					mail_exchanger = "sample.mx4.com"
					preference = 40
					comment = "test comment on MX record"
					ext_attrs = jsonencode({
						"Location" = "Test MX Location"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: testAccMXRecordCompare(t, "infoblox_mx_record.foo4", &ibclient.RecordMX{
					Name:          utils.StringPtr("testdemo.test.com"),
					View:          utils.StringPtr("default"),
					MailExchanger: utils.StringPtr("sample.mx4.com"),
					Preference:    utils.Uint32Ptr(40),
					Comment:       utils.StringPtr("test comment on MX record"),
					Ea: ibclient.EA{
						"Location": "Test MX Location",
					},
				}),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.RecordMX{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name":           "testdemo.test.com",
							"mail_exchanger": "sample.mx4.com",
						},
					)
					var res []ibclient.RecordA
					err := conn.GetObject(n, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].Ea["Site"] = "Sample MX site"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_mx_record" "foo4"{
					dns_view = "default"
					fqdn = "testdemo.test.com"
					mail_exchanger = "sample.mx4.com"
					preference = 40
					comment = "test comment on MX record"
					ext_attrs = jsonencode({
						"Location" = "Test MX Location"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_mx_record.foo4", "ext_attrs",
						`{"Location":"Test MX Location"}`,
					),
					// Actual API object should have Site EA
					testAccMXRecordCompare(t, "infoblox_mx_record.foo4", &ibclient.RecordMX{
						Name:          utils.StringPtr("testdemo.test.com"),
						View:          utils.StringPtr("default"),
						MailExchanger: utils.StringPtr("sample.mx4.com"),
						Preference:    utils.Uint32Ptr(40),
						Comment:       utils.StringPtr("test comment on MX record"),
						Ea: ibclient.EA{
							"Location": "Test MX Location",
							"Site":     "Sample MX site",
						},
					}),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_mx_record" "foo4"{
					dns_view = "default"
					fqdn = "testdemo.test.com"
					mail_exchanger = "sample.mx4.com"
					preference = 40
					comment = "updated comment on MX record"
					ext_attrs = jsonencode({
						"Location" = "Test MX Location"
					})
				}`,
				Check: testAccMXRecordCompare(t, "infoblox_mx_record.foo4", &ibclient.RecordMX{
					Name:          utils.StringPtr("testdemo.test.com"),
					View:          utils.StringPtr("default"),
					MailExchanger: utils.StringPtr("sample.mx4.com"),
					Preference:    utils.Uint32Ptr(40),
					Comment:       utils.StringPtr("updated comment on MX record"),
					Ea: ibclient.EA{
						"Location": "Test MX Location",
						"Site":     "Sample MX site",
					},
				}),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_mx_record" "foo4"{
					dns_view = "default"
					fqdn = "testdemo.test.com"
					mail_exchanger = "sample.mx4.com"
					preference = 40
					comment = "test comment on MX record"
					ext_attrs = jsonencode({
						"Location" = "Test MX Location"
						"Site" = "New Modern site"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: testAccMXRecordCompare(t, "infoblox_mx_record.foo4", &ibclient.RecordMX{
					Name:          utils.StringPtr("testdemo.test.com"),
					View:          utils.StringPtr("default"),
					MailExchanger: utils.StringPtr("sample.mx4.com"),
					Preference:    utils.Uint32Ptr(40),
					Comment:       utils.StringPtr("test comment on MX record"),
					Ea: ibclient.EA{
						"Location": "Test MX Location",
						"Site":     "New Modern site",
					},
				}),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_mx_record" "foo4"{
					dns_view = "default"
					fqdn = "testdemo.test.com"
					mail_exchanger = "sample.mx4.com"
					preference = 40
					comment = "test comment on MX record"
					ext_attrs = jsonencode({
						"Location" = "Test MX Location"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_mx_record.foo4", "ext_attrs",
						`{"Location":"Test MX Location"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_mx_record.foo4"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_mx_record.foo4")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						mrec, err := objMgr.GetMXRecordByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := mrec.Ea["Site"]; ok {
							return fmt.Errorf("Site EA should've been removed, but still present in the WAPI object")
						}
						return nil
					},
				),
			},
		},
	})
}
