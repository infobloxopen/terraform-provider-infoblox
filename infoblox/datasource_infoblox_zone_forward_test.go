package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testDataSourceZoneForward = fmt.Sprintf(
	`resource "infoblox_zone_forward" "zf_data_src" {
    fqdn = "test_fz_ds.ex.org"
    comment = "test sample forward zone"
    forward_to {
        name = "test_ds_123.dz.ex.com"
        address = "10.0.0.1"
    }
    forward_to {
        name = "test_ds_245.dz.ex.com"
        address = "10.0.0.2"
    }
}
data "infoblox_zone_forward" "zf_data_src_read" {	
	filters = {
	    fqdn = infoblox_zone_forward.zf_data_src.fqdn
    }
    depends_on = [infoblox_zone_forward.zf_data_src]
}`)

var testDataSourceZoneForwardEA = fmt.Sprintf(
	`resource "infoblox_zone_forward" "zf_data_src_ea" {
    fqdn = "test2.ex.org"
    comment = "test sample forward zone with EA"
    forward_to {
        name = "test_ds_123.ea.ex.com"
        address = "10.0.0.1"
    }
    forward_to {
        name = "test_ds_245.ea.ex.com"
        address = "10.0.0.2"
    }
	ext_attrs = jsonencode({
        "Location" = "TBD"
    })
}
data "infoblox_zone_forward" "zf_data_src_ea_read" {	
	filters = {
	    "*Location" = "TBD"
    }
    depends_on = [infoblox_zone_forward.zf_data_src_ea]
}`)

func TestAccDataSourceZoneForward(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testZoneForwardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceZoneForward,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.fqdn", "test_fz_ds.ex.org"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.comment", "test sample forward zone"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.forward_to.0.name", "test_ds_123.dz.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.forward_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.forward_to.1.name", "test_ds_245.dz.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.forward_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.ns_group", ""),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.forwarders_only", "false"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_read", "results.0.disable", "false"),
				),
			},
			{
				Config: testDataSourceZoneForwardEA,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.fqdn", "test2.ex.org"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.comment", "test sample forward zone with EA"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.forward_to.0.name", "test_ds_123.ea.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.forward_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.forward_to.1.name", "test_ds_245.ea.ex.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.forward_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.forwarders_only", "false"),
					resource.TestCheckResourceAttr("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.disable", "false"),
					resource.TestCheckResourceAttrPair("data.infoblox_zone_forward.zf_data_src_ea_read", "results.0.ext_attrs.Location", "infoblox_zone_forward.zf_data_src_ea", "ext_attrs.Location"),
				),
			},
		},
	})
}
