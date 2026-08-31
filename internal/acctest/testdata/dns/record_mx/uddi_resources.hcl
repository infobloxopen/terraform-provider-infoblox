# Auto-generated resource acceptance-test cases for RecordMx.
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
      name_in_zone = "mx"
      rdata        = { exchange = "m1.example.com", preference = 10 }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.exchange"   = "m1.example.com"
      "uddi.rdata.preference" = "10"
    }
  }

  step {
    uddi {
      name_in_zone = "mx"
      rdata        = { exchange = "m2.example.com", preference = 20 }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.exchange"   = "m2.example.com"
      "uddi.rdata.preference" = "20"
    }
  }

}
