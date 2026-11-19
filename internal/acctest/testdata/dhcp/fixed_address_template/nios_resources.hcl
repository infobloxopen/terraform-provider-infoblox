# Auto-generated resource acceptance-test cases for Fixedaddresstemplate.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                            = "{{random}}"
      "nios.deny_bootp"                      = "false"
      "nios.enable_ddns"                     = "false"
      "nios.enable_pxe_lease_time"           = "false"
      "nios.ext_attrs.%"                     = "0"
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "bootfile" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      bootfile = "bootfile-name"
    }
    check = {
      "nios.bootfile" = "bootfile-name"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      bootfile = "bootfile-name-updated"
    }
    check = {
      "nios.bootfile" = "bootfile-name-updated"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      bootserver = "10.0.0.0"
    }
    check = {
      "nios.bootserver" = "10.0.0.0"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      bootserver = "10.1.1.1"
    }
    check = {
      "nios.bootserver" = "10.1.1.1"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name            = "{{random}}"
      ddns_domainname = "example.com"
    }
    check = {
      "nios.ddns_domainname" = "example.com"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      ddns_domainname = "example.org"
    }
    check = {
      "nios.ddns_domainname" = "example.org"
    }
  }

}

case "ddns_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      ddns_hostname = "ddns_host.name"
    }
    check = {
      "nios.ddns_hostname" = "ddns_host.name"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      ddns_hostname = "updated_ddns_host.name"
    }
    check = {
      "nios.ddns_hostname" = "updated_ddns_host.name"
    }
  }

}

case "deny_bootp" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      deny_bootp = true
    }
    check = {
      "nios.deny_bootp" = "true"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      deny_bootp = false
    }
    check = {
      "nios.deny_bootp" = "false"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      enable_ddns = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      enable_ddns = false
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "enable_pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      enable_pxe_lease_time = true
      pxe_lease_time        = 100
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      enable_pxe_lease_time = false
      pxe_lease_time        = 200
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                            = "{{random}}"
      ignore_dhcp_option_list_request = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      name                            = "{{random}}"
      ignore_dhcp_option_list_request = false
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_filteroption" "filter1" {
    nios = {
      name = "{{random}}-1"
    }
  }
  resource "infoblox_filteroption" "filter2" {
    nios = {
      name = "{{random}}-2"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random}}"
      logic_filter_rules = [{ filter = infoblox_filteroption.filter1.nios.name, type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "{{random}}-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      logic_filter_rules = [{ filter = infoblox_filteroption.filter2.nios.name, type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "{{random}}-2"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "nextserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      nextserver = "10.0.0.0"
    }
    check = {
      "nios.nextserver" = "10.0.0.0"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      nextserver = "10.1.1.1"
    }
    check = {
      "nios.nextserver" = "10.1.1.1"
    }
  }

}

case "number_of_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      number_of_addresses = 3
      offset              = 2
    }
    check = {
      "nios.number_of_addresses" = "3"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      number_of_addresses = 5
      offset              = 2
    }
    check = {
      "nios.number_of_addresses" = "5"
    }
  }

}

case "offset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      offset              = 3
      number_of_addresses = 10
    }
    check = {
      "nios.offset" = "3"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      offset              = 2
      number_of_addresses = 10
    }
    check = {
      "nios.offset" = "2"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      options = [{ name = "time-offset", num = 2, value = "50" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.name"  = "time-offset"
      "nios.options.0.num"   = "2"
      "nios.options.0.value" = "50"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      options = [{ num = 51, value = "7200" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.num"   = "51"
      "nios.options.0.value" = "7200"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

}

case "pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      pxe_lease_time = 10
    }
    check = {
      "nios.pxe_lease_time" = "10"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      pxe_lease_time = 20
    }
    check = {
      "nios.pxe_lease_time" = "20"
    }
  }

}

case "import_basic" {
  backend       = "nios"
  parallel      = true
  import        = true
  import_ignore = ["nios.ext_attrs", "nios.ext_attrs_all", "nios.options"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "import_all_fields" {
  backend       = "nios"
  parallel      = true
  import        = true
  import_ignore = ["nios.ext_attrs", "nios.ext_attrs_all", "nios.options"]
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_filteroption" "filter1" {
    nios = {
      name = "{{random}}-1"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      comment             = "imported template"
      number_of_addresses = 3
      offset              = 2
      bootfile            = "bootfile-name"
      bootserver          = "10.0.0.1"
      nextserver          = "10.0.0.2"
      ddns_hostname       = "ddns_host.name"
      ddns_domainname     = "example.com"
      deny_bootp          = true
      enable_ddns         = true
      pxe_lease_time      = 10
      options             = [{ name = "time-offset", num = 2, value = "50" }, { name = "subnet-mask", value = "1.1.1.1" }]
      logic_filter_rules  = [{ filter = infoblox_filteroption.filter1.nios.name, type = "Option" }]
      ext_attrs           = { Site = "{{random}}" }
    }
  }

}
