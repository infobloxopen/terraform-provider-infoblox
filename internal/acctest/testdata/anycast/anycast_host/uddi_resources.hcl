case "basic" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
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
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
    }
    check = {
      "id"                         = "{{uddi_infra_host_legacy_id_1}}"
      "uddi.anycast_config_refs.#" = "1"
    }
  }

}


case "disappears" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
  disappears        = true
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
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
    }
  }
}

case "anycast_config_refs" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_anycast_config" "test_onprem_hosts" {
      uddi = {
        name = "{{random}}"
        anycast_ip_address = "{{random_ip}}"
        service = "DNS"
      }
    }

    resource "infoblox_anycast_config" "test_onprem_hosts2" {
          uddi = {
            name = "{{random2}}"
            anycast_ip_address = "{{random_ip2}}"
            service = "DNS"
          }
        }

    resource "infoblox_anycast_config" "test_onprem_hosts3" {
          uddi = {
            name = "{{random3}}"
            anycast_ip_address = "{{random_ip3}}"
            service = "DNS"
          }
        }
    PREREQ

  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }, { anycast_config_name = infoblox_anycast_config.test_onprem_hosts2.uddi.name }]
    }
    check = {
      "uddi.anycast_config_refs.#"                     = "2"
      "uddi.anycast_config_refs.0.anycast_config_name" = "{{random}}"
      "uddi.anycast_config_refs.1.anycast_config_name" = "{{random2}}"
    }
  }

  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts3.uddi.name }]
    }
    check = {
      "uddi.anycast_config_refs.#"                     = "1"
      "uddi.anycast_config_refs.0.anycast_config_name" = "{{random3}}"
    }
  }
}


case "enable_routing" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
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
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = 6500, holddown_secs = 180, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.anycast_config_refs.#"             = "1"
      "uddi.config_bgp.asn"                    = "6500"
      "uddi.config_bgp.holddown_secs"          = "180"
      "uddi.config_bgp.neighbors.#"            = "1"
      "uddi.config_bgp.neighbors.0.asn"        = "6501"
      "uddi.config_bgp.neighbors.0.ip_address" = "172.28.4.198"
    }
  }

  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "STANDARD", area = "10.10.0.1", authentication_type = "Clear", interface = "eth0", authentication_key = "YXV0aGV", hello_interval = 10, dead_interval = 40, retransmit_interval = 5, transmit_delay = 1 }
    }
    check = {
      "uddi.anycast_config_refs.#"           = "1"
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

}


case "bgp" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
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
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = "6500", holddown_secs = 180, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.anycast_config_refs.#"             = "1"
      "uddi.config_bgp.asn"                    = "6500"
      "uddi.config_bgp.holddown_secs"          = "180"
      "uddi.config_bgp.neighbors.0.ip_address" = "172.28.4.198"
    }
  }

  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_bgp          = { asn = "6601", holddown_secs = 200, neighbors = [{ asn = 6501, ip_address = "172.28.4.198" }] }
    }
    check = {
      "uddi.anycast_config_refs.#"             = "1"
      "uddi.config_bgp.asn"                    = "6601"
      "uddi.config_bgp.holddown_secs"          = "200"
      "uddi.config_bgp.neighbors.0.ip_address" = "172.28.4.198"
    }
  }

}


case "ospf" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"
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
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "STANDARD", area = "10.10.0.1", authentication_type = "Clear", interface = "eth0", authentication_key = "YXV0aGV", authentication_key_id = 1, hello_interval = 10, dead_interval = 40, retransmit_interval = 5, transmit_delay = 10 }
    }
    check = {
      "uddi.anycast_config_refs.#"             = "1"
      "uddi.config_ospf.area_type"             = "STANDARD"
      "uddi.config_ospf.area"                  = "10.10.0.1"
      "uddi.config_ospf.authentication_type"   = "Clear"
      "uddi.config_ospf.interface"             = "eth0"
      "uddi.config_ospf.authentication_key"    = "YXV0aGV"
      "uddi.config_ospf.authentication_key_id" = "1"
      "uddi.config_ospf.hello_interval"        = "10"
      "uddi.config_ospf.dead_interval"         = "40"
      "uddi.config_ospf.retransmit_interval"   = "5"
      "uddi.config_ospf.transmit_delay"        = "10"
    }
  }

  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      anycast_config_refs = [{ anycast_config_name = infoblox_anycast_config.test_onprem_hosts.uddi.name }]
      config_ospf         = { area_type = "NSSA", area = "10.10.0.2", authentication_type = "MD5", interface = "ens160", authentication_key = "YXV0aGV", authentication_key_id = 3, hello_interval = 20, dead_interval = 50, retransmit_interval = 10, transmit_delay = 2 }
    }
    check = {
      "uddi.anycast_config_refs.#"             = "1"
      "uddi.config_ospf.area_type"             = "NSSA"
      "uddi.config_ospf.area"                  = "10.10.0.2"
      "uddi.config_ospf.authentication_type"   = "MD5"
      "uddi.config_ospf.interface"             = "ens160"
      "uddi.config_ospf.authentication_key"    = "YXV0aGV"
      "uddi.config_ospf.authentication_key_id" = "3"
      "uddi.config_ospf.hello_interval"        = "20"
      "uddi.config_ospf.dead_interval"         = "50"
      "uddi.config_ospf.retransmit_interval"   = "10"
      "uddi.config_ospf.transmit_delay"        = "2"
    }
  }

}
