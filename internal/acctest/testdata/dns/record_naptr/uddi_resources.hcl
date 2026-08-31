# Auto-generated resource acceptance-test cases for RecordNaptr.
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
      name_in_zone = "naptr"
      rdata        = { order = 10, preference = 10, replacement = "+", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.order"       = "10"
      "uddi.rdata.preference"  = "10"
      "uddi.rdata.replacement" = "+"
      "uddi.rdata.services"    = "SIP+D2U"
    }
  }

  step {
    uddi {
      name_in_zone = "naptr"
      rdata        = { order = 20, preference = 20, replacement = ".", services = "SIP+E2U" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.order"       = "20"
      "uddi.rdata.preference"  = "20"
      "uddi.rdata.replacement" = "."
      "uddi.rdata.services"    = "SIP+E2U"
    }
  }

}

case "rdata_flags_and_regexp" {
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
      name_in_zone = "naptr"
      rdata        = { flags = "U", order = 100, preference = 10, regexp = "!^.*$!sip:jdoe@corpxyz.com!", replacement = ".", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags"  = "U"
      "uddi.rdata.regexp" = "!^.*$!sip:jdoe@corpxyz.com!"
    }
  }

  step {
    uddi {
      name_in_zone = "naptr"
      rdata        = { flags = "A", order = 100, preference = 10, regexp = "!^.*$!sip:jdoe@corpabc.com!", replacement = ".", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.flags"  = "A"
      "uddi.rdata.regexp" = "!^.*$!sip:jdoe@corpabc.com!"
    }
  }

}
