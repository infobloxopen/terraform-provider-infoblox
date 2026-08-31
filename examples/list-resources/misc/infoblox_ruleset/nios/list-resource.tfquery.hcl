// List specific Rulesets using filters
list "infoblox_ruleset" "list_rulesets_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ruleset_1"
    }
  }
}

// List Rulesets with resource details included
list "infoblox_ruleset" "list_rulesets_with_resource" {
  provider         = infoblox
  include_resource = true
}
