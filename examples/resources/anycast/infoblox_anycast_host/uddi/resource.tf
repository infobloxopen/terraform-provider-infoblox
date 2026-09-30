# Retrieve Infra Host with Anycast Service Configured ( Required as Parent )
data "infoblox_infra_host" "parent" {
  filters = {
    display_name = "my_host"
  }
}

# Create an Anycast Configuration ( Required as Parent )
resource "infoblox_anycast_config" "example" {
  uddi = {
    name               = "anycast_config_example"
    service            = "DNS"
    anycast_ip_address = "192.1.1.1"
  }
}

# Manage an Anycast Host
resource "infoblox_anycast_config" "example_advanced" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id
  uddi = {
    # Adding the anycast config profile and enabling BGP,OSPF routing protocol
    anycast_config_refs = [
      {
        anycast_config_name = bloxone_anycast_config.example.name
        routing_protocols   = ["BGP", "OSPF"]
      }
    ]

    # Adding the BGP configuration
    config_bgp = {
      asn           = "6500"
      holddown_secs = 180
      neighbors = [
        {
          asn        = "6501"
          ip_address = "10.20.0.3"
        }
      ]
    }

    # Adding the OSPF configuration
    config_ospf = {
      area_type           = "STANDARD"
      area                = "10.0.0.1"
      authentication_type = "Clear"
      interface           = "eth0"
      authentication_key  = "YXV0aGV"
      hello_interval      = 10
      dead_interval       = 40
      retransmit_interval = 5
      transmit_delay      = 1
    }
  }
}
