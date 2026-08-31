# HaGroup — uddi list cases
#  TODO: Objects to be present in the grid for testing
#  dhcp/host/470520
#  dhcp/host/470521
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      hosts = [
        { host = "dhcp/host/470520", role = "active" },
        { host = "dhcp/host/470521", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      hosts = [
        { host = "dhcp/host/470520", role = "active" },
        { host = "dhcp/host/470521", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      hosts = [
        { host = "dhcp/host/470520", role = "active" },
        { host = "dhcp/host/470521", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
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
