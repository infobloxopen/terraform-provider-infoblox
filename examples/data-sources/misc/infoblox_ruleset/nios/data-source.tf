// Retrieve a specific ruleset by filters
data "infoblox_ruleset" "get_ruleset_by_name" {
  filters = {
    name = "example_ruleset_1"
  }
}

// Retrieve all rulesets
data "infoblox_ruleset" "get_all_rulesets" {}
