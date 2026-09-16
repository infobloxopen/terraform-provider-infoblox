# Auto-generated resource acceptance-test cases for RecordDname.
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
      name_in_zone = "dname"
      rdata        = { target = "google.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target" = "google.com."
    }
  }

  step {
    uddi {
      name_in_zone = "dname"
      rdata        = { target = "apple.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target" = "apple.com."
    }
  }

}
