// GSS-TSIG/Kerberos keys are created by uploading a keytab, not by Terraform, so
// they are exposed as a data source only. Every field except comment and tags is
// read-only.

// Retrieve the Kerberos keys for a specific principal
data "infoblox_kerberos_key" "get_kerberos_key_using_filters" {
  filters = {
    principal = "DNS/ns.b1ddi.example.com"
  }
}

// Retrieve specific Kerberos keys using Tags
data "infoblox_kerberos_key" "get_kerberos_key_using_tag_filters" {
  tag_filters = {
    tag1 = "value1"
  }
}

// Retrieve all Kerberos keys
data "infoblox_kerberos_key" "get_all_kerberos_keys" {}
