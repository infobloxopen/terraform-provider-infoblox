# Filteroption — uddi datasource cases
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "prereq_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

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
          option_code  = infoblox_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "prereq_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

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
          option_code  = infoblox_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}
