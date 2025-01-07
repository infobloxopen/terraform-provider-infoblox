package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testAccDataSourceDtcLbdn = fmt.Sprintf(`resource "infoblox_dtc_lbdn" "testLbdn_src_1" {
    name = "testLbdn444"
  	lb_method = "RATIO"
    }
    data "infoblox_dtc_lbdn" "testLbdn_src_read1" {	
	filters = {
	    name = infoblox_dtc_lbdn.testLbdn_src_1.name
    }
    depends_on = [infoblox_dtc_lbdn.testLbdn_src_1]
}`)

var testAccDatasourceDtcLbdn = fmt.Sprintf(`resource "infoblox_dtc_lbdn" "testLbdn_src" {
    name = "testLbdn888"
    auth_zones = ["test.com"]
  	comment = "test lbdn with max params"
  	ext_attrs = jsonencode({
        "Site" = "Malpe"
  	})
  	lb_method = "TOPOLOGY"
  	patterns = ["test.com*", "info.com*"]
  	pools {
    	pool = "pool2"
    	ratio = 2
  	}
  	pools {
    	pool = "rrpool"
    	ratio = 3
  	}
  	pools {
    	pool = "test-pool"
    	ratio = 6
  	}
    topology = "test-topo"
  	ttl = 120
  	disable = true
  	types = ["A", "AAAA", "CNAME"]
  	persistence = 60
  	priority = 2
}
data "infoblox_dtc_lbdn" "testLbdn_src_read" {	
	filters = {
	    name = infoblox_dtc_lbdn.testLbdn_src.name
    }
    depends_on = [infoblox_dtc_lbdn.testLbdn_src]
}`)

func TestAccDataSourceDtcLbdn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDtcLbdn,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read1", "results.0.name", "testLbdn444"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read1", "results.0.lb_method", "RATIO"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read1", "results.0.types.0", "A"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read1", "results.0.types.1", "AAAA"),
				),
			},
		},
	})
}

func TestAccDataSourceDtcLbdnSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceDtcLbdn,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.name", "testLbdn888"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.auth_zones.0", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.comment", "test lbdn with max params"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.lb_method", "TOPOLOGY"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.topology", "test-topo"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.0.pool", "pool2"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.0.ratio", "2"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.1.pool", "rrpool"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.1.ratio", "3"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.2.pool", "test-pool"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.pools.2.ratio", "6"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.ttl", "120"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.disable", "true"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.types.0", "A"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.types.1", "AAAA"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.types.2", "CNAME"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.persistence", "60"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.priority", "2"),
					resource.TestCheckResourceAttrPair("data.infoblox_dtc_lbdn.testLbdn_src_read", "results.0.ext_attrs.Site", "infoblox_dtc_lbdn.testLbdn_src", "ext_attrs.Site"),
				),
			},
		},
	})
}
