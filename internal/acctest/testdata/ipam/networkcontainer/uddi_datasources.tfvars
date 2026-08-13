# Auto-generated datasource acceptance-test cases for Networkcontainer.
case "filters" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.compartment_id", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.inheritance_parent", "uddi.name", "uddi.parent", "uddi.space"]

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

}

case "tag_filters" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.address", "uddi.cidr", "uddi.comment", "uddi.compartment_id", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.inheritance_parent", "uddi.name", "uddi.parent", "uddi.space"]

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags    = { tag1 = "{{random2}}" }
    }
  }

}
