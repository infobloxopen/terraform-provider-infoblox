// Create a Discovery Credential Group with Basic Fields
resource "infoblox_credential_group" "discovery_credentialgroup_with_basic_fields" {
  nios = {
    name = "example_credential_group"
  }
}
