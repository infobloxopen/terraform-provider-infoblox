// Create Record RPZ CNAME - Block Domain (No Such Domain / NXDOMAIN rule)
// canonical = "" is the most common RPZ CNAME configuration.
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_nxdomain" {
  nios = {
    name      = "blocked.rpz.example.com"
    canonical = ""
    rp_zone   = "rpz.example.com"
  }
}

// Create Record RPZ CNAME - Block Domain (No Data rule)
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_nodata" {
  nios = {
    name      = "nodata.rpz.example.com"
    canonical = "*"
    rp_zone   = "rpz.example.com"
  }
}

// Create Record RPZ CNAME - Substitution rule with additional fields
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_substitution" {
  nios = {
    name      = "substituted.rpz.example.com"
    canonical = "walled-garden.example.com"
    rp_zone   = "rpz.example.com"
    view      = "default"
    ttl       = 10
    comment   = "Example RPZ CNAME record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
