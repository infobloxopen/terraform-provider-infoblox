# Auto-generated resource acceptance-test cases for RecordA.
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ipv4addr"           = "{{random_ip}}"
      "nios.name"               = "{{random2}}.{{random}}.com"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      comment  = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr       = "{{random_ip}}"
      view           = infoblox_zone_auth.test.nios.view
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_1"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_1"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr       = "{{random_ip}}"
      view           = infoblox_zone_auth.test.nios.view
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_2"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_2"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr       = "{{random_ip}}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr       = "{{random_ip}}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "extattrs" {
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
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr  = "{{random_ip}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr  = "{{random_ip}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
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
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr           = "{{random_ip}}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr           = "{{random_ip}}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
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
  resource "infoblox_network" "test_func_call" {
    nios = {
      network = "{{random_cidr_network}}"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      dynamic_allocation = { network = infoblox_network.test_func_call.nios.network, network_view = "default" }
      comment            = "Original Function Call"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      dynamic_allocation = { network = infoblox_network.test_func_call.nios.network, network_view = "default" }
      comment            = "Function Call with Update"
    }
  }

}

case "ipv4addr" {
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
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ipv4addr" = "{{random_ip}}"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip2}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ipv4addr" = "{{random_ip2}}"
    }
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
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ipv4addr = "{{random_ip}}"
      view     = infoblox_zone_auth.test.nios.view
      ttl      = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}
