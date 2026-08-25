// Create a Ruleset with Basic Fields
resource "infoblox_ruleset" "ruleset_basic" {
  nios = {
    name = "example_ruleset_1"
    type = "BLACKLIST"
  }
}

// Create a Ruleset with Additional Fields
resource "infoblox_ruleset" "ruleset_with_additional_fields" {
  nios = {
    name     = "example_ruleset_2"
    type     = "NXDOMAIN"
    comment  = "This ruleset handles NXDOMAIN redirection"
    disabled = false

    nxdomain_rules = [
      {
        action  = "PASS"
        pattern = "example.com"
      },
      {
        action  = "REDIRECT"
        pattern = "test.org"
      }
    ]
  }
}
