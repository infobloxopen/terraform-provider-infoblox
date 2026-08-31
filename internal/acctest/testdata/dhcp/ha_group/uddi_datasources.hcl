# Auto-generated datasource acceptance-test cases for HaGroup.
#  TODO: Objects to be present in the grid for testing
#  dhcp/host/470520,
#  dhcp/host/470521
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.anycast_config_id", "uddi.comment", "uddi.ip_space", "uddi.mode", "uddi.name"]

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

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.anycast_config_id", "uddi.comment", "uddi.ip_space", "uddi.mode", "uddi.name"]

  step {
    uddi {
      hosts = [
        { host = "dhcp/host/470520", role = "active" },
        { host = "dhcp/host/470521", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
      tags = { tag1 = "{{random}}" }
    }
  }

}
