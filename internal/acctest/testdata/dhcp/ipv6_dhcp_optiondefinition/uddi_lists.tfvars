# Ipv6DhcpOptiondefinition — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_{{random_int}}"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_{{random_int}}"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
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
        type = "uddi.type"
      }
    }
  }

}
