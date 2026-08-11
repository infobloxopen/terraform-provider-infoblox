resource "infoblox_ipv6networkcontainer" "example" {
  uddi = {
    address = "192.168.1.0"
    cidr    = 24
    name    = "example_address_block"
  }
}

resource "infoblox_ipv6networkcontainer" "example_tags" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 8

    // Other optional fields
    name    = "example_address_block_tags"
    comment = "Example address block with tags created by the terraform provider"
    tags = {
      Site = "location-1"
    }
    asm_config = {
      asm_threshold       = 90
      enable              = "true"
      enable_notification = "true"
      forecast_period     = 10
      growth_factor       = 10
      growth_type         = "percent"
      history             = 30
      min_total           = 2
      min_unused          = 10
      reenable_date       = "2024-01-24T10:10:00+00:00"
    }
    dhcp_config = {
      allow_unknown = true
      ignore_list = [
        {
          type  = "hardware"
          value = "aa:bb:cc:dd:ee:ff"
        },
        {
          type  = "client_text"
          value = "001d.a18b.36d0"
        },
        {
          type  = "client_hex"
          value = "333964392D4769302F31"
        }
      ]
    }
  }
}

// Next available address block
resource "infoblox_ipv6networkcontainer" "example_na_ab" {
  uddi = {
    next_available_id = infoblox_ipv6networkcontainer.example.id
    cidr              = 26

    // Other optional fields
    name    = "example_address_block_tags"
    comment = "Example address block with tags created by the terraform provider"
    tags = {
      Site = "location-1"
    }
    asm_config = {
      asm_threshold       = 90
      enable              = "true"
      enable_notification = "true"
      forecast_period     = 10
      growth_factor       = 10
      growth_type         = "percent"
      history             = 30
      min_total           = 2
      min_unused          = 10
      reenable_date       = "2024-01-24T10:10:00+00:00"
    }
    dhcp_config = {
      allow_unknown = true
      ignore_list = [
        {
          type  = "hardware"
          value = "aa:bb:cc:dd:ee:ff"
        },
        {
          type  = "client_text"
          value = "001d.a18b.36d0"
        },
        {
          type  = "client_hex"
          value = "333964392D4769302F31"
        }
      ]
    }
  }
}
