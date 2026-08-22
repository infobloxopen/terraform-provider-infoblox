# Auto-generated datasource acceptance-test cases for DhcpOptiondefinition.
case "filters" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
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
      name         = "test_option_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }

}
