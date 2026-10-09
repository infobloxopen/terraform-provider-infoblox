// Create a TSIG Key with Basic Fields
resource "infoblox_tsig_key" "tsig_key_with_basic_fields" {
  uddi = {
    name   = "tsig-key-basic.example.com."
    secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
  }
}

// Create a TSIG Key with Additional Fields
resource "infoblox_tsig_key" "tsig_key_with_additional_fields" {
  uddi = {
    name   = "tsig-key-additional.example.com."
    secret = "FzpyuZuQAHxLmwZVGlYcwaPB7Ow9MSWqSyyJlNR1XUc="

    // Additional Fields
    algorithm = "hmac_sha512"
    comment   = "TSIG Key created by Terraform"
    tags = {
      Site = "location-1"
    }
  }
}
