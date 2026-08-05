# Auto-generated resource acceptance-test cases for RecordNs.
case "rdata" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      name_in_zone = "ns"
      rdata        = { dname = "ns1.example.com" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.dname" = "ns1.example.com"
    }
  }

  step {
    uddi {
      name_in_zone = "ns"
      rdata        = { dname = "ns2.example.com" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.dname" = "ns2.example.com"
    }
  }

}
