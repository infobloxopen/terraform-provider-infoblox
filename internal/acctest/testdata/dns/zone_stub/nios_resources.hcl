# Auto-generated resource acceptance-test cases for ZoneStub.
// Objects to be present on the GRID for the tests to run 
// NS Group Forward Stub Server - nsgroup_forwardstubserver_1 , nsgroup_forwardstubserver_2
// NS Group Stub Member - ns_group_stub_member_1 , ns_group_stub_member_2

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
    }
    check = {
      "nios.fqdn"                = "{{random}}.com"
      "nios.stub_from.0.address" = "1.1.1.1"
      "nios.stub_from.0.name"    = "{{random2}}"
      "nios.disable"             = "false"
      "nios.disable_forwarding"  = "false"
      "nios.locked"              = "false"
      "nios.ms_ad_integrated"    = "false"
      "nios.ms_ddns_mode"        = "NONE"
      "nios.view"                = "default"
      "nios.zone_format"         = "FORWARD"
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
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      comment   = "Example Comment"
    }
    check = {
      "nios.comment" = "Example Comment"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      comment   = "Updated Comment"
    }
    check = {
      "nios.comment" = "Updated Comment"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "disable_forwarding" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn               = "{{random}}.com"
      stub_from          = [{ address = "1.1.1.1", name = "{{random2}}" }]
      disable_forwarding = false
    }
    check = {
      "nios.disable_forwarding" = "false"
    }
  }

  step {
    nios {
      fqdn               = "{{random}}.com"
      stub_from          = [{ address = "1.1.1.1", name = "{{random2}}" }]
      disable_forwarding = true
    }
    check = {
      "nios.disable_forwarding" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "external_ns_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn              = "{{random}}.com"
      stub_from         = [{ address = "1.1.1.1", name = "{{random2}}" }]
      external_ns_group = "nsgroup_forwardstubserver_1"
    }
    check = {
      "nios.external_ns_group" = "nsgroup_forwardstubserver_1"
    }
  }

  step {
    nios {
      fqdn              = "{{random}}.com"
      stub_from         = [{ address = "1.1.1.1", name = "{{random2}}" }]
      external_ns_group = "nsgroup_forwardstubserver_2"
    }
    check = {
      "nios.external_ns_group" = "nsgroup_forwardstubserver_2"
    }
  }

}

case "locked" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      locked    = true
    }
    check = {
      "nios.locked" = "true"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      locked    = false
    }
    check = {
      "nios.locked" = "false"
    }
  }

}

case "ms_ad_integrated" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      stub_from        = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ms_ad_integrated = true
    }
    check = {
      "nios.ms_ad_integrated" = "true"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      stub_from        = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ms_ad_integrated = false
    }
    check = {
      "nios.ms_ad_integrated" = "false"
    }
  }

}

case "ms_ddns_mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      stub_from    = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ms_ddns_mode = "ANY"
    }
    check = {
      "nios.ms_ddns_mode" = "ANY"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      stub_from    = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ms_ddns_mode = "SECURE"
    }
    check = {
      "nios.ms_ddns_mode" = "SECURE"
    }
  }

}

case "ns_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ns_group  = "ns_group_stub_member_1"
    }
    check = {
      "nios.ns_group" = "ns_group_stub_member_1"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ns_group  = "ns_group_stub_member_2"
    }
    check = {
      "nios.ns_group" = "ns_group_stub_member_2"
    }
  }

}

case "prefix" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      prefix    = "STUB-b"
    }
    check = {
      "nios.prefix" = "STUB-b"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      prefix    = "{{random3}}"
    }
    check = {
      "nios.prefix" = "{{random3}}"
    }
  }

}

case "stub_from" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
    }
    check = {
      "nios.stub_from.0.name"    = "{{random2}}"
      "nios.stub_from.0.address" = "1.1.1.1"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "2.2.2.2", name = "{{random3}}" }]
    }
    check = {
      "nios.stub_from.0.name"    = "{{random3}}"
      "nios.stub_from.0.address" = "2.2.2.2"
    }
  }

}

case "stub_members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      stub_from    = [{ address = "1.1.1.1", name = "{{random2}}" }]
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.stub_members.0.name" = "{{grid_member_hostname}}"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      stub_from    = [{ address = "1.1.1.1", name = "{{random2}}" }]
      stub_members = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.stub_members.0.name" = "{{grid_master_hostname}}"
    }
  }

}

case "stub_msservers" {
  backend     = "nios"
  parallel    = true

  step {
    nios {
      fqdn           = "{{random}}.com"
      stub_from      = [{ address = "1.1.1.1", name = "{{random2}}" }]
      stub_msservers = [{ address = "10.10.10.10", is_master = false, ns_ip = "1.1.1.1", ns_name = "ns_server" }]
    }
    check = {
      "nios.stub_msservers.0.address"   = "10.10.10.10"
      "nios.stub_msservers.0.is_master" = "false"
      "nios.stub_msservers.0.ns_ip"     = "1.1.1.1"
      "nios.stub_msservers.0.ns_name"   = "ns_server"
    }
  }

  step {
    nios {
      fqdn           = "{{random}}.com"
      stub_from      = [{ address = "1.1.1.1", name = "{{random2}}" }]
      stub_msservers = [{ address = "10.0.0.0", is_master = false, ns_ip = "2.1.1.1", ns_name = "ns_server2" }]
    }
    check = {
      "nios.stub_msservers.0.address"   = "10.0.0.0"
      "nios.stub_msservers.0.is_master" = "false"
      "nios.stub_msservers.0.ns_ip"     = "2.1.1.1"
      "nios.stub_msservers.0.ns_name"   = "ns_server2"
    }
  }

}

case "zone_format_ipv4" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn        = "10.1.0.0/25"
      stub_from   = [{ address = "1.1.1.1", name = "{{random}}" }]
      zone_format = "IPV4"
      prefix      = "{{random2}}"
    }
    check = {
      "nios.zone_format" = "IPV4"
    }
  }

}

case "zone_format_ipv6" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn        = "2001:db8:85a3:8::/64"
      stub_from   = [{ address = "1.1.1.1", name = "{{random}}" }]
      zone_format = "IPV6"
    }
    check = {
      "nios.zone_format" = "IPV6"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_dns_view" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      view      = infoblox_view.test_dns_view.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

}
