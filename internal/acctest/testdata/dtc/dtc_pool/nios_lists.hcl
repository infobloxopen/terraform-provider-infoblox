# DtcPool — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
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
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
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

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ext_attrs           = { Site = "{{random2}}" }
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
