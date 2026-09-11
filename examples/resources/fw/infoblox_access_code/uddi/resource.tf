// Create Named Lists to reference in Access Code rules
resource "infoblox_named_list" "example_1" {
  uddi = {
    name = "example-named-list-1"
    type = "custom_list"
    items_described = [
      {
        item        = "malicious-domain.example.com"
        description = "Known malicious domain"
      },
    ]
  }
}

resource "infoblox_named_list" "example_2" {
  uddi = {
    name = "example-named-list-2"
    type = "custom_list"
    items_described = [
      {
        item        = "blocked-site.example.com"
        description = "Blocked site"
      },
    ]
  }
}

// Create a basic Access Code with a single rule
resource "infoblox_access_code" "example_basic" {
  uddi = {
    name       = "example-access-code"
    activation = "2030-01-01T00:00:00Z"
    expiration = "2031-01-01T00:00:00Z"
    rules = [{
      type = "custom_list"
      data = infoblox_named_list.example_1.uddi.name
    }]
  }
}

// Create an Access Code with a description and multiple rules
resource "infoblox_access_code" "example_full" {
  uddi = {
    name        = "example-access-code-full"
    description = "Access code for remote users"
    activation  = "2030-06-01T00:00:00Z"
    expiration  = "2031-06-01T00:00:00Z"
    rules = [
      {
        type = "custom_list"
        data = infoblox_named_list.example_1.uddi.name
      },
      {
        type        = "custom_list"
        data        = infoblox_named_list.example_2.uddi.name
        description = "Secondary rule"
      },
    ]
  }
}
