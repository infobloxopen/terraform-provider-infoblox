// Retrieve all Discovery Credential Groups
//
// discovery:credentialgroup exposes no searchable field, so an unfiltered read
// is the only form this data source supports.
data "infoblox_credential_group" "get_all_discovery_credentialgroups" {}
