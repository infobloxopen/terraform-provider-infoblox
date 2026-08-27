# Auto-generated datasource acceptance-test cases for Ruleset.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disabled", "nios.name", "nios.type"]

  step {
    nios {
      name = "{{random}}"
      type = "NXDOMAIN"
    }
  }

}
