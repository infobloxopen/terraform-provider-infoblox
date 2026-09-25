// Create a Network View (Required as Parent)
resource "infoblox_network_view" "example_network_view" {
  nios = {
    name = "example-range-view"
  }
}

// Create an IPv4 Network inside the Network View (Required as Parent of the Range)
resource "infoblox_network" "example_network" {
  nios = {
    network      = "10.0.0.0/24"
    network_view = infoblox_network_view.example_network_view.nios.name
  }
}

// Create a DHCP Range with Basic Fields
resource "infoblox_range" "range_basic_fields" {
  nios = {
    start_addr   = "10.0.0.10"
    end_addr     = "10.0.0.20"
    network      = infoblox_network.example_network.nios.network
    network_view = infoblox_network.example_network.nios.network_view
    comment      = "Example Range created by the terraform provider"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a DHCP Range with Additional Fields
resource "infoblox_range" "range_additional_fields" {
  nios = {
    name         = "example_range"
    start_addr   = "10.0.0.30"
    end_addr     = "10.0.0.50"
    network      = infoblox_network.example_network.nios.network
    network_view = infoblox_network.example_network.nios.network_view
    comment      = "Example Range with additional fields"

    // Addresses excluded from the range
    exclude = [
      {
        start_address = "10.0.0.35"
        end_address   = "10.0.0.40"
        comment       = "Reserved for static assignment"
      }
    ]

    // DHCP options served to clients of this range
    options = [
      {
        name  = "domain-name-servers"
        num   = 6
        value = "10.0.0.2,10.0.0.3"
      },
      {
        name  = "time-offset"
        num   = 2
        value = "1000"
      }
    ]

    // DHCP threshold monitoring
    enable_dhcp_thresholds = true
    high_water_mark        = 90
    low_water_mark         = 5
    enable_email_warnings  = true
    email_list             = ["admin@infoblox.com"]

    // Client filtering
    deny_all_clients = false
    deny_bootp       = false
    unknown_clients  = "Allow"
    known_clients    = "Allow"

    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a DHCP Range served by a Grid Member
resource "infoblox_range" "range_with_member" {
  nios = {
    start_addr   = "10.0.0.60"
    end_addr     = "10.0.0.70"
    network      = infoblox_network.example_network.nios.network
    network_view = infoblox_network.example_network.nios.network_view
    comment      = "Example Range assigned to a grid member"

    # server_association_type = "MEMBER"
    # member = {
    #   name = "infoblox.172_28_83_113" // Replace with the name of a Grid Member
    # }
  }
}

	  terraform {
	    required_providers {
	      infoblox = {
	        source  = "infobloxopen/infoblox"
	        version = "0.0.1"
	      }
	    }
	  }
	  
	  provider "infoblox" {
	    nios = {
	      host_url = "https://172.28.82.8"
	      username = "admin"
	      password = "Infoblox@123"
	    }
}