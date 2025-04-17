//zone delegated, with fqdn and delegate_to
resource "infoblox_zone_delegated" "zone_delegated1" {
  fqdn = "subdomain.example.com"

  delegate_to {
    name    = "ns-1488.awsdns-58.org"
    address = "10.1.1.1"
  }
  delegate_to {
    name    = "ns-2034.awsdns-62.co.uk"
    address = "10.10.1.1"
  }
}

//zone delegated, with fqdn and ns_group
resource "infoblox_zone_delegated" "zone_delegated2" {
  fqdn     = "min_params.ex.org"
  ns_group = "test"
}

//zone delegated with full set of parameters
resource "infoblox_zone_delegated" "zone_delegated3" {
  fqdn          = "max_params.ex.org"
  view          = "nondefault_view"
  zone_format   = "FORWARD"
  locked        = true
  delegated_ttl = 60
  comment       = "test sample delegated zone"
  disable       = true

  delegate_to {
    name    = "te32.dz.ex.com"
    address = "10.0.0.1"
  }

  ext_attrs = jsonencode({
    "Site" = "LA"
  })
}

//zone delegated IPV6 reverse mapping zone
resource "infoblox_zone_delegated" "zone_delegated4" {
  fqdn        = "3001:db8::/64"
  comment     = "zone delegated IPV6"
  zone_format = "IPV6"

  delegate_to {
    name    = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
}
