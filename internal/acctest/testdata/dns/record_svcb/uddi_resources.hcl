# Auto-generated resource acceptance-test cases for RecordSvcb.
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
      name_in_zone = "svcb"
      rdata        = { target_name = "google.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target_name" = "google.com."
    }
  }

  step {
    uddi {
      name_in_zone = "svcb"
      rdata        = { target_name = "apple.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target_name" = "apple.com."
    }
  }

}
