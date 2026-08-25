// Create Network View with Basic Fields
resource "infoblox_network_view" "create_network_view" {
  uddi = {
    name = "example_network_view"
  }
}

// Create Network View with Tags
resource "infoblox_network_view" "create_network_view_with_tags" {
  uddi = {
    name    = "example_network_view_tags"
    comment = "Example Network View with tags created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

// Create Network View with Additional Fields
resource "infoblox_network_view" "create_network_view_with_additional_fields" {
  uddi = {
    name    = "example_network_view_full"
    comment = "Full Network View example"

    // DDNS settings (defaults shown with non-default values)
    ddns_client_update            = "server"
    ddns_conflict_resolution_mode = "no_check_with_dhcid"
    ddns_generate_name            = true
    ddns_generated_prefix         = "myhost"
    ddns_send_updates             = false
    ddns_update_on_renew          = true
    ddns_use_conflict_resolution  = false

    // Hostname rewrite settings
    hostname_rewrite_enabled = true
    hostname_rewrite_char    = "-"
    hostname_rewrite_regex   = "[^a-zA-Z0-9_.]"

    // ASM configuration
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

    // DHCP configuration
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
