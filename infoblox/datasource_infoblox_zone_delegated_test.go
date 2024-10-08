package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testDataSourceZoneDelegated = fmt.Sprintf(

	`resource "infoblox_zone_auth" "zone11" {
  fqdn = "test33.com"
  view = "default"
  zone_format = "FORWARD"
  ns_group = ""
  restart_if_needed = true
  soa_default_ttl = 36000
  soa_expire = 72000
  soa_negative_ttl = 600
  soa_refresh = 1800
  soa_retry = 900
  comment = "Zone Auth created newly"
  ext_attrs = jsonencode({
    Location = "AcceptanceTerraform"
  })
}

resource "infoblox_zone_delegated" "zone_delegated_data_src" {
    fqdn = "test_fz_ds.test33.com"
    comment = "test sample delegate zone"
    delegate_to {
        name = "test_ds_123.dz.ex.com"
        address = "10.0.0.1"
    }
    delegate_to {
        name = "test_ds_245.dz.ex.com"
        address = "10.0.0.2"
    }
	depends_on = [infoblox_zone_auth.zone11]
}
data "infoblox_zone_delegated" "zone_delegated_data_src_read" {	
	filters = {
	    fqdn = infoblox_zone_delegated.zone_delegated_data_src.fqdn
    }
    depends_on = [infoblox_zone_delegated.zone_delegated_data_src]
}`)

var testDataSourceZoneDelegatedEA = fmt.Sprintf(
	`resource "infoblox_zone_auth" "zone12" {
fqdn = "test313.com"
view = "default"
zone_format = "FORWARD"
ns_group = ""
restart_if_needed = true
soa_default_ttl = 36000
soa_expire = 72000
soa_negative_ttl = 600
soa_refresh = 1800
soa_retry = 900
comment = "Zone Auth created newly"
ext_attrs = jsonencode({
Location = "AcceptanceTerraform"
})
}

resource "infoblox_zone_delegated" "zone_delegated_data_src_ea" {
    fqdn = "test2.test313.com"
    comment = "test sample delegate zone with EA"
    delegate_to {
        name = "test_ds_123.ea.ex.com"
        address = "10.0.0.1"
    }
    delegate_to {
        name = "test_ds_245.ea.ex.com"
        address = "10.0.0.2"
    }
	ext_attrs = jsonencode({
        "Location" = "TBD"
    })
	depends_on = [infoblox_zone_auth.zone12]
}

data "infoblox_zone_delegated" "zone_delegated_data_src_ea_read" {	
	filters = {
	    "*Location" = "TBD"
    }
    depends_on = [infoblox_zone_delegated.zone_delegated_data_src_ea]
}`)

func TestAccDataSourceZoneDelegated(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDelegatedDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceZoneDelegated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.fqdn", "test_fz_ds.test33.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.comment", "test sample delegate zone"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.delegate_to.0.name", "test_ds_123.dz.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.delegate_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.delegate_to.1.name", "test_ds_245.dz.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.delegate_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.ns_group", ""),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_read", "results.0.disable", "false"),
				),
			},
			{
				Config: testDataSourceZoneDelegatedEA,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.fqdn", "test2.test313.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.comment", "test sample delegate zone with EA"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.delegate_to.0.name", "test_ds_123.ea.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.delegate_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.delegate_to.1.name", "test_ds_245.ea.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.delegate_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.disable", "false"),
					resource.TestCheckResourceAttrPair("data.infoblox_zone_delegated.zone_delegated_data_src_ea_read", "results.0.ext_attrs.Location", "infoblox_zone_delegated.zone_delegated_data_src_ea", "ext_attrs.Location"),
				),
			},
		},
	})
}
