# Auto-generated list acceptance-test cases for RecordNs.
case "basic" {
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{ address = "20.0.0.0", auto_create_ptr = false }]
      view       = "default"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{ address = "20.0.0.0", auto_create_ptr = false }]
      view       = "default"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name       = "nios.name"
        nameserver = "nios.nameserver"
      }
    }
  }

}

case "creator_filter" {
  backend = "nios"
  min_tf_version = "1.14.0"
  skip        = true
  skip_reason = "unmappable filter key creator"
}
