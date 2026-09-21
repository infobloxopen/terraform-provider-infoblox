# Auto-generated datasource acceptance-test cases for Bulkhostnametemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      template_name   = "nios.template_name"
      template_format = "nios.template_format"
    }
  }

  pair_checks = ["nios.template_format", "nios.template_name"]

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "host-$4"
    }
  }

}
