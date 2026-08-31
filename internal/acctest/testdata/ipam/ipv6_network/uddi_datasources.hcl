# Auto-generated datasource acceptance-test cases for Ipv6network.
case "filters" {
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  filter {
    type   = "filters"
    values = {
      address = "uddi.address"
      space   = "uddi.space"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.dhcp_host", "uddi.disable_dhcp", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.inheritance_parent", "uddi.name", "uddi.parent", "uddi.rebind_time", "uddi.renew_time", "uddi.space"]

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.dhcp_host", "uddi.disable_dhcp", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.inheritance_parent", "uddi.name", "uddi.parent", "uddi.rebind_time", "uddi.renew_time", "uddi.space"]

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags    = { tag1 = "{{random}}" }
    }
  }

}
