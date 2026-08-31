# Auto-generated resource acceptance-test cases for RecordSvcb.
# TODO - Add Zone Auth as a PREREQ for the records
# As of 31st Aug , adding Zone Auth as PREREQ gives 500 Internal Server Error for Record SVCB
# Refer B1DDISPT-2207 for the same


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
      zone                = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { target_name = "{{random}}.com" }
      #zone               = infoblox_zone_auth.test.id
      zone                = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
      inheritance_sources = { ttl = { action = "override" } }
      ttl                 = 57600
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

  step {
    uddi {
      rdata               = { target_name = "{{random}}.com" }
      #zone               = infoblox_zone_auth.test.id
      zone                = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
      inheritance_sources = { ttl = { action = "inherit" } }
      ttl                 = 57600
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
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
      rdata = { target_name = "{{random}}.com", priority = 0 }
      zone  = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name" = "{{random}}.com"
      "uddi.rdata.priority"    = 0
    }
  }

  step {
    uddi {
      rdata = { target_name = "{{random}}_updated.com", priority = 2 }
      zone  = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name" = "{{random}}_updated.com"
      "uddi.rdata.priority"    = 2
    }
  }

}

case "rdata_svc_params" {
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
      rdata = {
        target_name = "{{random}}.com"
        priority    = 5
        svc_params  = [
          { key = "port",     value = "80" },
          { key = "ipv4hint", value = "1.1.1.1" },
          { key = "ipv6hint", value = "2001:db8::2" },
          { key = "ech",      value = "bWVvdw==" },
        ]
      }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name"        = "{{random}}.com"
      "uddi.rdata.priority"           = 5
      "uddi.rdata.svc_params.0.key"   = "port"
      "uddi.rdata.svc_params.0.value" = "80"
      "uddi.rdata.svc_params.1.key"   = "ipv4hint"
      "uddi.rdata.svc_params.1.value" = "1.1.1.1"
      "uddi.rdata.svc_params.2.key"   = "ipv6hint"
      "uddi.rdata.svc_params.2.value" = "2001:db8::2"
      "uddi.rdata.svc_params.3.key"   = "ech"
      "uddi.rdata.svc_params.3.value" = "bWVvdw=="
    }
  }

  step {
    uddi {
      rdata = {
        target_name = "{{random}}.com"
        priority    = 5
        svc_params  = [
          { key = "port",      value = "80" },
          { key = "ipv4hint",  value = "1.1.1.1" },
          { key = "ipv6hint",  value = "2001:db8::2" },
          { key = "ech",       value = "bWVvdw==" },
          { key = "mandatory", value = "ech,ipv6hint" },
        ]
      }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name"        = "{{random}}.com"
      "uddi.rdata.priority"           = 5
      "uddi.rdata.svc_params.0.key"   = "port"
      "uddi.rdata.svc_params.0.value" = "80"
      "uddi.rdata.svc_params.1.key"   = "ipv4hint"
      "uddi.rdata.svc_params.1.value" = "1.1.1.1"
      "uddi.rdata.svc_params.2.key"   = "ipv6hint"
      "uddi.rdata.svc_params.2.value" = "2001:db8::2"
      "uddi.rdata.svc_params.3.key"   = "ech"
      "uddi.rdata.svc_params.3.value" = "bWVvdw=="
      "uddi.rdata.svc_params.4.key"   = "mandatory"
      "uddi.rdata.svc_params.4.value" = "ech,ipv6hint"
    }
  }

  step {
    uddi {
      rdata = {
        target_name = "{{random}}.com"
        priority    = 5
        svc_params  = [
          { key = "alpn", value = "h3,h2,h9,h15" },
        ]
      }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name"        = "{{random}}.com"
      "uddi.rdata.priority"           = 5
      "uddi.rdata.svc_params.0.key"   = "alpn"
      "uddi.rdata.svc_params.0.value" = "h3,h2,h9,h15"
    }
  }

  step {
    uddi {
      rdata = {
        target_name = "{{random}}.com"
        priority    = 5
        svc_params  = [
          { key = "ohttp" },
          { key = "key13",     value = "69206c6f766520796f75" },
          { key = "key15",     value = "69206c6f766520796f7" },
          { key = "dohpath",   value = "/dns-query{?dns}" },
          { key = "mandatory", value = "dohpath,key13" },
        ]
      }
      zone = "dns/auth_zone/cf7a5e79-82c2-4de1-9788-4397c846d317"
    }
    check = {
      "uddi.rdata.target_name"        = "{{random}}.com"
      "uddi.rdata.priority"           = 5
      "uddi.rdata.svc_params.0.key"   = "ohttp"
      "uddi.rdata.svc_params.1.key"   = "key13"
      "uddi.rdata.svc_params.1.value" = "69206c6f766520796f75"
      "uddi.rdata.svc_params.2.key"   = "key15"
      "uddi.rdata.svc_params.2.value" = "69206c6f766520796f7"
      "uddi.rdata.svc_params.3.key"   = "dohpath"
      "uddi.rdata.svc_params.3.value" = "/dns-query{?dns}"
      "uddi.rdata.svc_params.4.key"   = "mandatory"
      "uddi.rdata.svc_params.4.value" = "dohpath,key13"
    }
  }

}
