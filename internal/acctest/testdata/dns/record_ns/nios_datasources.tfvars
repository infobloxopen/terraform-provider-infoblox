# Auto-generated datasource acceptance-test cases for RecordNs.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name       = "nios.name"
      nameserver = "nios.nameserver"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL"
      view       = "default"
    }
  }

}
