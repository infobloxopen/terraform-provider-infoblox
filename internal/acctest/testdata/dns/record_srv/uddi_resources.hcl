# Auto-generated resource acceptance-test cases for RecordSrv.
case "rdata" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      name_in_zone = "srv"
      rdata        = { port = 80, priority = 10, target = "abc.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.port"     = "80"
      "uddi.rdata.priority" = "10"
      "uddi.rdata.target"   = "abc.com."
    }
  }

  step {
    uddi {
      name_in_zone = "srv"
      rdata        = { port = 90, priority = 20, target = "xyz.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.port"     = "90"
      "uddi.rdata.priority" = "20"
      "uddi.rdata.target"   = "xyz.com."
    }
  }

  step {
    uddi {
      name_in_zone = "srv"
      rdata        = { port = 90, priority = 20, target = "xyz.com.", weight = 10 }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.port"     = "90"
      "uddi.rdata.priority" = "20"
      "uddi.rdata.target"   = "xyz.com."
      "uddi.rdata.weight"   = "10"
    }
  }

}
