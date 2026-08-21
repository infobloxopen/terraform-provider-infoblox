# Basic IP space with only the required name field.
resource "infoblox_networkview" "example" {
  uddi = {
    name = "example_ip_space"
  }
}

# IP space with a comment and resource tags.
resource "infoblox_networkview" "example_tags" {
  uddi = {
    name    = "example_ip_space_tags"
    comment = "Example IP space with tags created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

# IP space with DDNS, hostname-rewrite, ASM, and DHCP configuration.
resource "infoblox_networkview" "example_full" {
  uddi = {
    name    = "example_ip_space_full"
    comment = "Full IP space example"

    # DDNS settings (defaults shown with non-default values)
    ddns_client_update            = "server"
    ddns_conflict_resolution_mode = "no_check_with_dhcid"
    ddns_generate_name            = true
    ddns_generated_prefix         = "myhost"
    ddns_send_updates             = false
    ddns_update_on_renew          = true
    ddns_use_conflict_resolution  = false

    # Hostname rewrite settings
    hostname_rewrite_enabled = true
    hostname_rewrite_char    = "-"
    hostname_rewrite_regex   = "[^a-zA-Z0-9_.]"

    # ASM configuration (overriding defaults)
    asm_config = {
      asm_threshold       = 80
      enable              = true
      enable_notification = true
      forecast_period     = 14
      growth_factor       = 20
      growth_type         = "percent"
      history             = 30
      min_total           = 10
      min_unused          = 10
      reenable_date       = "1970-01-01T00:00:00Z"
    }

    # DHCP configuration (overriding defaults)
    dhcp_config = {
      abandoned_reclaim_time    = 3600
      abandoned_reclaim_time_v6 = 3600
      allow_unknown             = true
      allow_unknown_v6          = true
      echo_client_id            = true
      ignore_client_uid         = false
      lease_time                = 3600
      lease_time_v6             = 3600
    }

    tags = {
      Site = "location-1"
    }
  }
}
