# Ipv6DhcpOptiondefinition — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name              = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name              = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
        type = "nios.type"
      }
    }
  }

}
