# Auto-generated resource acceptance-test cases for RecordAaaa.
case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ipv6addr"           = "{{random_ipv6}}"
      "nios.name"               = "aaaa-record.{{random}}.com"
      "nios.view"               = "default"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
    }
  }

}

case "comment" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      comment  = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      comment  = "This is an updated record"
    }
    check = {
      "nios.comment" = "This is an updated record"
    }
  }

}

case "creator" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      creator  = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr       = "{{random_ipv6}}"
      view           = infoblox_zone_auth.test.nios.view
      creator        = "DYNAMIC"
      ddns_principal = "ddns_principal"
    }
    check = {
      "nios.ddns_principal" = "ddns_principal"
    }
  }

  step {
    nios {
      name           = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr       = "{{random_ipv6}}"
      view           = infoblox_zone_auth.test.nios.view
      creator        = "DYNAMIC"
      ddns_principal = "updated_ddns_principal"
    }
    check = {
      "nios.ddns_principal" = "updated_ddns_principal"
    }
  }

}

case "ddns_protected" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr       = "{{random_ipv6}}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

  step {
    nios {
      name           = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr       = "{{random_ipv6}}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

}

case "disable" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "forbid_reclamation" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name               = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr           = "{{random_ipv6}}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

  step {
    nios {
      name               = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr           = "{{random_ipv6}}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

}

case "ipv6addr" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      view     = infoblox_zone_auth.test.nios.view
      ipv6addr = "{{random_ipv6}}"
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
    }
  }

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      view     = infoblox_zone_auth.test.nios.view
      ipv6addr = "{{random_ipv62}}"
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv62}}"
    }
  }

}

case "func_call" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  # resource "infoblox_ipv6network" "test_func_call" {
  #   nios = {
  #     network = "{{random_ipv6_network}}"
  #     network_view = "default"
  #   }
  # }
  PREREQ

  step {
    nios {
      name               = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      # dynamic_allocation = { network = infoblox_ipv6network.test.nios.network, network_view = "default" }
      dynamic_allocation = { network = "2001:db8:abcd:12::/64", network_view = "default" }
      comment            = "Original Function Call"
    }
    # depends_on = [infoblox_ipv6network.test_func_call]
  }

  step {
    nios {
      name               = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      # dynamic_allocation = { network = infoblox_ipv6network.test.nios.network, network_view = "default" }
      dynamic_allocation = { network = "2001:db8:abcd:12::/64", network_view = "default" }
      comment            = "Updated Function Call"
    }
    # depends_on = [infoblox_ipv6network.test_func_call]
  }

}

case "name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv62}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
    }
  }

}

case "ttl" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
      ttl      = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}
