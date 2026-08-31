// Create an IPV4 Network Container with Basic Fields
resource "infoblox_network_container" "networkcontainer_with_basic_fields" {
  uddi = {
    address = "192.168.1.0"
    cidr    = 24
    name    = "example_address_block"
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
  }
}

// Create an IPV4 Network Container with Additional Fields
resource "infoblox_network_container" "networkcontainer_with_additional_fields" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 8
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_address_block_with_additional_fields"
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

// Create an IPV4 Network Container with Dynamic Allocation
resource "infoblox_network_container" "networkcontainer_with_dynamic_allocation" {
  uddi = {
    next_available_id = infoblox_network_container.networkcontainer_with_basic_fields.id
    cidr              = 26
    space             = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_address_block_with_dynamic_allocation"
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
