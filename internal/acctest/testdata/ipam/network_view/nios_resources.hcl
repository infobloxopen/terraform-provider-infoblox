# Auto-generated resource acceptance-test cases for Networkview.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"        = "{{random}}"
      "nios.mgm_private" = "false"
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
      name = "{{random}}"
    }
  }

}

case "cloud_info" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Requires Cloud API Configurations to Run this test"
  parallel    = true

  step {
    nios {
      name       = "{{random}}"
      cloud_info = { delegated_member = { name = "{{grid_member_hostname}}" } }
    }
    check = {
      "nios.name"                             = "{{random}}"
      "nios.cloud_info.authority_type"        = "Member"
      "nios.cloud_info.delegated_scope"       = "ROOT"
      "nios.cloud_info.owned_by_adaptor"      = "false"
      "nios.cloud_info.delegated_member.name" = "{{grid_member_hostname}}"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      cloud_info = { delegated_member = null }
    }
    check = {
      "nios.cloud_info.#" = "0"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a new network view"
    }
    check = {
      "nios.comment" = "This is a new network view"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a modified network view"
    }
    check = {
      "nios.comment" = "This is a modified network view"
    }
  }

}

case "ddns_dns_view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
    }
    check = {
      "nios.ddns_dns_view" = "default.{{random}}"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      ddns_dns_view = "default.{{random}}"
    }
    check = {
      "nios.ddns_dns_view" = "default.{{random}}"
    }
  }

}

case "ddns_zone_primaries" {
  backend     = "nios"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "parent_zone1" {
    nios = {
      fqdn = "{{random2}}.com"
      grid_primary = [{ name = "{{grid_member_hostname}}" }]
      view = "default"
    }
  }
  resource "infoblox_zone_auth" "parent_zone2" {
    nios = {
      fqdn = "{{random3}}.com"
      grid_primary = [{ name = "{{grid_member_hostname}}" }]
      view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ddns_zone_primaries = [{ dns_grid_primary = "{{grid_member_hostname}}", zone_match = "GRID", dns_grid_zone = { ref = "${infoblox_zone_auth.parent_zone1.id}" } }]
    }
    check = {
      "nios.ddns_zone_primaries.0.dns_grid_primary" = "{{grid_member_hostname}}"
      "nios.ddns_zone_primaries.0.zone_match"       = "GRID"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ddns_zone_primaries = [{ dns_grid_primary = "{{grid_member_hostname}}", zone_match = "GRID", dns_grid_zone = { ref = "${infoblox_zone_auth.parent_zone2.id}" } }]
    }
    check = {
      "nios.ddns_zone_primaries.0.dns_grid_primary" = "{{grid_member_hostname}}"
      "nios.ddns_zone_primaries.0.zone_match"       = "GRID"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "federated_realms" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Requires UDDI Configurations to Run this test"
  parallel    = true

  step {
    nios {
      name             = "{{random}}"
      federated_realms = [{ id = 22, name = "federated_realm_1" }]
    }
    check = {
      "nios.federated_realms.0.id"   = "11"
      "nios.federated_realms.0.name" = "federated_realm_1"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      federated_realms = [{ id = 22, name = "federated_realm_2" }]
    }
    check = {
      "nios.federated_realms.0.id"   = "22"
      "nios.federated_realms.0.name" = "federated_realm_2"
    }
  }

}

case "internal_forward_zones" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone1" {
    nios = {
      fqdn = "{{random}}1"
      view = "default.{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test_zone2" {
    nios = {
      fqdn = "{{random}}2"
      view = "default.{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      internal_forward_zones = ["${infoblox_zone_auth.test_zone1.id}"]
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      internal_forward_zones = ["${infoblox_zone_auth.test_zone1.id}", "${infoblox_zone_auth.test_zone2.id}"]
    }
  }

}

case "mgm_private" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      mgm_private = false
    }
    check = {
      "nios.mgm_private" = "false"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      mgm_private = true
    }
    check = {
      "nios.mgm_private" = "true"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "remote_forward_zones" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      remote_forward_zones = [{ fqdn = "fwdzone1.com", key_type = "TSIG", server_address = "192.168.12.12", tsig_key_name = "tsigkey", tsig_key_alg = "HMAC-SHA256", tsig_key = "dGhpc2lzdGVzdHRzaWdrZXk=" }]
    }
    check = {
      "nios.remote_forward_zones.0.fqdn"           = "fwdzone1.com"
      "nios.remote_forward_zones.0.key_type"       = "TSIG"
      "nios.remote_forward_zones.0.server_address" = "192.168.12.12"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      remote_forward_zones = [{ fqdn = "fwdzone2.com", key_type = "TSIG", server_address = "192.168.12.13", tsig_key_name = "tsigkey2", tsig_key_alg = "HMAC-SHA256", tsig_key = "dGhpc2lzdGBzdHRzaWdrZXk=" }]
    }
    check = {
      "nios.remote_forward_zones.0.fqdn"           = "fwdzone2.com"
      "nios.remote_forward_zones.0.server_address" = "192.168.12.13"
      "nios.remote_forward_zones.0.key_type"       = "TSIG"
      "nios.remote_forward_zones.0.tsig_key_name"  = "tsigkey2"
      "nios.remote_forward_zones.0.tsig_key_alg"   = "HMAC-SHA256"
      "nios.remote_forward_zones.0.tsig_key"       = "dGhpc2lzdGBzdHRzaWdrZXk="
    }
  }

}

case "remote_reverse_zones" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      remote_reverse_zones = [{ fqdn = "0.168.192.in-addr.arpa", key_type = "NONE", server_address = "192.168.12.12" }]
    }
    check = {
      "nios.remote_reverse_zones.0.fqdn"           = "0.168.192.in-addr.arpa"
      "nios.remote_reverse_zones.0.key_type"       = "NONE"
      "nios.remote_reverse_zones.0.server_address" = "192.168.12.12"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      remote_reverse_zones = [{ fqdn = "2.168.192.in-addr.arpa", key_type = "NONE", server_address = "192.168.12.13" }]
    }
    check = {
      "nios.remote_reverse_zones.0.fqdn"           = "2.168.192.in-addr.arpa"
      "nios.remote_reverse_zones.0.key_type"       = "NONE"
      "nios.remote_reverse_zones.0.server_address" = "192.168.12.13"
    }
  }

}
