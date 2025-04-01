package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testAccDataSourceAliasRecord = fmt.Sprintf(`resource "infoblox_alias_record" "record1" {
    name = "alias_read1.test.com"
    target_name = "aa.bb.com"
	target_type = "AAAA"
}
    data "infoblox_alias_record" "read_alias" {	
	filters = {
	    name = infoblox_alias_record.record1.name
		view = infoblox_alias_record.record1.dns_view 
    }
}`)

var testAccDataSourceAliasRecord1 = fmt.Sprintf(`resource "infoblox_alias_record" "record2" {
    name = "alias_read2.test.com"
    comment = "test alias record"
    target_name = "bb.kk.com"
	target_type = "NAPTR"
	dns_view = "default"
	ttl = 36000
	disable = true
	ext_attrs = jsonencode({
    	"Site" = "Romania"
  	})
}
data "infoblox_alias_record" "read_alias2" {	
	filters = {
	    name = infoblox_alias_record.record2.name
    }
}`)

func TestAccDataSourceAliasRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAliasRecord,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias", "results.0.name", "alias_read1.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias", "results.0.target_name", "aa.bb.com"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias", "results.0.target_type", "AAAA"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias", "results.0.dns_view", "default"),
				),
			},
		},
	})
}

func TestAccDataSourceAliasRecordEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAliasRecord1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.name", "alias_read2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.comment", "test alias record"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.target_name", "bb.kk.com"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.target_type", "NAPTR"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.ttl", "36000"),
					resource.TestCheckResourceAttr("data.infoblox_alias_record.read_alias2", "results.0.disable", "true"),
					resource.TestCheckResourceAttrPair("data.infoblox_alias_record.read_alias2", "results.0.ext_attrs.Site", "infoblox_alias_record.record2", "ext_attrs.Site"),
				),
			},
		},
	})
}
