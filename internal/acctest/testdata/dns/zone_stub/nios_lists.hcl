# Auto-generated list acceptance-test cases for ZoneStub.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
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
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        fqdn = "nios.fqdn"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ext_attrs = { Site = "{{random3}}" }
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
