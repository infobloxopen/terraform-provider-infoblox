// Create Network View with Basic Fields
resource "infoblox_network_view" "create_network_view" {
  nios = {
    name = "example_network_view"
  }
}

// Create Network View with Additional Fields
resource "infoblox_network_view" "create_network_view_with_additional_fields" {
  nios = {
    name    = "example-network-view2"
    comment = "Example Network View with Additional Fields"

    remote_reverse_zones = [
      {
        fqdn           = "0.168.192.in-addr.arpa"
        key_type       = "NONE"
        server_address = "192.168.12.12"
      },
      {
        fqdn           = "1.168.192.in-addr.arpa"
        key_type       = "TSIG"
        server_address = "192.168.12.12"
        tsig_key_name  = "aeiou"
        tsig_key_alg   = "HMAC-SHA256"
        tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
      }
    ]

    remote_forward_zones = [
      {
        fqdn           = "fwdzone1.com"
        key_type       = "NONE"
        server_address = "192.168.12.12"
      },
      {
        fqdn           = "fwdzone2.com"
        key_type       = "TSIG"
        server_address = "192.168.12.12"
        tsig_key_name  = "aeiou"
        tsig_key_alg   = "HMAC-SHA256"
        tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
      }
    ]
    mgm_private = true

    ext_attrs = {
      Site = "location-1"
    }
  }
}
