# Auto-generated datasource acceptance-test cases for Natgroup.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.name"]

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a test natgroup"
    }
  }

}
