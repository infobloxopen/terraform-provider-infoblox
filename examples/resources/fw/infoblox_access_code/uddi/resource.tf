// Create a basic Access Code with a single rule
resource "infoblox_access_code" "example_basic" {
  uddi = {
    name       = "example-access-code"
    activation = "2030-01-01T00:00:00Z"
    expiration = "2031-01-01T00:00:00Z"
    rules = [{
      type = "custom_list"
      data = "tf-provider-test-access-code"
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
        data = "tf-provider-test-access-code"
      },
    ]
  }
}
