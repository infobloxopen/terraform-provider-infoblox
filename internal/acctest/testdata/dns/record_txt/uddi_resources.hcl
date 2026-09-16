# Auto-generated resource acceptance-test cases for RecordTxt.
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
      name_in_zone = "txt"
      rdata        = { text = "abc" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.text" = "abc"
    }
  }

  step {
    uddi {
      name_in_zone = "txt"
      rdata        = { text = "xyz" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.text" = "xyz"
    }
  }

}
