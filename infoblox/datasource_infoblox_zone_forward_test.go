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
}`)

var testDataSourceZoneForwardEA = fmt.Sprintf(
	`resource "infoblox_zone_forward" "zf_data_src_ea" {
    fqdn = "test_fz_ea.ex.org"
    comment = "test sample forward zone with EA"
    ns_group = "foward_ns_group"
    forward_to {
        name = "test_ds_123.ea.ex.com"
        address = "10.0.0.1"
    }
    forward_to {
        name = "test_ds_245.ea.ex.com"
        address = "10.0.0.2"
    }
	extattrs = jsonencode({
        Site = "Main DNS Site"
    })
}
data "infoblox_zone_forward" "zf_data_src_ea_read" {	
	filters = {
	fqdn = infoblox_zone_forward.zf_data_src_ea.fqdn
}`)

func TestAccDataSourceZoneForward(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceZoneForward,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "fqdn", "test_fz_ds.ex.org"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "comment", "test sample forward zone"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "forward_to.0.name", "test_ds_123.dz.ex.com"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "forward_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "forward_to.1.name", "test_ds_245.dz.ex.com"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "forward_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "ns_group", ""),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "forwarders_only", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_read", "disable", "false"),
				),
			},
			{
				Config: testDataSourceZoneForwardEA,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "fqdn", "test_fz_ea.ex.org"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "comment", "test sample forward zone with EA"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "forward_to.0.name", "test_ds_123.ea.ex.com"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "forward_to.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "forward_to.1.name", "test_ds_245.ea.ex.com"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "forward_to.1.address", "10.0.0.2"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "ns_group", "foward_ns_group"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "forwarders_only", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "disable", "false"),
					resource.TestCheckResourceAttr("infoblox_zone_forward.zf_data_src_ea_read", "extattrs.Site", "Main DNS Site"),
				),
			},
		},
	})
}
