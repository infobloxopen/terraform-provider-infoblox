# Auto-generated resource acceptance-test cases for RecordCaa.
case "rdata" {
  backend = "uddi"
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
      rdata = { tag = "issue", value = "ca.example.com" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags" = "0"
      "uddi.rdata.tag"   = "issue"
      "uddi.rdata.value" = "ca.example.com"
    }
  }

  step {
    uddi {
      rdata = { tag = "issuewild", value = "*.example.com" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags" = "0"
      "uddi.rdata.tag"   = "issuewild"
      "uddi.rdata.value" = "*.example.com"
    }
  }

  step {
    uddi {
      rdata = { flags = 1, tag = "issuewild", value = "*.example.com" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags" = "1"
      "uddi.rdata.tag"   = "issuewild"
      "uddi.rdata.value" = "*.example.com"
    }
  }

  step {
    uddi {
      rdata = { flags = 0, tag = "issuewild", value = "*.example.com" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags" = "0"
      "uddi.rdata.tag"   = "issuewild"
      "uddi.rdata.value" = "*.example.com"
    }
  }

}
