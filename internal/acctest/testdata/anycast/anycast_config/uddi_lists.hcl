case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name               = "{{random}}"
              service            = "NTP"
              anycast_ip_address = "{{random_ip}}"
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
  min_tf_version = "1.14.0"

  step {
      uddi {
        name               = "{{random}}"
                service            = "NTP"
                anycast_ip_address = "{{random_ip}}"
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
  min_tf_version = "1.14.0"

  step {
    uddi {
      name               = "{{random}}"
              service            = "NTP"
              anycast_ip_address = "{{random_ip}}"
      tags = { tag1 = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
