# Ipv6filteroption — uddi datasource cases
case "filters" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
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

  pair_checks = ["uddi.comment", "uddi.lease_time", "uddi.name", "uddi.protocol", "uddi.role"]

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
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

  pair_checks = ["uddi.comment", "uddi.lease_time", "uddi.name", "uddi.protocol", "uddi.role"]

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random3}}" }
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}
