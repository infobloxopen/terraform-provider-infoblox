package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceZoneAuthBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCNAMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone2" {
						fqdn = "test2.com"
					}
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
						ns_group = "nsgroup1"
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ext_attrs", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ns_group", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "restart_if_needed", "false"),
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ns_group", "nsgroup1"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "restart_if_needed", "true"),
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
						ns_group = "nsgroup2"
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
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
						ns_group = "nsgroup2"
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "comment", "Zone Auth created by terraform acceptance test 22"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ext_attrs", "{\"Location\":\"AcceptanceTerraform 22\"}"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ns_group", "nsgroup2"),
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ns_group", "nsgroup2"),
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
					resource "infoblox_zone_auth" "test_zone2" {
						fqdn = "test2.com"
						ns_group = "nsgroup2"
					}
					resource "infoblox_zone_auth" "test_zone3" {
						fqdn = "test3.com"
						view = "nondefault_view"
						zone_format = "FORWARD"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "fqdn", "test2.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "zone_format", "FORWARD"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ext_attrs", ""),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "ns_group", "nsgroup2"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "restart_if_needed", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone2", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "fqdn", "test3.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "view", "nondefault_view"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "zone_format", "FORWARD"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ext_attrs", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "ns_group", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone3", "restart_if_needed", "false"),
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
						ns_group = "nsgroup1"
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
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ext_attrs", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ns_group", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "restart_if_needed", "false"),
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ns_group", "nsgroup1"),
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
						ns_group = "nsgroup2"
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
						ns_group = "nsgroup2"
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ns_group", "nsgroup2"),
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
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ns_group", "nsgroup2"),
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
						ns_group = "nsgroup2"
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
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ext_attrs", ""),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "ns_group", "nsgroup2"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "restart_if_needed", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_default_ttl", "28800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_expire", "2419200"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_negative_ttl", "900"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_refresh", "10800"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone4", "soa_retry", "3600"),

					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "fqdn", "2345::/64"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "zone_format", "IPV6"),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "comment", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ext_attrs", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "ns_group", ""),
					//resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone5", "restart_if_needed", "false"),
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
