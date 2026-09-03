// Create an RPZ CNAME Client IP Address DN Record with Basic Fields
resource "infoblox_record_rpz_cname_clientipaddressdn" "basic" {
  nios = {
    name      = "10.10.0.0/24.rpz-test.infoblox.com"
    canonical = "block.example.com"
    rp_zone   = "rpz-test.infoblox.com"
  }
}

// Create an RPZ CNAME Client IP Address DN Record with Additional Fields
resource "infoblox_record_rpz_cname_clientipaddressdn" "additional_fields" {
  nios = {
    name      = "192.168.1.0/24.rpz-test.infoblox.com"
    canonical = "redirect.example.com"
    rp_zone   = "rpz-test.infoblox.com"

    // Additional Fields
    comment = "Block client IP network from resolving"
    disable = false
    ttl     = 3600

    // Extensible Attributes
    ext_attrs = {
      Site = "datacenter-1"
    }
  }
}
