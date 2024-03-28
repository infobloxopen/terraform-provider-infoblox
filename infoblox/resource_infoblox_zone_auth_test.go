package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"testing"
)

func testAccCheckZoneAuthDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_zone_auth" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetZoneAuthByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("zone not found")
		}

	}
	return nil
}

func testAccZoneAuthCompare(t *testing.T, resPath string, expectedObj *ibclient.ZoneAuth) resource.TestCheckFunc {
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
		zone, err := objMgr.SearchObjectByAltId("ZoneAuth", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedObj == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var zoneObj *ibclient.ZoneAuth
		recJson, _ := json.Marshal(zone)
		err = json.Unmarshal(recJson, &zoneObj)

		if zoneObj.Fqdn != expectedObj.Fqdn {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'", zoneObj.Fqdn,
				expectedObj.Fqdn)
		}

		if zoneObj.View == nil {
			return fmt.Errorf("'view' is expected to be defined but it is not")
		}
		if *zoneObj.View != *expectedObj.View {
			return fmt.Errorf(
				"'view' does not match: got '%s', expected '%s'",
				*zoneObj.View, *expectedObj.View)
		}

		if expectedObj.ZoneFormat != "" {
			if zoneObj.ZoneFormat != expectedObj.ZoneFormat {
				return fmt.Errorf(
					"'zone_format' does not match: got '%s', expected '%s'", zoneObj.ZoneFormat,
					expectedObj.ZoneFormat)
			}
		}

		if zoneObj.SoaDefaultTtl != nil {
			if *zoneObj.SoaDefaultTtl != *expectedObj.SoaDefaultTtl {
				return fmt.Errorf(
					"'soa_default_ttl' does not match: got '%b', expected '%b'",
					*zoneObj.SoaDefaultTtl,
					*expectedObj.SoaDefaultTtl)
			}
		}

		if zoneObj.SoaExpire != nil {
			if *zoneObj.SoaExpire != *expectedObj.SoaExpire {
				return fmt.Errorf(
					"'soa_expire' does not match: got '%b', expected '%b'",
					*zoneObj.SoaExpire,
					*expectedObj.SoaExpire)
			}
		}

		if zoneObj.SoaNegativeTtl != nil {
			if *zoneObj.SoaNegativeTtl != *expectedObj.SoaNegativeTtl {
				return fmt.Errorf(
					"'soa_negative_ttl' does not match: got '%b', expected '%b'",
					*zoneObj.SoaNegativeTtl,
					*expectedObj.SoaNegativeTtl)
			}
		}

		if zoneObj.SoaRefresh != nil {
			if *zoneObj.SoaRefresh != *expectedObj.SoaRefresh {
				return fmt.Errorf(
					"'soa_refresh' does not match: got '%b', expected '%b'",
					*zoneObj.SoaRefresh,
					*expectedObj.SoaRefresh)
			}
		}

		if zoneObj.SoaRetry != nil {
			if *zoneObj.SoaRetry != *expectedObj.SoaRetry {
				return fmt.Errorf(
					"'soa_retry' does not match: got '%b', expected '%b'",
					*zoneObj.SoaRetry,
					*expectedObj.SoaRetry)
			}
		}

		if zoneObj.Comment != nil {
			if expectedObj.Comment == nil {
				return fmt.Errorf("'comment' is expected to be defined but it is not")
			}
			if *zoneObj.Comment != *expectedObj.Comment {
				return fmt.Errorf(
					"'comment' does not match: got '%s', expected '%s'",
					*zoneObj.Comment, *expectedObj.Comment)
			}
		}

		return validateEAs(zoneObj.Ea, expectedObj.Ea)
	}
}

func TestAccResourceZoneAuthBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneAuthDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone2" {
						fqdn = "test2.com"
					}
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
						//ns_group = "nsgroup1"
						restart_if_needed = true
						soa_default_ttl = 36000
						soa_expire = 72000
						soa_negative_ttl = 600
						soa_refresh = 1800
						soa_retry = 900
						comment = "Zone Auth created by terraform acceptance test"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform"
						})
						depends_on = [infoblox_dns_view.view]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "fqdn", "test3.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "view", "nondefault_view"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "comment", "Zone Auth created by terraform acceptance test"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ext_attrs", "{\"Location\":\"AcceptanceTerraform\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_default_ttl", "36000"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_expire", "72000"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_negative_ttl", "600"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_refresh", "1800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_retry", "900"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone2" {
						fqdn = "test2.com"
						//ns_group = "nsgroup2"
						restart_if_needed = true
						soa_default_ttl = 36002
						soa_expire = 72002
						soa_negative_ttl = 602
						soa_refresh = 1802
						soa_retry = 902
						comment = "Zone Auth created by terraform acceptance test 22"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform 22"
						})
					}
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
						//ns_group = "nsgroup2"
						restart_if_needed = false
						soa_default_ttl = 36001
						soa_expire = 72001
						soa_negative_ttl = 601
						soa_refresh = 1801
						soa_retry = 901
						comment = "Zone Auth created by terraform acceptance test 2"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform 2"
						})
						depends_on = [infoblox_dns_view.view]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "comment", "Zone Auth created by terraform acceptance test 22"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ext_attrs", "{\"Location\":\"AcceptanceTerraform 22\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "restart_if_needed", "true"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_default_ttl", "36002"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_expire", "72002"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_negative_ttl", "602"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_refresh", "1802"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_retry", "902"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "fqdn", "test3.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "view", "nondefault_view"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "comment", "Zone Auth created by terraform acceptance test 2"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ext_attrs", "{\"Location\":\"AcceptanceTerraform 2\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "restart_if_needed", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_default_ttl", "36001"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_expire", "72001"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_negative_ttl", "601"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_refresh", "1801"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_retry", "901"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "view" {
						name = "nondefault_view"
					}
					resource "infoblox_zone_auth" "test_zone2" {
						fqdn = "test2.com"
						//ns_group = "nsgroup2"
					}
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
						depends_on = [infoblox_dns_view.view]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "fqdn", "test3.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "view", "nondefault_view"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "soa_retry", "3600"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone4" {
						fqdn = "10.20.30.0/24"
						zone_format = "IPV4"
					}
					resource "infoblox_zone_auth" "test_zone5" {
						fqdn = "2345::/64"
						zone_format = "IPV6"
						//ns_group = "nsgroup1"
						restart_if_needed = true
						soa_default_ttl = 36000
						soa_expire = 72000
						soa_negative_ttl = 600
						soa_refresh = 1800
						soa_retry = 900
						comment = "Zone Auth created by terraform acceptance test"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "fqdn", "10.20.30.0/24"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "zone_format", "IPV4"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "fqdn", "2345::/64"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "zone_format", "IPV6"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "comment", "Zone Auth created by terraform acceptance test"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ext_attrs", "{\"Location\":\"AcceptanceTerraform\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "restart_if_needed", "true"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_default_ttl", "36000"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_expire", "72000"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_negative_ttl", "600"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_refresh", "1800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_retry", "900"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone4" {
						fqdn = "10.20.30.0/24"
						zone_format = "IPV4"
						//ns_group = "nsgroup2"
						restart_if_needed = true
						soa_default_ttl = 36002
						soa_expire = 72002
						soa_negative_ttl = 602
						soa_refresh = 1802
						soa_retry = 902
						comment = "Zone Auth created by terraform acceptance test 22"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform 22"
						})
					}
					resource "infoblox_zone_auth" "test_zone5" {
						fqdn = "2345::/64"
						zone_format = "IPV6"
						//ns_group = "nsgroup2"
						restart_if_needed = false
						soa_default_ttl = 36001
						soa_expire = 72001
						soa_negative_ttl = 601
						soa_refresh = 1801
						soa_retry = 901
						comment = "Zone Auth created by terraform acceptance test 2"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform 2"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "fqdn", "10.20.30.0/24"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "zone_format", "IPV4"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "comment", "Zone Auth created by terraform acceptance test 22"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ext_attrs", "{\"Location\":\"AcceptanceTerraform 22\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "restart_if_needed", "true"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_default_ttl", "36002"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_expire", "72002"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_negative_ttl", "602"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_refresh", "1802"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_retry", "902"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "fqdn", "2345::/64"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "zone_format", "IPV6"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "comment", "Zone Auth created by terraform acceptance test 2"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ext_attrs", "{\"Location\":\"AcceptanceTerraform 2\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "restart_if_needed", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_default_ttl", "36001"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_expire", "72001"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_negative_ttl", "601"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_refresh", "1801"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_retry", "901"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone4" {
						fqdn = "10.20.30.0/24"
						zone_format = "IPV4"
						//ns_group = "nsgroup2"
					}
					resource "infoblox_zone_auth" "test_zone5" {
						fqdn = "2345::/64"
						zone_format = "IPV6"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "fqdn", "10.20.30.0/24"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "zone_format", "IPV4"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "fqdn", "2345::/64"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "zone_format", "IPV6"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "soa_retry", "3600"),
				),
			},
		},
	})
}

func TestAcc_resourceZoneAuth_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneAuthDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone6" {
						fqdn = "10.10.0.0/24"
						view = "default"
						zone_format = "IPV4"
						comment = "test reverse mapping zone"
						ext_attrs = jsonencode({
							"Location" = "Test RM location"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					testAccZoneAuthCompare(t, "infoblox_zone_auth.test_zone6", &ibclient.ZoneAuth{
						View:    utils.StringPtr("default"),
						Fqdn:    "10.10.0.0/24",
						Comment: utils.StringPtr("test reverse mapping zone"),
						Ea: ibclient.EA{
							"Location": "Test RM location",
						},
					}),
				),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					zobj := &ibclient.ZoneAuth{}
					zobj.SetReturnFields(append(zobj.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"fqdn": "10.10.0.0/24",
							"view": "default",
						},
					)

					var res []ibclient.ZoneAuth
					err := conn.GetObject(zobj, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].Ea["Site"] = "Test RM Site"
					res[0].Fqdn = ""

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone6" {
						fqdn = "10.10.0.0/24"
						view = "default"
						zone_format = "IPV4"
						comment = "test reverse mapping zone"
						ext_attrs = jsonencode({
							"Location" = "Test RM location"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Site EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_zone_auth.test_zone6", "ext_attrs",
						`{"Location":"Test RM location"}`,
					),
					// Actual API object should have Site EA
					testAccZoneAuthCompare(t, "infoblox_zone_auth.test_zone6", &ibclient.ZoneAuth{
						View:    utils.StringPtr("default"),
						Fqdn:    "10.10.0.0/24",
						Comment: utils.StringPtr("test reverse mapping zone"),
						Ea: ibclient.EA{
							"Location": "Test RM location",
							"Site":     "Test RM Site",
						},
					}),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone6" {
						fqdn = "10.10.0.0/24"
						view = "default"
						zone_format = "IPV4"
						comment = "updated reverse mapping zone"
						ext_attrs = jsonencode({
							"Location" = "Test RM location"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					testAccZoneAuthCompare(t, "infoblox_zone_auth.test_zone6", &ibclient.ZoneAuth{
						View:    utils.StringPtr("default"),
						Fqdn:    "10.10.0.0/24",
						Comment: utils.StringPtr("updated reverse mapping zone"),
						Ea: ibclient.EA{
							"Location": "Test RM location",
							"Site":     "Test RM Site",
						},
					}),
				),
			},
			// Validate that inherited EA can be updated
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone6" {
						fqdn = "10.10.0.0/24"
						view = "default"
						zone_format = "IPV4"
						comment = "test reverse mapping zone"
						ext_attrs = jsonencode({
							"Location" = "Test RM location"
							"Site" = "New RM Site"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					testAccZoneAuthCompare(t, "infoblox_zone_auth.test_zone6", &ibclient.ZoneAuth{
						View:    utils.StringPtr("default"),
						Fqdn:    "10.10.0.0/24",
						Comment: utils.StringPtr("test reverse mapping zone"),
						Ea: ibclient.EA{
							"Location": "Test RM location",
							"Site":     "New RM Site",
						},
					}),
				),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone6" {
						fqdn = "10.10.0.0/24"
						view = "default"
						zone_format = "IPV4"
						comment = "updated reverse mapping zone"
						ext_attrs = jsonencode({
							"Location" = "Test RM location"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_zone_auth.test_zone6", "ext_attrs",
						`{"Location":"Test RM location"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_zone_auth.test_zone6"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_zone_auth.test_zone6")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						zobj, err := objMgr.GetZoneAuthByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}
						if _, ok := zobj.Ea["Site"]; ok {
							return fmt.Errorf("Site EA should've been removed, but still present in the WAPI object")
						}
						return nil
					},
				),
			},
		},
	})
}
