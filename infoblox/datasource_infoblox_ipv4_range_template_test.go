package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testDataSourceRangeTemplate = fmt.Sprintf(
	`resource "infoblox_ipv4_range_template" "range_template" {
    name = "rangetemplate111"
    number_of_addresses = 10
    cloud_api_compatible = true
    offset = 70
}
data "infoblox_ipv4_range_template" "range_template_read" {	
	filters = {
	    name = infoblox_ipv4_range_template.range_template.name
    }
}`)

var testDataSourceRangeTemplateEA = fmt.Sprintf(
	`resource "infoblox_ipv4_range_template" "range_template_ea" {
    name = "range-template333"
    number_of_addresses = 60
	cloud_api_compatible = true
    offset = 76
    comment = "Temporary Range Template"
    use_options = true
    ext_attrs = jsonencode({
      "Site" = "Kobe"
    })
    options {
		name = "domain-name-servers"
		value = "11.22.33.44"
		vendor_class = "DHCP"
		num = 6
		use_option = true
	}
}
data "infoblox_ipv4_range_template" "range_template_ea_read" {	
	filters = {
	    "*Site" = "Kobe"
    }
    depends_on = [infoblox_ipv4_range_template.range_template_ea]
}`)

func TestAccDataSourceRangeTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRangeTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceRangeTemplate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.name", "rangetemplate111"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.number_of_addresses", "10"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.offset", "70"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.options.0.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.options.0.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_read", "results.0.options.0.use_option", "false"),
				),
			},
			{
				Config: testDataSourceRangeTemplateEA,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.name", "range-template333"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.comment", "Temporary Range Template"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.number_of_addresses", "60"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.offset", "76"),

					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.0.name", "domain-name-servers"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.0.num", "6"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.0.value", "11.22.33.44"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.0.use_option", "true"),

					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.1.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.1.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.1.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.1.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.options.1.use_option", "false"),
					resource.TestCheckResourceAttrPair("data.infoblox_ipv4_range_template.range_template_ea_read", "results.0.ext_attrs.Site", "infoblox_ipv4_range_template.range_template_ea", "ext_attrs.Site"),
				),
			},
		},
	})
}
