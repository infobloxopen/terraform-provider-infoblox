# RecordRpzTxt — nios list cases
# TODO : rp_zone must be present
# test-rpz.com

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "List test text"
      rp_zone = "test-rpz.com"
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
      name    = "{{random2}}.test-rpz.com"
      text    = "Filter test text"
      rp_zone = "test-rpz.com"
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
      name      = "{{random2}}.test-rpz.com"
      text      = "Ext attr filter text"
      rp_zone   = "test-rpz.com"
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
