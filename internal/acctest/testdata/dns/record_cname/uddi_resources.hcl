# Auto-generated resource acceptance-test cases for RecordCname.
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
      name_in_zone = "cname"
      rdata        = { cname = "c1.${infoblox_zone_auth.test.uddi.fqdn}" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.cname" = "c1.{{random}}.com."
    }
  }

  step {
    uddi {
      name_in_zone = "cname"
      rdata        = { cname = "c2.${infoblox_zone_auth.test.uddi.fqdn}" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.cname" = "c2.{{random}}.com."
    }
  }

}
