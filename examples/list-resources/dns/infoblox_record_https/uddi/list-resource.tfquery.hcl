// List specific HTTPS Records using filters
list "infoblox_record_https" "list_record_https_using_filters" {
  provider = infoblox
  config {
    filters = {
      absolute_name_spec = "abc.example.com"
    }
  }
  limit = 10
}

// List specific HTTPS Records using Tags
list "infoblox_record_https" "list_record_https_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List HTTPS Records with resource details included
list "infoblox_record_https" "list_record_https_with_resource" {
  provider         = infoblox
  include_resource = true
}
