# Auto-generated resource acceptance-test cases for RecordNs.
case "basic" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "example.com"
      view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.name"                        = "example.com"
      "nios.nameserver"                  = "{{random}}.example.com"
      "nios.addresses.0.address"         = "20.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
      "nios.view"                        = "default"
      //"nios.ms_delegation_name"          = ""
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "example.com"
      view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
  }

}

case "addresses" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "example.com"
      view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.addresses.0.address"         = "20.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{address = "40.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.addresses.0.address"         = "40.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "ns1.example.com"
      addresses  = [{address = "40.0.0.0", auto_create_ptr = true}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.addresses.0.address"         = "40.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "true"
    }
  }

}

case "nameserver" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "example.com"
      view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.nameserver" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random2}}.example.com"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
    check = {
      "nios.nameserver" = "{{random2}}.example.com"
    }
  }

}
