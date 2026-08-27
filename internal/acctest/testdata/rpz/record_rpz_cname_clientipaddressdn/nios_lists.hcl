case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
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
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "{{random}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
