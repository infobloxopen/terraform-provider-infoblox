# Auto-generated resource acceptance-test cases for RecordA.
# TODO - Add Zone Auth as a PREREQ for the records
# As of 31st Aug , adding Zone Auth as PREREQ gives 500 Internal Server Error for Record HTTPS
# Refer B1DDISPT-2207 for the same


case "inheritance_sources" {
  backend           = "uddi"
  parallel          = true
#   prerequisites_hcl = <<-PREREQ
#   resource "infoblox_zone_auth" "test" {
#     uddi = {
#       fqdn = "{{random}}.com."
#       primary_type = "cloud"
#     }
#   }
#   PREREQ

  step {
    uddi {
      rdata               = { target_name = "{{random}}.com" }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { target_name = "{{random}}.com" }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
      inheritance_sources = { ttl = { action = "override" } }
      ttl = 57600
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

  step {
      uddi {
        rdata               = { target_name = "{{random}}.com" }
        #zone                = infoblox_zone_auth.test.id
        zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
        inheritance_sources = { ttl = { action = "inherit" } }
        ttl = 57600
      }
      check = {
        "uddi.inheritance_sources.ttl.action" = "inherit"
      }
    }

}

case "rdata" {
  backend           = "uddi"
  parallel          = true
#   prerequisites_hcl = <<-PREREQ
#   resource "infoblox_zone_auth" "test" {
#     uddi = {
#       fqdn = "{{random}}.com."
#       primary_type = "cloud"
#     }
#   }
#   PREREQ

  step {
    uddi {
      rdata = { target_name = "{{random}}.com" , priority = 0 }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name" = "{{random}}.com"
      "uddi.rdata.priority" = 0
    }
  }

  step {
    uddi {
      rdata = { target_name = "{{random}}_updated.com" , priority = 2 }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name" = "{{random}}_updated.com"
      "uddi.rdata.priority" = 2
    }
  }

}
