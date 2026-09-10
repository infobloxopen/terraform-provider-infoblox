case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
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
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      outbound_member_type = "GM"
      uri                  = "https://example.com"
      ext_attrs            = { Site = "{{random2}}" }
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
