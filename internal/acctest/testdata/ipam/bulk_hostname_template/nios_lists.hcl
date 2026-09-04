# Auto-generated list acceptance-test cases for Bulkhostnametemplate.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "host-$4"
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
      template_name   = "{{random}}"
      template_format = "host-$4"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        template_name = "nios.template_name"
      }
    }
  }

}
