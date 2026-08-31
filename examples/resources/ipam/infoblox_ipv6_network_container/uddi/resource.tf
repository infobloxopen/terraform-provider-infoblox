// Create an IPv6 Network Container with Basic Fields
resource "infoblox_ipv6_network_container" "example" {
  uddi = {
    address = "2001:db8::"
    cidr    = 64
    name    = "example_ipv6_network_container"
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
  }
}


// Create an IPv6 Network Container with Additional Fields
resource "infoblox_ipv6_network_container" "example_tags" {
  uddi = {
    address = "2001:db8:1::"
    cidr    = 48
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_ipv6_network_container_additional"
    comment = "Created by Terraform"
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
