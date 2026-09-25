# HardwareFilter — uddi datasource cases
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.lease_time", "uddi.name", "uddi.role"]

  step {
    uddi {
      name = "{{random}}"
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

  pair_checks = ["uddi.comment", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.lease_time", "uddi.name", "uddi.role"]

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random2}}" }
    }
  }

}
