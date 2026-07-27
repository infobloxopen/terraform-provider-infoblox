# Auto-generated resource acceptance-test cases for RecordAaaa.
case "rdata" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      rdata = { address = "2001:db8::1" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "2001:db8::1"
    }
  }

  step {
    uddi {
      rdata = { address = "2001:db8::2" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "2001:db8::2"
    }
  }

}

case "options" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "rmz" {
    uddi = {
      fqdn = "1.0.0.2.ip6.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      rdata   = { address = "2001:db8::1" }
      options = { create_ptr = true, check_rmz = true }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "true"
    }
  }

  step {
    uddi {
      rdata   = { address = "2001:db8::1" }
      options = { create_ptr = true, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "false"
    }
  }

  step {
    uddi {
      rdata   = { address = "2001:db8::1" }
      options = { create_ptr = false, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "false"
      "uddi.options.check_rmz"  = "false"
    }
  }

}
