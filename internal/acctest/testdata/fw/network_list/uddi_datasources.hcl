# Auto-generated datasource acceptance-test cases for NetworkList.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.description", "uddi.name"]

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
  }

}
