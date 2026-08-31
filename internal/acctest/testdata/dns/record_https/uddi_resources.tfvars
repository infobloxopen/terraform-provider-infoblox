# Auto-generated resource acceptance-test cases for RecordA.
case "inheritance_sources" {
  backend           = "uddi"
  parallel          = true
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
      rdata               = { target_name = "{{random}}.com" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { target_name = "{{random}}.com" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "rdata" {
  backend           = "uddi"
  parallel          = true
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
      rdata = { target_name = "{{random}}.com" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target_name" = "{{random}}.com"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip2}}" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.target_name" = "{{random_ip2}}"
    }
  }

}
