// Retrieve an Infra Host
data "infoblox_infra_host" "example_by_attribute" {
  filters = {
    display_name = "example_host"
  }
}

// Create DNS Server ( Required as Parent )
resource "infoblox_dns_server" "test" {
  uddi = {
    name = "example_dns_server"
  }
}

// Manage a DNS Host
resource "infoblox_dns_host" "example_dns_host" {
  id = "dns/host/${data.infoblox_infra_host.example_by_attribute.results.0.uddi.legacy_id}"
  uddi = {
    server        = infoblox_dns_server.test.id
    absolute_name = "example_dns_host."
  }
}
