# Auto-generated resource acceptance-test cases for RecordNs.
case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.name"                        = "{{random}}.com"
      "nios.nameserver"                  = "{{random2}}.{{random}}.com"
      "nios.addresses.0.address"         = "{{random_ip}}"
      "nios.addresses.0.auto_create_ptr" = "false"
      "nios.view"                        = "default"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
  }

}

case "addresses" {
  backend  = "nios"
  parallel = true

  # auto_create_ptr needs the reverse-mapping zone of the glue address to exist.
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  resource "infoblox_zone_auth" "rmz" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.addresses.0.address"         = "{{random_ip}}"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip2}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.addresses.0.address"         = "{{random_ip2}}"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_octet}}.0.0.1", auto_create_ptr = true }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone, infoblox_zone_auth.rmz]
    check = {
      "nios.addresses.0.address"         = "{{random_octet}}.0.0.1"
      "nios.addresses.0.auto_create_ptr" = "true"
    }
  }

}

case "nameserver" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.nameserver" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random3}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.nameserver" = "{{random3}}.{{random}}.com"
    }
  }

}
