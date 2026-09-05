case "filters" {
  backend = "nios"
  filter {
    type = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.name", "nios.delegate_to"]

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  filter {
    type = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.name", "nios.delegate_to"]

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
    ext_attrs = { Site = "{{random2}}" }
  }
}
