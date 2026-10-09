// Create an RPZ Zone (Required as Parent)
resource "infoblox_zone_rp" "example" {
  nios = {
    fqdn = "rpz.example.com"
  }
}

// Create Record RPZ CNAME - Block Domain (No Such Domain / NXDOMAIN rule)
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_nxdomain" {
  nios = {
    name      = "blocked.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = ""
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ CNAME - Block Domain (No Data rule)
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_nodata" {
  nios = {
    name      = "nodata.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "*"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
  }
}

// Create Record RPZ CNAME - Substitution rule with additional fields
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_substitution" {
  nios = {
    name      = "substituted.${infoblox_zone_rp.example.nios.fqdn}"
    canonical = "walled-garden.example.com"
    rp_zone   = infoblox_zone_rp.example.nios.fqdn
    ttl       = 10
    comment   = "Example RPZ CNAME record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
