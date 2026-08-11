// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_auth_zone" {
  nios = {
    fqdn = "example-auth-zone.com"
  }
}

// Create Record SRV with Basic Fields
resource "infoblox_record_srv" "example_1" {
  nios = {
    name     = "example-srv-record.${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
    target   = "example.target.${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
    port     = 5060
    priority = 10
    weight   = 5
  }
}

// Create Record SRV with Additional Fields
resource "infoblox_record_srv" "create_with_additional_config" {
  nios = {
    name     = "example-srv-record-with-config.${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
    target   = "example_updated.target.${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
    port     = 8080
    priority = 2
    weight   = 100
    view     = "default"
    use_ttl  = true
    ttl      = 10
    creator  = "DYNAMIC"
    comment  = "Example SRV record"
    ext_attrs = {
      Site = "location-1"
    }
  }

  depends_on = [infoblox_zone_auth.parent_auth_zone]
}
