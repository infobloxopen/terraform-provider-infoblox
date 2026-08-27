# Auto-generated datasource acceptance-test cases for DhcpOptiondefinition.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name  = "nios.name"
      code  = "nios.code"
      space = "nios.space"
      type  = "nios.type"
    }
  }

  pair_checks = ["nios.code", "nios.name", "nios.space", "nios.type"]

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
  }

}
