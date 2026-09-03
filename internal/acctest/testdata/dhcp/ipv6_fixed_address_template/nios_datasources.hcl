# Auto-generated datasource acceptance-test cases for Ipv6fixedaddresstemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.domain_name", "nios.name", "nios.number_of_addresses", "nios.offset", "nios.preferred_lifetime", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.domain_name", "nios.name", "nios.number_of_addresses", "nios.offset", "nios.preferred_lifetime", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
