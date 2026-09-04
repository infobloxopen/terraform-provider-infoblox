# Auto-generated datasource acceptance-test cases for Ipv6rangetemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.cloud_api_compatible", "nios.comment", "nios.name", "nios.number_of_addresses", "nios.offset", "nios.recycle_leases", "nios.server_association_type", "nios.use_logic_filter_rules", "nios.use_recycle_leases"]

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 10
      offset               = 50
      cloud_api_compatible = true
    }
  }

}
