# Auto-generated datasource acceptance-test cases for Ipv6DhcpOptiondefinition.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.array", "uddi.code", "uddi.comment", "uddi.name", "uddi.option_space", "uddi.type"]

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_{{random_int}}"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }

}

case "filters_code" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      code = "uddi.code"
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.array", "uddi.code", "uddi.comment", "uddi.name", "uddi.option_space", "uddi.type"]

  step {
    uddi {
      code         = 235
      name         = "basic_opt_code_{{random_int}}"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }

}
