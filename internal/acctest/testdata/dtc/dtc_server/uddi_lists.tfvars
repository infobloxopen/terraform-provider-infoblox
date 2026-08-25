# DtcServer — uddi list cases

case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
      tags          = { env = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        env = "uddi.tags.env"
      }
    }
  }

}
