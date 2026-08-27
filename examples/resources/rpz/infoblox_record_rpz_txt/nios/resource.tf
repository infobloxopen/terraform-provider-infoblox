// Create Record RPZ TXT with Basic Fields
resource "infoblox_record_rpz_txt" "create_record_rpz_txt_basic" {
  nios = {
    name    = "blocked.rpz.example.com"
    text    = "Example text"
    rp_zone = "rpz.example.com"
  }
}

// Create Record RPZ TXT with Additional Fields
resource "infoblox_record_rpz_txt" "create_record_rpz_txt_additional" {
  nios = {
    name    = "blocked-with-ttl.rpz.example.com"
    text    = "Example text with Additional Config"
    rp_zone = "rpz.example.com"
    view    = "default"
    ttl     = 10
    comment = "Example RPZ TXT record"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
