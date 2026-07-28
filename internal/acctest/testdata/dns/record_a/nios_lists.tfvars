# Auto-generated list acceptance-test cases for RecordA (nios).
case "basic" {
  # basic — generated from terraform-provider-nios
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  # filters — generated from terraform-provider-nios
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.21"
      view     = "default"
    }
  }

  step {
    query    = true
    provider = infoblox
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
  # ext_attr_filters — generated from terraform-provider-nios
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name      = "{{random}}.example.com"
      ipv4addr  = "10.0.0.22"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
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
