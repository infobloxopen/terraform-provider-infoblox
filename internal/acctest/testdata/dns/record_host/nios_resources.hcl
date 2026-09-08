# Auto-generated resource acceptance-test cases for RecordHost.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      ipv4addrs = [{ ipv4addr = "192.168.1.10" }]
      view      = "default"
    }
    check = {
      "nios.name"                 = "{{random}}.example.com"
      "nios.view"                 = "default"
      "nios.ipv4addrs.0.ipv4addr" = "192.168.1.10"
      "nios.configure_for_dns"    = "true"
      "nios.ddns_protected"       = "false"
      "nios.disable"              = "false"
      "nios.disable_discovery"    = "false"
      "nios.network_view"         = "default"
      "nios.rrset_order"          = "cyclic"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name      = "{{random}}.example.com"
      ipv4addrs = [{ ipv4addr = "192.168.1.11" }]
      view      = "default"
    }
  }

}

case "aliases" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      aliases   = ["{{random2}}.example.com"]
      ipv4addrs = [{ ipv4addr = "192.168.1.12" }]
    }
    check = {
      "nios.aliases.#" = "1"
      "nios.aliases.0" = "{{random2}}.example.com"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      aliases   = ["{{random3}}.example.com"]
      ipv4addrs = [{ ipv4addr = "192.168.1.12" }]
    }
    check = {
      "nios.aliases.#" = "1"
      "nios.aliases.0" = "{{random3}}.example.com"
    }
  }

}

case "allow_telnet" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping the test as backend isn't setting the values correctly"
  parallel    = true

  step {
    nios {
      name         = "{{random}}.example.com"
      view         = "default"
      ipv4addrs    = [{ ipv4addr = "192.168.1.13" }]
      allow_telnet = false
    }
    check = {
      "nios.allow_telnet" = "false"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "cli_credentials" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#" = "0"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#"                 = "1"
      "nios.cli_credentials.0.user"            = "user1"
      "nios.cli_credentials.0.credential_type" = "SSH"
      "nios.cli_credentials.0.comment"         = "cli credential comment"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#"                 = "1"
      "nios.cli_credentials.0.user"            = "user1"
      "nios.cli_credentials.0.credential_type" = "SSH"
      "nios.cli_credentials.0.comment"         = "cli credential comment"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#"                 = "1"
      "nios.cli_credentials.0.user"            = "user2"
      "nios.cli_credentials.0.credential_type" = "SSH"
      "nios.cli_credentials.0.comment"         = "cli credential comment update"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#"                 = "1"
      "nios.cli_credentials.0.user"            = "user1"
      "nios.cli_credentials.0.credential_type" = "SSH"
      "nios.cli_credentials.0.comment"         = "cli credential comment"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#" = "0"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.cli_credentials.#"                 = "1"
      "nios.cli_credentials.0.user"            = "user2"
      "nios.cli_credentials.0.credential_type" = "SSH"
      "nios.cli_credentials.0.comment"         = "cli credential comment update"
    }
  }

}

case "cloud_info" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.14" }]
    }
    check = {
      "nios.cloud_info.authority_type"   = "GM"
      "nios.cloud_info.delegated_scope"  = "NONE"
      "nios.cloud_info.mgmt_platform"    = ""
      "nios.cloud_info.owned_by_adaptor" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.15" }]
      comment   = "new host record"
    }
    check = {
      "nios.comment" = "new host record"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.15" }]
      comment   = "updated host record"
    }
    check = {
      "nios.comment" = "updated host record"
    }
  }

}

case "configure_for_dns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name              = "{{random}}.example.com"
      view              = "default"
      ipv4addrs         = [{ ipv4addr = "10.0.0.249" }]
      configure_for_dns = true
    }
    check = {
      "nios.configure_for_dns" = "true"
    }
  }

  step {
    nios {
      name              = "{{random}}.example.com"
      view              = "default"
      ipv4addrs         = [{ ipv4addr = "10.0.0.249" }]
      configure_for_dns = false
    }
    check = {
      "nios.configure_for_dns" = "false"
    }
  }

}

case "ddns_protected" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}.example.com"
      view           = "default"
      ipv4addrs      = [{ ipv4addr = "192.168.1.17" }]
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random}}.example.com"
      view           = "default"
      ipv4addrs      = [{ ipv4addr = "192.168.1.17" }]
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "device_description" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}.example.com"
      view               = "default"
      ipv4addrs          = [{ ipv4addr = "192.168.1.18" }]
      device_description = "device description"
    }
    check = {
      "nios.device_description" = "device description"
    }
  }

  step {
    nios {
      name               = "{{random}}.example.com"
      view               = "default"
      ipv4addrs          = [{ ipv4addr = "192.168.1.18" }]
      device_description = "updated device description"
    }
    check = {
      "nios.device_description" = "updated device description"
    }
  }

}

case "device_location" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name            = "{{random}}.example.com"
      view            = "default"
      ipv4addrs       = [{ ipv4addr = "192.168.1.19" }]
      device_location = "device location"
    }
    check = {
      "nios.device_location" = "device location"
    }
  }

  step {
    nios {
      name            = "{{random}}.example.com"
      view            = "default"
      ipv4addrs       = [{ ipv4addr = "192.168.1.19" }]
      device_location = "updated device location"
    }
    check = {
      "nios.device_location" = "updated device location"
    }
  }

}

case "device_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      view        = "default"
      ipv4addrs   = [{ ipv4addr = "192.168.1.20" }]
      device_type = "device type"
    }
    check = {
      "nios.device_type" = "device type"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      view        = "default"
      ipv4addrs   = [{ ipv4addr = "192.168.1.20" }]
      device_type = "updated device type"
    }
    check = {
      "nios.device_type" = "updated device type"
    }
  }

}

case "device_vendor" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}.example.com"
      view          = "default"
      ipv4addrs     = [{ ipv4addr = "192.168.1.21" }]
      device_vendor = "device vendor"
    }
    check = {
      "nios.device_vendor" = "device vendor"
    }
  }

  step {
    nios {
      name          = "{{random}}.example.com"
      view          = "default"
      ipv4addrs     = [{ ipv4addr = "192.168.1.21" }]
      device_vendor = "updated device vendor"
    }
    check = {
      "nios.device_vendor" = "updated device vendor"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.22" }]
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.22" }]
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "disable_discovery" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name              = "{{random}}.example.com"
      view              = "default"
      ipv4addrs         = [{ ipv4addr = "192.168.1.23" }]
      disable_discovery = true
    }
    check = {
      "nios.disable_discovery" = "true"
    }
  }

  step {
    nios {
      name              = "{{random}}.example.com"
      view              = "default"
      ipv4addrs         = [{ ipv4addr = "192.168.1.23" }]
      disable_discovery = false
    }
    check = {
      "nios.disable_discovery" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.26" }]
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.26" }]
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv4addrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.27" }]
    }
    check = {
      "nios.ipv4addrs.#"          = "1"
      "nios.ipv4addrs.0.ipv4addr" = "192.168.1.27"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.1.28" }]
    }
    check = {
      "nios.ipv4addrs.#"          = "1"
      "nios.ipv4addrs.0.ipv4addr" = "192.168.1.28"
    }
  }

}

case "ipv6addrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv6addrs = [{ ipv6addr = "fd00:1234:5678::1" }]
    }
    check = {
      "nios.ipv6addrs.#"          = "1"
      "nios.ipv6addrs.0.ipv6addr" = "fd00:1234:5678::1"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv6addrs = [{ ipv6addr = "fd00:1234:5678::12" }]
    }
    check = {
      "nios.ipv6addrs.#"          = "1"
      "nios.ipv6addrs.0.ipv6addr" = "fd00:1234:5678::12"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.10" }]
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name      = "{{random2}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.10" }]
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.11" }]
    }
    check = {
      "nios.network_view" = "default"
    }
  }

}

case "rrset_order" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      view        = "default"
      ipv4addrs   = [{ ipv4addr = "192.168.2.12" }]
      rrset_order = "cyclic"
    }
    check = {
      "nios.rrset_order" = "cyclic"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      view        = "default"
      ipv4addrs   = [{ ipv4addr = "192.168.2.12" }]
      rrset_order = "random"
    }
    check = {
      "nios.rrset_order" = "random"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "snmp3_credential" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random2}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random2}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random3}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random3}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random4}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
  }

  step {
    nios {
      name                 = "{{random4}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random5}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random5}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
  }

  step {
    nios {
      name                 = "{{random6}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user1"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

  step {
    nios {
      name                 = "{{random6}}.example.com"
      ipv4addrs            = [{ ipv4addr = "192.168.1.10" }]
    }
    check = {
      "nios.snmp3_credential.user"                    = "user2"
      "nios.snmp3_credential.authentication_protocol" = "SHA"
      "nios.snmp3_credential.privacy_protocol"        = "AES"
    }
  }

}

case "snmp_credential" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}.example.com"
      view                = "default"
      ipv4addrs           = [{ ipv4addr = "192.168.2.30" }]
      snmp_credential     = { community_string = "COMMUNITY_STRING", comment = "SNMP Credential Comment", credential_group = "default" }
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      view                = "default"
      ipv4addrs           = [{ ipv4addr = "192.168.2.30" }]
      snmp_credential     = { community_string = "COMMUNITY_STRING_UPDATED", comment = "SNMP Credential Comment Updated", credential_group = "default" }
    }
    check = {
      "nios.snmp_credential.community_string" = "COMMUNITY_STRING_UPDATED"
      "nios.snmp_credential.comment"          = "SNMP Credential Comment Updated"
      "nios.snmp_credential.credential_group" = "default"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.13" }]
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.13" }]
      ttl       = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "use_dns_ea_inheritance" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}.example.com"
      view                   = "default"
      ipv4addrs              = [{ ipv4addr = "192.168.2.14" }]
      use_dns_ea_inheritance = true
    }
    check = {
      "nios.use_dns_ea_inheritance" = "true"
    }
  }

  step {
    nios {
      name                   = "{{random}}.example.com"
      view                   = "default"
      ipv4addrs              = [{ ipv4addr = "192.168.2.14" }]
      use_dns_ea_inheritance = false
    }
    check = {
      "nios.use_dns_ea_inheritance" = "false"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.16" }]
    }
    check = {
      "nios.view" = "default"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ipv4addrs = [{ ipv4addr = "192.168.2.16" }]
    }
    check = {
      "nios.view" = "default"
    }
  }

}
