# Auto-generated datasource acceptance-test cases for Ipv6DhcpOptiondefinition.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      code = "nios.code"
      name = "nios.name"
      type = "nios.type"
    }
  }

  pair_checks = ["nios.code", "nios.name", "nios.space", "nios.type"]

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
  }

}
