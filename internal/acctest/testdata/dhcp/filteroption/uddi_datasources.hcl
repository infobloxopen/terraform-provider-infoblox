# Filteroption — uddi datasource cases
// An option code has to be created before running the test cases.

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
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }]
      }
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
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }]
      }
    }
  }

}
