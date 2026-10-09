// Create a Named ACL with Basic Fields
resource "infoblox_namedacl" "create_namedacl" {
  nios = {
    name    = "example_namedacl"
    comment = "Base ACL structure created for future assignment of access control entries"

    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Named ACL with Additional Fields
resource "infoblox_namedacl" "create_namedacl_with_additional_fields" {
  nios = {
    name    = "example_namedacl_advanced"
    comment = "ACL to allow or deny access to specific network resources"

    // ACL entries: address-based and TSIG-based
    access_list = [
      {
        struct     = "addressac"
        address    = "10.0.0.1"
        permission = "ALLOW"
      },
      {
        struct     = "addressac"
        address    = "10.0.0.2"
        permission = "DENY"
      },
      {
        struct        = "tsigac"
        tsig_key      = "X4oRe92t54I+T98NdQpV2w=="
        tsig_key_name = "example_tsig_key"
        tsig_key_alg  = "HMAC-SHA256"
      }
    ]

    ext_attrs = {
      Site = "location-1"
    }
  }
}
