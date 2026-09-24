# Auto-generated datasource acceptance-test cases for HaGroup.
#  TODO: Objects to be present in the grid for testing
#  DHCP Hosts
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.anycast_config_id", "uddi.comment", "uddi.ip_space", "uddi.mode", "uddi.name"]

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.anycast_config_id", "uddi.comment", "uddi.ip_space", "uddi.mode", "uddi.name"]

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
      tags = { tag1 = "{{random}}" }
    }
  }

}
