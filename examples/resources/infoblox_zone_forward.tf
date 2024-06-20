//forward mapping zone, with minimum set of parameters
resource "infoblox_zone_forward" "forward_zone_forwardTo" {
  fqdn = "min_params.ex.org"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
}

//forward zone with full set of parameters
resource "infoblox_zone_forward" "forward_zone_full_parameters" {
  fqdn = "max_params.ex.org"
  view = "nondefault_view"
  zone_format = "FORWARD"
  comment = "test sample forward zone"
  forward_to {
    name = "te32.dz.ex.com"
    address = "10.0.0.1"
  }
  forwarding_servers {
    name = "infoblox.172_28_83_216"
    forwarders_only = true
    use_override_forwarders = false
    forward_to {
      name = "cc.fwd.com"
      address = "10.1.1.1"
    }
  }
  forwarding_servers {
    name = "infoblox.172_28_83_0"
    forwarders_only = true
    use_override_forwarders = true
    forward_to {
      name = "kk.fwd.com"
      address = "10.2.1.31"
    }
  }
}

//forward zone with ns_group and external_ns_group
resource "infoblox_zone_forward" "forward_zone_nsGroup_externalNsGroup" {
  fqdn = "params_ns_ens.ex.org"
  ns_group = "test"
  external_ns_group = "stub server"
}

//forward zone with forwarding_servers and forward_to
resource "infoblox_zone_forward" "forward_zone_forwardTo_forwardingServers" {
  fqdn = "params_fs_ft.ex.org"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
  forwarding_servers {
    name = "infoblox.172_28_83_0"
    forwarders_only = true
    use_override_forwarders = true
    forward_to {
      name = "kk.fwd.com"
      address = "10.2.1.31"
    }
  }
}

//forward zone IPV4 reverse mapping zone
resource "infoblox_zone_forward" "forward_zone_IPV4_nsGroup_externalNsGroup_comment" {
  fqdn = "192.1.0.0/24"
  comment = "Forward zone IPV4"
  external_ns_group = "stub server"
  zone_format = "IPV4"
  ns_group = "test"
}

//forward zone IPV6 reverse mapping zone
resource "infoblox_zone_forward" "forward_zone_IPV6_forwardTo_forwardingServers" {
  fqdn = "3001:db8::/64"
  comment = "Forward zone IPV6"
  zone_format = "IPV6"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
  forwarding_servers {
    name = "infoblox.172_28_83_0"
    forwarders_only = true
    use_override_forwarders = true
    forward_to {
      name = "kk.fwd.com"
      address = "10.2.1.31"
    }
  }
}
