// Create DHCP Range Template with required Fields
resource "infoblox_range_template" "range_template_required_fields" {
  nios = {
    name                = "example_range_template"
    number_of_addresses = 10
    offset              = 20
    // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
    cloud_api_compatible = true
  }
}

// Create DHCP Range Template with Additional Fields
resource "infoblox_range_template" "range_template_additional_fields" {
  nios = {
    name                = "example_range_template_additional_fields"
    number_of_addresses = 10
    offset              = 20
    // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
    cloud_api_compatible    = true
    bootfile                = "bootfile.iso"
    bootserver              = "boot_server1"
    comment                 = "Example comment for range template"
    email_list              = ["abc@example.com", "xyz@example.com"]
    failover_association    = "example_failover_association"
    server_association_type = "FAILOVER"
    high_water_mark         = 100
    high_water_mark_reset   = 20
    low_water_mark          = 10
    low_water_mark_reset    = 5
    nextserver              = "next_server1"
    options = [
      {
        name  = "domain-name-servers"
        num   = "6"
        value = "11.22.1.2,11.22.1.3"
      },
      {
        name  = "time-offset"
        num   = "2"
        value = "1000"
      },
      {
        name  = "domain-name"
        num   = "15"
        value = "aa.bb.com"
      },
    ]
    ext_attrs = {
      "Tenant ID" = "tenant-1"
    }
  }
}

// Create DHCP Range Template with filters and exclude fields
resource "infoblox_range_template" "range_template_additional_fields2" {
  nios = {
    name                = "example_range_template3"
    number_of_addresses = 60
    offset              = 70
    exclude = [
      {
        "number_of_addresses" = 10
        "offset"              = 2
        "comment"             = "Example comment for range template exclude"
      }
    ]
    fingerprint_filter_rules = [
      {
        "filter"     = "fingerprint_filter"
        "permission" = "Allow"
      }
    ]
    logic_filter_rules = [
      {
        "filter" = "example-option-filter-1"
        "type"   = "Option"
      }
    ]
    mac_filter_rules = [
      {
        "filter"     = "mac_filter"
        "permission" = "Deny"
      }
    ]
    nac_filter_rules = [
      {
        "filter"     = "nac_filter"
        "permission" = "Allow"
      }
    ]
    option_filter_rules = [
      {
        "filter"     = "example-option-filter-1"
        "permission" = "Deny"
      }
    ]
    relay_agent_filter_rules = [
      {
        "filter"     = "relay_agent_filter"
        "permission" = "Allow"
      }
    ]
  }
}
