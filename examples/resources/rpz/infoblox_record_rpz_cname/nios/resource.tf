// Create Record RPZ CNAME with Basic Fields
// A canonical name of "*" is the RPZ NODATA rule.
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_basic" {
  nios = {
    name      = "blocked.rpz.example.com"
    canonical = "*"
    rp_zone   = "rpz.example.com"
  }
}

// Create Record RPZ CNAME with Additional Fields
// A fully qualified canonical name is a substitution rule.
resource "infoblox_record_rpz_cname" "create_record_rpz_cname_additional" {
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
