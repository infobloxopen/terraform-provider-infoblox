# Auto-generated datasource acceptance-test cases for Ipv6DhcpOptionspace.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      enterprise_number = "nios.enterprise_number"
      name              = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.enterprise_number", "nios.name"]

  step {
    nios {
      enterprise_number = 5674
      name              = "{{random}}"
    }
  }

}
