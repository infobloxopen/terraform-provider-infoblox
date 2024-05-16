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

func testAccCheckSRVRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_srv_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetSRVRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testAccSRVRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordSRV) resource.TestCheckFunc {
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
		recSRV, err := objMgr.SearchObjectByAltId("SRV", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.RecordSRV
		recJson, _ := json.Marshal(recSRV)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil {
			return fmt.Errorf("'name' is expected to be defined but it is not")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'name' does not match: got '%s', expected '%s'",
				*rec.Name, *expectedRec.Name)
		}

		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
		}

		if rec.Priority == nil {
			return fmt.Errorf("'priority' is expected to be defined but it is not")
		}
		if *rec.Priority != *expectedRec.Priority {
			return fmt.Errorf(
				"'priority' does not match: got '%d', expected '%d'",
				rec.Priority, expectedRec.Priority)
		}

		if rec.Weight == nil {
			return fmt.Errorf("'weight' is expected to be defined but it is not")
		}
		if *rec.Weight != *expectedRec.Weight {
			return fmt.Errorf(
				"'weight' does not match: got '%d', expected '%d'",
				rec.Weight, expectedRec.Weight)
		}

		if rec.Port == nil {
			return fmt.Errorf("'port' is expected to be defined but it is not")
		}
		if *rec.Port != *expectedRec.Port {
			return fmt.Errorf(
				"'port' does not match: got '%d', expected '%d'",
				rec.Port, expectedRec.Port)
		}

		if rec.Target == nil {
			return fmt.Errorf("'target' is expected to be defined but it is not")
		}
		if *rec.Target != *expectedRec.Target {
			return fmt.Errorf(
				"'target' does not match: got '%s', expected '%s'",
				*rec.Target, *expectedRec.Target)
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

func TestAccResourceSRVRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSRVRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone" {
						fqdn = "test.com"
					}
					resource "infoblox_srv_record" "foo"{
						name = "_sip._tcp.host1.test.com"
						priority = 50
						weight = 30
						port = 80
						target = "sample.target1.com"
						depends_on = [infoblox_zone_auth.zone]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo", &ibclient.RecordSRV{
						View:     "default",
						Name:     utils.StringPtr("_sip._tcp.host1.test.com"),
						Priority: utils.Uint32Ptr(50),
						Weight:   utils.Uint32Ptr(30),
						Port:     utils.Uint32Ptr(80),
						Target:   utils.StringPtr("sample.target1.com"),
						UseTtl:   utils.BoolPtr(false),
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
					resource "infoblox_srv_record" "foo1" {
						dns_view = "nondefault_view"
						name = "_sip._udp.host2.test.com"
						priority = 60
						weight = 40
						port = 36
						target = "sample.target2.com"
						ttl = 300 //300s
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Location" = "France"
							"Site" = "DHQ"
						})
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo1", &ibclient.RecordSRV{
						View:     "nondefault_view",
						Name:     utils.StringPtr("_sip._udp.host2.test.com"),
						Priority: utils.Uint32Ptr(60),
						Weight:   utils.Uint32Ptr(40),
						Port:     utils.Uint32Ptr(36),
						Target:   utils.StringPtr("sample.target2.com"),
						Ttl:      utils.Uint32Ptr(300),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 1"),
						Ea: ibclient.EA{
							"Location": "France",
							"Site":     "DHQ",
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_http._tcp.demo.host3.test.com"
						priority = 100
						weight = 50
						port = 88
						target = "sample.target3.com"
						ttl = 140
						comment = "test comment 2"
						ext_attrs = jsonencode({
							"Site" = "DHQ"
						})
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo2", &ibclient.RecordSRV{
						View:     "nondefault_view",
						Name:     utils.StringPtr("_http._tcp.demo.host3.test.com"),
						Priority: utils.Uint32Ptr(100),
						Weight:   utils.Uint32Ptr(50),
						Port:     utils.Uint32Ptr(88),
						Target:   utils.StringPtr("sample.target3.com"),
						Ttl:      utils.Uint32Ptr(140),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 2"),
						Ea: ibclient.EA{
							"Site": "DHQ",
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_http._tcp.demo.host4.test.com"
						priority = 101
						weight = 51
						port = 89
						target = "sample.target4.com"
						ttl = 141
						comment = "test comment 3"
						ext_attrs = jsonencode({
							"Site" = "None"
						})
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo2", &ibclient.RecordSRV{
						View:     "nondefault_view",
						Name:     utils.StringPtr("_http._tcp.demo.host4.test.com"),
						Priority: utils.Uint32Ptr(101),
						Weight:   utils.Uint32Ptr(51),
						Port:     utils.Uint32Ptr(89),
						Target:   utils.StringPtr("sample.target4.com"),
						Ttl:      utils.Uint32Ptr(141),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 3"),
						Ea: ibclient.EA{
							"Site": "None",
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101
						weight = 51
						port = 89
						target = "sample.target4.com"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo2", &ibclient.RecordSRV{
						View:     "nondefault_view",
						Name:     utils.StringPtr("_customservice._newcoolproto.demo.host4.test.com"),
						Priority: utils.Uint32Ptr(101),
						Weight:   utils.Uint32Ptr(51),
						Port:     utils.Uint32Ptr(89),
						Target:   utils.StringPtr("sample.target4.com"),
						UseTtl:   utils.BoolPtr(false),
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101
						weight = 51
						port = 89
						target = "sample.target4.com"
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101
						weight = 51
						port = 89000
						target = "sample.target4.com"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("'port' must be integer and must be in the range from 0 to 65535 inclusively"),
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101000
						weight = 51
						port = 89
						target = "sample.target4.com"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("'priority' must be integer and must be in the range from 0 to 65535 inclusively"),
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
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101
						weight = 510000
						port = 89
						target = "sample.target4.com"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexp.MustCompile("'weight' must be integer and must be in the range from 0 to 65535 inclusively"),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view2"
					}
					resource "infoblox_zone_auth" "zone2" {
						fqdn = "test.com"
						view = infoblox_dns_view.view1.name
					}
					resource "infoblox_srv_record" "foo2"{
						dns_view = "nondefault_view2"
						name = "_customservice._newcoolproto.demo.host4.test.com"
						priority = 101
						weight = 51
						port = 89
						target = "sample.target4.com"
						depends_on = [infoblox_zone_auth.zone2]
					}`),
				ExpectError: regexp.MustCompile("changing the value of 'dns_view' field is not allowed"),
			},
		},
	})
}

func TestAcc_resourceSRVRecord_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSRVRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_srv_record" "foo3" {
					dns_view = "default"
					name = "_sip._udp.test.com"
					priority = 12
					weight = 10
					port = 5060
					target = "sip.example2.org"
					comment = "example SRV record"
					ext_attrs = jsonencode({
						"Location" = "Newyork"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: testAccSRVRecordCompare(t, "infoblox_srv_record.foo3", &ibclient.RecordSRV{
					View:     "default",
					Name:     utils.StringPtr("_sip._udp.test.com"),
					Priority: utils.Uint32Ptr(12),
					Weight:   utils.Uint32Ptr(10),
					Port:     utils.Uint32Ptr(5060),
					UseTtl:   utils.BoolPtr(false),
					Target:   utils.StringPtr("sip.example2.org"),
					Comment:  utils.StringPtr("example SRV record"),
					Ea: ibclient.EA{
						"Location": "Newyork",
					},
				}),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.RecordSRV{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name":   "_sip._udp.test.com",
							"target": "sip.example2.org",
							"port":   "5060",
						},
					)
					var res []ibclient.RecordSRV
					err := conn.GetObject(n, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].View = ""
					res[0].Ea["Site"] = "SRV new site"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_srv_record" "foo3" {
					dns_view = "default"
					name = "_sip._udp.test.com"
					priority = 12
					weight = 10
					port = 5060
					target = "sip.example2.org"
					comment = "example SRV record"
					ext_attrs = jsonencode({
						"Location" = "Newyork"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_srv_record.foo3", "ext_attrs",
						`{"Location":"Newyork"}`,
					),
					// Actual API object should have Site EA
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo3", &ibclient.RecordSRV{
						View:     "default",
						Name:     utils.StringPtr("_sip._udp.test.com"),
						Priority: utils.Uint32Ptr(12),
						Weight:   utils.Uint32Ptr(10),
						Port:     utils.Uint32Ptr(5060),
						UseTtl:   utils.BoolPtr(false),
						Target:   utils.StringPtr("sip.example2.org"),
						Comment:  utils.StringPtr("example SRV record"),
						Ea: ibclient.EA{
							"Location": "Newyork",
							"Site":     "SRV new site",
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
				resource "infoblox_srv_record" "foo3" {
					dns_view = "default"
					name = "_sip._udp.test.com"
					priority = 12
					weight = 10
					port = 5060
					target = "sip.example2.org"
					comment = "updated example SRV record"
					ext_attrs = jsonencode({
						"Location" = "Newyork"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: testAccSRVRecordCompare(t, "infoblox_srv_record.foo3", &ibclient.RecordSRV{
					View:     "default",
					Name:     utils.StringPtr("_sip._udp.test.com"),
					Priority: utils.Uint32Ptr(12),
					Weight:   utils.Uint32Ptr(10),
					Port:     utils.Uint32Ptr(5060),
					UseTtl:   utils.BoolPtr(false),
					Target:   utils.StringPtr("sip.example2.org"),
					Comment:  utils.StringPtr("updated example SRV record"),
					Ea: ibclient.EA{
						"Location": "Newyork",
						"Site":     "SRV new site",
					},
				}),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_srv_record" "foo3" {
					dns_view = "default"
					name = "_sip._udp.test.com"
					priority = 12
					weight = 10
					port = 5060
					target = "sip.example2.org"
					comment = "example SRV record"
					ext_attrs = jsonencode({
						"Location" = "Newyork"
						"Site" = "random new site"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: testAccSRVRecordCompare(t, "infoblox_srv_record.foo3", &ibclient.RecordSRV{
					View:     "default",
					Name:     utils.StringPtr("_sip._udp.test.com"),
					Priority: utils.Uint32Ptr(12),
					Weight:   utils.Uint32Ptr(10),
					Port:     utils.Uint32Ptr(5060),
					UseTtl:   utils.BoolPtr(false),
					Target:   utils.StringPtr("sip.example2.org"),
					Comment:  utils.StringPtr("example SRV record"),
					Ea: ibclient.EA{
						"Location": "Newyork",
						"Site":     "random new site",
					},
				}),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone" {
					fqdn = "test.com"
				}
				resource "infoblox_srv_record" "foo3" {
					dns_view = "default"
					name = "_sip._udp.test.com"
					priority = 12
					weight = 10
					port = 5060
					target = "sip.example2.org"
					comment = "example SRV record"
					ext_attrs = jsonencode({
						"Location" = "Newyork"
					})
					depends_on = [infoblox_zone_auth.zone]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_srv_record.foo3", "ext_attrs",
						`{"Location":"Newyork"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_srv_record.foo3"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_srv_record.foo3")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						srec, err := objMgr.GetSRVRecordByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := srec.Ea["Site"]; ok {
							return fmt.Errorf("Site EA should've been removed, but still present in the WAPI object")
						}
						return nil
					},
				),
			},
		},
	})
}
