// List Custom Redirects with resource details included
list "infoblox_custom_redirect" "list_custom_redirect_with_resource" {
  provider         = infoblox
  include_resource = true
}
