// List all Credential Groups
//
// discovery:credentialgroup exposes no searchable field, so the list resource
// supports no filters block.
list "infoblox_credential_group" "list_credential_group_all" {
  provider = infoblox
  limit    = 10
}

// List Credential Groups with resource details included
list "infoblox_credential_group" "list_credential_group_with_resource" {
  provider         = infoblox
  include_resource = true
}
