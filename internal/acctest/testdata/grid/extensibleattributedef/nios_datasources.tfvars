# Auto-generated datasource acceptance-test cases for Extensibleattributedef.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.default_value", "nios.flags", "nios.max", "nios.min", "nios.name", "nios.type"]

  step {
    nios {
      name = "{{random}}"
      type = "STRING"
    }
  }

}
