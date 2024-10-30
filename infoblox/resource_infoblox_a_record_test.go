package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"net"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckARecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_a_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetARecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccARecordCompare(
	t *testing.T,
	resPath string,
	expectedRec *ibclient.RecordA,
	notExpectedIpAddr string,
	expectedCidr string,
	expectedFilterParams string) resource.TestCheckFunc {

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
		recA, err := objMgr.SearchObjectByAltId("A", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.RecordA
		recJson, _ := json.Marshal(recA)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil {
			return fmt.Errorf("'fqdn' is expected to be defined but it is not")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
		}
		if rec.Ipv4Addr == nil {
			return fmt.Errorf("'ipv4addr' is expected to be defined but it is not")
		}
		if notExpectedIpAddr != "" && notExpectedIpAddr == *rec.Ipv4Addr {
			return fmt.Errorf(
				"'ip_addr' field has value '%s' but that is not expected to happen",
				notExpectedIpAddr)
		}
		if expectedCidr != "" {
			_, parsedCidr, err := net.ParseCIDR(expectedCidr)
			if err != nil {
				panic(fmt.Sprintf("cannot parse CIDR '%s': %s", expectedCidr, err))
			}

			if !parsedCidr.Contains(net.ParseIP(*rec.Ipv4Addr)) {
				return fmt.Errorf(
					"IP address '%s' does not belong to the expected CIDR '%s'",
					*rec.Ipv4Addr, expectedCidr)
			}
		}
		if expectedRec.Ipv4Addr != nil {
			if *expectedRec.Ipv4Addr == "" {
				expectedRec.Ipv4Addr = utils.StringPtr(res.Primary.Attributes["ip_addr"])
			}
			if *rec.Ipv4Addr != *expectedRec.Ipv4Addr {
				return fmt.Errorf(
					"'ipv4address' does not match: got '%s', expected '%s'",
					*rec.Ipv4Addr, *expectedRec.Ipv4Addr)
			}
		}
		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
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

		actualFilterParams, exists := res.Primary.Attributes["filter_params"]
		if expectedFilterParams != "" {
			if !exists {
				return fmt.Errorf("'filter_params' is expected to be defined but it is not")
			}
			if actualFilterParams != expectedFilterParams {
				return fmt.Errorf(
					"'filter_params' does not match: got '%s', expected '%s'",
					actualFilterParams, expectedFilterParams)
			}
		} else if exists {
			return fmt.Errorf("'filter_params' is expected to be undefined but it is not")
		}
		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

var (
	regexpRequiredMissingIPv4    = regexp.MustCompile("either of 'ip_addr' or 'cidr' or 'filter_params' values is required")
	regexpCidrIpAddrConflictIPv4 = regexp.MustCompile("only one of 'ip_addr' or 'cidr' or 'filter_params' values is allowed to be defined")

	regexpUpdateConflictIPv4      = regexp.MustCompile("only one of 'ip_addr' and 'cidr' values is allowed to update")
	regexpNetviewUpdateNotAllowed = regexp.MustCompile("changing the value of 'network_view' field is not allowed")
	regexpDnsviewUpdateNotAllowed = regexp.MustCompile("changing the value of 'dns_view' field is not allowed")
)

func TestAccResourceARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckARecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = "nondefault_view"
						depends_on = [infoblox_dns_view.view1]
					}
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
						ip_addr = "10.0.0.2"
						cidr = "10.20.30.0/24"
                        network_view = "default"
					}`),
				ExpectError: regexpCidrIpAddrConflictIPv4,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = "nondefault_view"
						depends_on = [infoblox_dns_view.view1]
					}
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
					}`),
				ExpectError: regexpRequiredMissingIPv4,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = "nondefault_view"
						depends_on = [infoblox_dns_view.view1]
					}
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
						ip_addr = "10.0.0.2"
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.0.0.2"),
						Name:     utils.StringPtr("name1.test.com"),
						View:     "nondefault_view",
						Ttl:      utils.Uint32Ptr(0),
						UseTtl:   utils.BoolPtr(false),
						Comment:  nil,
						Ea:       nil,
					}, "", "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = "nondefault_view"
						depends_on = [infoblox_dns_view.view1]
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name2.test.com"
						ip_addr = "192.168.31.31"
						ttl = 10
						dns_view = "nondefault_view"
						comment = "test comment 1"
						ext_attrs = jsonencode({
						  "Location" = "New York"
						  "Site" = "HQ"
						})
						depends_on = [infoblox_zone_auth.zone1, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("192.168.31.31"),
						Name:     utils.StringPtr("name2.test.com"),
						View:     "nondefault_view",
						Ttl:      utils.Uint32Ptr(10),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 1"),
						Ea: ibclient.EA{
							"Location": "New York",
							"Site":     "HQ",
						},
					}, "", "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.1"
						ttl = 155
						dns_view = "nondefault_view"
						comment = "test comment 2"
						depends_on = [infoblox_zone_auth.zone1, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.1"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						Ttl:      utils.Uint32Ptr(155),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 2"),
					}, "", "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.1"
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.1"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						UseTtl:   utils.BoolPtr(false),
					}, "", "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
			        resource "infoblox_ipv4_network" "net1" {
			            cidr = "10.20.30.0/24"
			            network_view = "default"
			        }
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
			            cidr = infoblox_ipv4_network.net1.cidr
			            network_view = infoblox_ipv4_network.net1.network_view
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_ipv4_network.net1, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Name:   utils.StringPtr("name3.test.com"),
						View:   "nondefault_view",
						UseTtl: utils.BoolPtr(false),
					}, "10.10.0.1", "10.20.30.0/24", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
			        resource "infoblox_ipv4_network" "netA" {
			            cidr = "11.20.0.0/24"
			            network_view = "default"
			        }
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
			            cidr = infoblox_ipv4_network.netA.cidr
						ip_addr = "10.10.0.7"
			            network_view = infoblox_ipv4_network.netA.network_view
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_ipv4_network.netA, infoblox_dns_view.view1]
					}`),
				ExpectError: regexpUpdateConflictIPv4,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
			        resource "infoblox_ipv4_network" "net2" {
			            cidr = "10.20.33.0/24"
			            network_view = "default"
			        }
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
			            cidr = infoblox_ipv4_network.net2.cidr
			            network_view = infoblox_ipv4_network.net2.network_view
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_ipv4_network.net2, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Name:   utils.StringPtr("name3.test.com"),
						View:   "nondefault_view",
						UseTtl: utils.BoolPtr(false),
					}, "", "10.20.33.0/24", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
					resource "infoblox_network_view" "view1" {
						name = "nondefault_netview"
					}
			      	resource "infoblox_ipv4_network" "net3" {
			          	cidr = "10.20.34.0/24"
			          	network_view = infoblox_network_view.view1.name
			      	}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
			          	cidr = infoblox_ipv4_network.net3.cidr
			          	network_view = infoblox_ipv4_network.net3.network_view
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_network_view.view1, infoblox_ipv4_network.net3]
					}`),
				ExpectError: regexpNetviewUpdateNotAllowed,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view1" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
						view = infoblox_dns_view.view1.name
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.2"
						dns_view = "nondefault_view"
						depends_on = [infoblox_zone_auth.zone1, infoblox_dns_view.view1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.2"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						UseTtl:   utils.BoolPtr(false),
					}, "", "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.2"
						dns_view = "default"
						depends_on = [infoblox_zone_auth.zone1]
					}`),
				ExpectError: regexpDnsviewUpdateNotAllowed,
			},
			{
				Config: fmt.Sprintf(`
				resource "infoblox_zone_auth" "zone1" {
					fqdn = "test1.com"
				}
				resource "infoblox_ipv4_network" "net2" {
					cidr = "10.1.0.0/24"
					ext_attrs = jsonencode({
 						"Site" = "Blr"
					})
				}
				resource "infoblox_a_record" "rec4" {
					fqdn = "dynamic.test1.com"
					filter_params = jsonencode({
						"*Site" = "Blr"})
					dns_view = "default"
				}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.rec4", &ibclient.RecordA{
						Name:   utils.StringPtr("dynamic.test1.com"),
						View:   "default",
						UseTtl: utils.BoolPtr(false),
					}, "", "", `{"*Site":"Blr"}`),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "infoblox_zone_auth" "zone1" {
					fqdn = "test1.com"
				}
				resource "infoblox_a_record" "rec1" {
					fqdn = "missing_fields.test1.com"
					comment = "missing required fields"
					dns_view = "default"
					depends_on = [infoblox_zone_auth.zone1]
				}`),
				ExpectError: regexpRequiredMissingIPv4,
			},
		},
	})
}

func TestAcc_resourceARecord_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckARecordDestroy,

		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
				resource "infoblox_a_record" "foo3"{
					dns_view = "default"
					fqdn = "samplearec2.test.com"
					ip_addr = "10.1.0.2"
					comment = "test comment on A record"
					ext_attrs = jsonencode({
						"Location" = "test location A"
					})
					depends_on = [infoblox_zone_auth.zone1]
				}`,
				Check: testAccARecordCompare(t, "infoblox_a_record.foo3", &ibclient.RecordA{
					Ipv4Addr: utils.StringPtr("10.1.0.2"),
					Name:     utils.StringPtr("samplearec2.test.com"),
					View:     "default",
					UseTtl:   utils.BoolPtr(false),
					Comment:  utils.StringPtr("test comment on A record"),
					Ea: ibclient.EA{
						"Location": "test location A",
					},
				}, "", "", ""),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.RecordA{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"name":     "samplearec2.test.com",
							"ipv4addr": "10.1.0.2",
						},
					)
					var res []ibclient.RecordA
					err := conn.GetObject(n, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].View = ""
					res[0].Ea["Site"] = "Test site"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
				resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
				resource "infoblox_a_record" "foo3"{
					dns_view = "default"
					fqdn = "samplearec2.test.com"
					ip_addr = "10.1.0.2"
					comment = "test comment on A record"
					ext_attrs = jsonencode({
						"Location" = "test location A"
					})
					depends_on = [infoblox_zone_auth.zone1]
				}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_a_record.foo3", "ext_attrs",
						`{"Location":"test location A"}`,
					),
					// Actual API object should have Site EA
					testAccARecordCompare(t, "infoblox_a_record.foo3", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.1.0.2"),
						Name:     utils.StringPtr("samplearec2.test.com"),
						View:     "default",
						UseTtl:   utils.BoolPtr(false),
						Comment:  utils.StringPtr("test comment on A record"),
						Ea: ibclient.EA{
							"Location": "test location A",
							"Site":     "Test site",
						},
					}, "", "", ""),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
				resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
				resource "infoblox_a_record" "foo3"{
					dns_view = "default"
					fqdn = "samplearec2.test.com"
					ip_addr = "10.1.0.2"
					comment = "updated comment on A record"
					ext_attrs = jsonencode({
						"Location" = "test location A"
					})
					depends_on = [infoblox_zone_auth.zone1]
				}`,
				Check: testAccARecordCompare(t, "infoblox_a_record.foo3", &ibclient.RecordA{
					Ipv4Addr: utils.StringPtr("10.1.0.2"),
					Name:     utils.StringPtr("samplearec2.test.com"),
					View:     "default",
					UseTtl:   utils.BoolPtr(false),
					Comment:  utils.StringPtr("updated comment on A record"),
					Ea: ibclient.EA{
						"Location": "test location A",
						"Site":     "Test site",
					},
				}, "", "", ""),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
				resource "infoblox_a_record" "foo3"{
					dns_view = "default"
					fqdn = "samplearec2.test.com"
					ip_addr = "10.1.0.2"
					comment = "test comment on A record"
					ext_attrs = jsonencode({
						"Location" = "test location A"
						"Site" = "new extensible site"
					})
					depends_on = [infoblox_zone_auth.zone1]
				}`,
				Check: testAccARecordCompare(t, "infoblox_a_record.foo3", &ibclient.RecordA{
					Ipv4Addr: utils.StringPtr("10.1.0.2"),
					Name:     utils.StringPtr("samplearec2.test.com"),
					View:     "default",
					UseTtl:   utils.BoolPtr(false),
					Comment:  utils.StringPtr("test comment on A record"),
					Ea: ibclient.EA{
						"Location": "test location A",
						"Site":     "new extensible site",
					},
				}, "", "", ""),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
				resource "infoblox_zone_auth" "zone1" {
						fqdn = "test.com"
					}
				resource "infoblox_a_record" "foo3"{
					dns_view = "default"
					fqdn = "samplearec2.test.com"
					ip_addr = "10.1.0.2"
					comment = "test comment on A record"
					ext_attrs = jsonencode({
						"Location" = "test location A"
					})
					depends_on = [infoblox_zone_auth.zone1]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_a_record.foo3", "ext_attrs",
						`{"Location":"test location A"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_a_record.foo3"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_a_record.foo3")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						arec, err := objMgr.GetARecordByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := arec.Ea["Site"]; ok {
							return fmt.Errorf("Site EA should've been removed, but still present in the WAPI object")
						}
						return nil
					},
				),
			},
		},
	})
}
