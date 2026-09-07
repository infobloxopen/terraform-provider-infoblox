# Hand-authored list acceptance-test cases for RecordRpzAaaaIpaddress.
# rp_zone is hardcoded to "rpz-test.infoblox.com" (persistent zone on the test NIOS grid)
# because infoblox_zone_rp is not yet registered in the unified provider.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true

  step {
    nios {
      name      = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr  = "2001:db8::10"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "{{random3}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
