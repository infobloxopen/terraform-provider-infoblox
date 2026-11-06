# Auto-generated datasource acceptance-test cases for Fixedaddresstemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_hostname", "nios.deny_bootp", "nios.enable_ddns", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.name", "nios.nextserver", "nios.number_of_addresses", "nios.offset", "nios.pxe_lease_time", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_dhcp_option_list_request", "nios.use_logic_filter_rules", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time"]

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

  pair_checks = ["nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_hostname", "nios.deny_bootp", "nios.enable_ddns", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.name", "nios.nextserver", "nios.number_of_addresses", "nios.offset", "nios.pxe_lease_time", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_dhcp_option_list_request", "nios.use_logic_filter_rules", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
