case "basic" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_anycast_config not yet implemented"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_anycast_config" "test_onprem_hosts" {
    uddi = {
      name = "{{random}}"
      anycast_ip_address = "{{random_ip}}"
      service = "DNS"
    }
  }
  PREREQ

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
    }
  }

}


case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "requires_resource: infoblox_anycast_config not yet implemented"
  disappears            = true
  expect_non_empty_plan = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_anycast_config" "test_onprem_hosts" {
    uddi = {
      name = "{{random}}"
      anycast_ip_address = "{{random_ip}}"
      service = "DNS"
    }
  }
  PREREQ

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccAnycastHostAnycastConfigRefs' could not be parsed (no resource block found)
case "anycast_config_refs" {
  backend     = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAnycastHostAnycastConfigRefs' could not be parsed (no resource block found)"
}


case "enable_routing" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_anycast_config not yet implemented"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_anycast_config" "test_onprem_hosts" {
    uddi = {
      name = "{{random}}"
      anycast_ip_address = "{{random_ip}}"
      service = "DNS"
    }
  }
  PREREQ

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = 6500, holddown_secs = 180, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.config_bgp.asn"             = "6500"
      "uddi.config_bgp.holddown_secs"   = "180"
      "uddi.config_bgp.neighbors.#"     = "1"
      "uddi.config_bgp.neighbors.0.asn" = "6501"
    }
  }

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "STANDARD", area = "10.10.0.1", authentication_type = "Clear", interface = "eth0", authentication_key = "YXV0aGV", hello_interval = 10, dead_interval = 40, retransmit_interval = 5, transmit_delay = 1 }
    }
    check = {
      "uddi.config_ospf.area_type"           = "STANDARD"
      "uddi.config_ospf.area"                = "10.10.0.1"
      "uddi.config_ospf.authentication_type" = "Clear"
      "uddi.config_ospf.interface"           = "eth0"
      "uddi.config_ospf.authentication_key"  = "YXV0aGV"
      "uddi.config_ospf.hello_interval"      = "10"
    }
  }

}


case "bgp" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_anycast_config not yet implemented"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_anycast_config" "test_onprem_hosts" {
    uddi = {
      name = "{{random}}"
      anycast_ip_address = "{{random_ip}}"
      service = "DNS"
    }
  }
  PREREQ

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = "BGP", holddown_secs = 6500, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.config_bgp.asn"           = "6500"
      "uddi.config_bgp.holddown_secs" = "180"
    }
  }

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = "BGP", holddown_secs = 6601, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.config_bgp.asn"           = "6601"
      "uddi.config_bgp.holddown_secs" = "200"
    }
  }

}


case "ospf" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_anycast_config not yet implemented"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_anycast_config" "test_onprem_hosts" {
    uddi = {
      name = "{{random}}"
      anycast_ip_address = "{{random_ip}}"
      service = "DNS"
    }
  }
  PREREQ

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "OSPF", area = "STANDARD", authentication_type = "10.10.0.1", interface = "Clear", authentication_key = "YXV0aGV", authentication_key_id = 1, hello_interval = "eth0", dead_interval = 10, retransmit_interval = 40, transmit_delay = 5 }
    }
    check = {
      "uddi.config_ospf.area_type"           = "STANDARD"
      "uddi.config_ospf.area"                = "10.10.0.1"
      "uddi.config_ospf.authentication_type" = "Clear"
      "uddi.config_ospf.interface"           = "eth0"
      "uddi.config_ospf.authentication_key"  = "YXV0aGV"
      "uddi.config_ospf.hello_interval"      = "10"
      "uddi.config_ospf.dead_interval"       = "40"
      "uddi.config_ospf.retransmit_interval" = "5"
      "uddi.config_ospf.transmit_delay"      = "1"
    }
  }

  step {
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "OSPF", area = "NSSA", authentication_type = "10.10.0.2", interface = "MD5", authentication_key = "YXV0aGV", authentication_key_id = 1, hello_interval = "ens160", dead_interval = 20, retransmit_interval = 50, transmit_delay = 10 }
    }
    check = {
      "uddi.config_ospf.area_type"           = "NSSA"
      "uddi.config_ospf.area"                = "10.10.0.2"
      "uddi.config_ospf.authentication_type" = "MD5"
      "uddi.config_ospf.interface"           = "ens160"
      "uddi.config_ospf.authentication_key"  = "YXV0aGV"
      "uddi.config_ospf.hello_interval"      = "20"
      "uddi.config_ospf.dead_interval"       = "50"
      "uddi.config_ospf.retransmit_interval" = "10"
      "uddi.config_ospf.transmit_delay"      = "2"
    }
  }

}
