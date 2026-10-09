---
page_title: "Migrating to the Unified Provider"
subcategory: "Guides"
description: |-
  Maps resources and data sources from the legacy Infoblox, NIOS, and BloxOne providers to the unified Infoblox provider.
---

# Migrating to the Unified Provider

The Unified Infoblox provider replaces three earlier providers:

- The legacy Infoblox provider (`infobloxopen/infoblox` 2.x)
- The NIOS provider (`infobloxopen/nios`)
- The BloxOne provider (`infobloxopen/bloxone`)

The tables below map each resource and data source in those providers to its name in the unified provider. Each name applies to both the resource and the data source, where the object has one.

 Every object now has a nested `nios` or `uddi` block, and many attributes have changed. Objects that keep the same name also have a new schema. Update your configuration and re-import existing objects into state. See the [Resources and Data Sources](resources-datasources.md) guide for the backends each object supports.


## Legacy Infoblox Provider to Unified Provider

All legacy objects map to the NIOS backend.

| Legacy Provider                   | Unified Provider                  |
|-----------------------------------|-----------------------------------|
| `infoblox_network_view`           | `infoblox_network_view`           |
| `infoblox_ipv4_network_container` | `infoblox_network_container`      |
| `infoblox_ipv6_network_container` | `infoblox_ipv6_network_container` |
| `infoblox_ipv4_network`           | `infoblox_network`                |
| `infoblox_ipv6_network`           | `infoblox_ipv6_network`           |
| `infoblox_ipv4_fixed_address`     | `infoblox_fixed_address`          |
| `infoblox_ipv4_range`             | `infoblox_range`                  |
| `infoblox_ipv4_range_template`    | `infoblox_range_template`         |
| `infoblox_ipv4_shared_network`    | `infoblox_sharednetwork`          |
| `infoblox_ip_allocation`          | `infoblox_record_host`            |
| `infoblox_ip_association`         | `infoblox_ip_association`         |
| `infoblox_host_record`            | `infoblox_record_host`            |
| `infoblox_a_record`               | `infoblox_record_a`               |
| `infoblox_aaaa_record`            | `infoblox_record_aaaa`            |
| `infoblox_alias_record`           | `infoblox_record_alias`           |
| `infoblox_cname_record`           | `infoblox_record_cname`           |
| `infoblox_mx_record`              | `infoblox_record_mx`              |
| `infoblox_ns_record`              | `infoblox_record_ns`              |
| `infoblox_ptr_record`             | `infoblox_record_ptr`             |
| `infoblox_srv_record`             | `infoblox_record_srv`             |
| `infoblox_txt_record`             | `infoblox_record_txt`             |
| `infoblox_dns_view`               | `infoblox_view`                   |
| `infoblox_zone_auth`              | `infoblox_zone_auth`              |
| `infoblox_zone_delegated`         | `infoblox_zone_delegated`         |
| `infoblox_zone_forward`           | `infoblox_zone_forward`           |
| `infoblox_dtc_lbdn`               | `infoblox_dtc_lbdn`               |
| `infoblox_dtc_pool`               | `infoblox_dtc_pool`               |
| `infoblox_dtc_server`             | `infoblox_dtc_server`             |

## NIOS Provider to Unified Provider

All NIOS provider objects map to the NIOS backend.

| NIOS Provider                             | Unified Provider                              |
|-------------------------------------------|-----------------------------------------------|
| `nios_acl_namedacl`                       | `infoblox_namedacl`                           |
| `nios_cloud_aws_user`                     | `infoblox_aws_user`                           |
| `nios_dhcp_filteroption`                  | `infoblox_filteroption`                       |
| `nios_dhcp_fixed_address`                 | `infoblox_fixed_address`                      |
| `nios_dhcp_ipv6_range_template`           | `infoblox_ipv6_range_template`                |
| `nios_dhcp_ipv6filteroption`              | `infoblox_ipv6_filteroption`                  |
| `nios_dhcp_ipv6fixedaddress`              | `infoblox_ipv6_fixed_address`                 |
| `nios_dhcp_ipv6fixedaddresstemplate`      | `infoblox_ipv6_fixed_address_template`        |
| `nios_dhcp_ipv6optiondefinition`          | `infoblox_ipv6_dhcp_optiondefinition`         |
| `nios_dhcp_ipv6optionspace`               | `infoblox_ipv6_dhcp_optionspace`              |
| `nios_dhcp_ipv6sharednetwork`             | `infoblox_ipv6_sharednetwork`                 |
| `nios_dhcp_optiondefinition`              | `infoblox_dhcp_optiondefinition`              |
| `nios_dhcp_optionspace`                   | `infoblox_dhcp_optionspace`                   |
| `nios_dhcp_range`                         | `infoblox_range`                              |
| `nios_dhcp_range_template`                | `infoblox_range_template`                     |
| `nios_dhcp_shared_network`                | `infoblox_sharednetwork`                      |
| `nios_discovery_credentialgroup`          | `infoblox_discovery_credential_group`         |
| `nios_dns_nsgroup`                        | `infoblox_nsgroup`                            |
| `nios_dns_nsgroup_delegation`             | `infoblox_nsgroup_delegation`                 |
| `nios_dns_nsgroup_forwardingmember`       | `infoblox_nsgroup_forwardingmember`           |
| `nios_dns_nsgroup_forwardstubserver`      | `infoblox_nsgroup_forwardstubserver`          |
| `nios_dns_nsgroup_stubmember`             | `infoblox_nsgroup_stubmember`                 |
| `nios_dns_record_a`                       | `infoblox_record_a`                           |
| `nios_dns_record_aaaa`                    | `infoblox_record_aaaa`                        |
| `nios_dns_record_alias`                   | `infoblox_record_alias`                       |
| `nios_dns_record_caa`                     | `infoblox_record_caa`                         |
| `nios_dns_record_cname`                   | `infoblox_record_cname`                       |
| `nios_dns_record_dname`                   | `infoblox_record_dname`                       |
| `nios_dns_record_mx`                      | `infoblox_record_mx`                          |
| `nios_dns_record_naptr`                   | `infoblox_record_naptr`                       |
| `nios_dns_record_ns`                      | `infoblox_record_ns`                          |
| `nios_dns_record_ptr`                     | `infoblox_record_ptr`                         |
| `nios_dns_record_srv`                     | `infoblox_record_srv`                         |
| `nios_dns_record_txt`                     | `infoblox_record_txt`                         |
| `nios_dns_sharedrecord_a`                 | `infoblox_sharedrecord_a`                     |
| `nios_dns_sharedrecord_aaaa`              | `infoblox_sharedrecord_aaaa`                  |
| `nios_dns_sharedrecord_cname`             | `infoblox_sharedrecord_cname`                 |
| `nios_dns_sharedrecord_mx`                | `infoblox_sharedrecord_mx`                    |
| `nios_dns_sharedrecord_srv`               | `infoblox_sharedrecord_srv`                   |
| `nios_dns_sharedrecord_txt`               | `infoblox_sharedrecord_txt`                   |
| `nios_dns_sharedrecordgroup`              | `infoblox_sharedrecord_group`                 |
| `nios_dns_view`                           | `infoblox_view`                               |
| `nios_dns_zone_auth`                      | `infoblox_zone_auth`                          |
| `nios_dns_zone_delegated`                 | `infoblox_zone_delegated`                     |
| `nios_dns_zone_forward`                   | `infoblox_zone_forward`                       |
| `nios_dns_zone_rp`                        | `infoblox_zone_rp`                            |
| `nios_dns_zone_stub`                      | `infoblox_zone_stub`                          |
| `nios_record_host`                        | `infoblox_record_host`                        |
| `nios_ip_allocation`                      | `infoblox_record_host`                        |
| `nios_ip_association`                     | `infoblox_ip_association`                     |
| `nios_dtc_lbdn`                           | `infoblox_dtc_lbdn`                           |
| `nios_dtc_monitor_http`                   | `infoblox_dtc_monitor_http`                   |
| `nios_dtc_monitor_icmp`                   | `infoblox_dtc_monitor_icmp`                   |
| `nios_dtc_monitor_pdp`                    | `infoblox_dtc_monitor_pdp`                    |
| `nios_dtc_monitor_snmp`                   | `infoblox_dtc_monitor_snmp`                   |
| `nios_dtc_monitor_tcp`                    | `infoblox_dtc_monitor_tcp`                    |
| `nios_dtc_pool`                           | `infoblox_dtc_pool`                           |
| `nios_dtc_server`                         | `infoblox_dtc_server`                         |
| `nios_dtc_topology`                       | `infoblox_dtc_topology`                       |
| `nios_grid_distributionschedule`          | `infoblox_distribution_schedule`              |
| `nios_grid_extensibleattributedef`        | `infoblox_extensible_attribute_def`           |
| `nios_grid_natgroup`                      | `infoblox_nat_group`                          |
| `nios_grid_servicerestart_group`          | `infoblox_servicerestart_group`               |
| `nios_grid_upgradegroup`                  | `infoblox_upgrade_group`                      |
| `nios_ipam_bulk_hostname_template`        | `infoblox_bulk_hostname_template`             |
| `nios_ipam_ipv6network`                   | `infoblox_ipv6_network`                       |
| `nios_ipam_ipv6network_container`         | `infoblox_ipv6_network_container`             |
| `nios_ipam_network`                       | `infoblox_network`                            |
| `nios_ipam_network_container`             | `infoblox_network_container`                  |
| `nios_ipam_network_view`                  | `infoblox_network_view`                       |
| `nios_ipam_superhost`                     | `infoblox_superhost`                          |
| `nios_ipam_vlan`                          | `infoblox_vlan`                               |
| `nios_ipam_vlanrange`                     | `infoblox_vlan_range`                         |
| `nios_ipam_vlanview`                      | `infoblox_vlan_view`                          |
| `nios_misc_bfdtemplate`                   | `infoblox_bfd_template`                       |
| `nios_misc_ruleset`                       | `infoblox_ruleset`                            |
| `nios_notification_rest_endpoint`         | `infoblox_notification_rest_endpoint`         |
| `nios_notification_rule`                  | `infoblox_notification_rule`                  |
| `nios_rir_organization`                   | `infoblox_rir_organization`                   |
| `nios_rpz_record_a`                       | `infoblox_record_rpz_a`                       |
| `nios_rpz_record_a_ipaddress`             | `infoblox_record_rpz_a_ipaddress`             |
| `nios_rpz_record_aaaa`                    | `infoblox_record_rpz_aaaa`                    |
| `nios_rpz_record_aaaa_ipaddress`          | `infoblox_record_rpz_aaaa_ipaddress`          |
| `nios_rpz_record_cname`                   | `infoblox_record_rpz_cname`                   |
| `nios_rpz_record_cname_clientipaddress`   | `infoblox_record_rpz_cname_clientipaddress`   |
| `nios_rpz_record_cname_clientipaddressdn` | `infoblox_record_rpz_cname_clientipaddressdn` |
| `nios_rpz_record_cname_ipaddress`         | `infoblox_record_rpz_cname_ipaddress`         |
| `nios_rpz_record_cname_ipaddressdn`       | `infoblox_record_rpz_cname_ipaddressdn`       |
| `nios_rpz_record_naptr`                   | `infoblox_record_rpz_naptr`                   |
| `nios_rpz_record_ptr`                     | `infoblox_record_rpz_ptr`                     |
| `nios_rpz_record_txt`                     | `infoblox_record_rpz_txt`                     |
| `nios_security_admin_user`                | `infoblox_admin_user`                         |

## BloxOne Provider to Unified Provider

All BloxOne provider objects map to the UDDI backend. BloxOne data source names are plural, while unified data source names match the resource name. Where the BloxOne provider uses one object for IPv4 and IPv6, the unified provider has a separate object for each.

| BloxOne Provider                             | Unified Provider                                                                        |
|----------------------------------------------|-----------------------------------------------------------------------------------------|
| `bloxone_anycast_config`                     | `infoblox_anycast_config`                                                               |
| `bloxone_anycast_host`                       | `infoblox_anycast_host`                                                                 |
| `bloxone_cloud_discovery_provider`           | `infoblox_cloud_discovery_provider`                                                     |
| `bloxone_dhcp_fixed_address`                 | `infoblox_fixed_address` (IPv4) or `infoblox_ipv6_fixed_address` (IPv6)                 |
| `bloxone_dhcp_ha_group`                      | `infoblox_ha_group`                                                                     |
| `bloxone_dhcp_option_code`                   | `infoblox_dhcp_optiondefinition` (IPv4) or `infoblox_ipv6_dhcp_optiondefinition` (IPv6) |
| `bloxone_dhcp_option_group`                  | `infoblox_option_group`                                                                 |
| `bloxone_dhcp_option_space`                  | `infoblox_dhcp_optionspace` (IPv4) or `infoblox_ipv6_dhcp_optionspace` (IPv6)           |
| `bloxone_dns_a_record`                       | `infoblox_record_a`                                                                     |
| `bloxone_dns_aaaa_record`                    | `infoblox_record_aaaa`                                                                  |
| `bloxone_dns_caa_record`                     | `infoblox_record_caa`                                                                   |
| `bloxone_dns_cname_record`                   | `infoblox_record_cname`                                                                 |
| `bloxone_dns_dname_record`                   | `infoblox_record_dname`                                                                 |
| `bloxone_dns_https_record`                   | `infoblox_record_https`                                                                 |
| `bloxone_dns_mx_record`                      | `infoblox_record_mx`                                                                    |
| `bloxone_dns_naptr_record`                   | `infoblox_record_naptr`                                                                 |
| `bloxone_dns_ns_record`                      | `infoblox_record_ns`                                                                    |
| `bloxone_dns_ptr_record`                     | `infoblox_record_ptr`                                                                   |
| `bloxone_dns_srv_record`                     | `infoblox_record_srv`                                                                   |
| `bloxone_dns_svcb_record`                    | `infoblox_record_svcb`                                                                  |
| `bloxone_dns_txt_record`                     | `infoblox_record_txt`                                                                   |
| `bloxone_dns_acl`                            | `infoblox_namedacl`                                                                     |
| `bloxone_dns_auth_nsg`                       | `infoblox_auth_nsg`                                                                     |
| `bloxone_dns_auth_zone`                      | `infoblox_zone_auth`                                                                    |
| `bloxone_dns_delegation`                     | `infoblox_zone_delegated`                                                               |
| `bloxone_dns_forward_nsg`                    | `infoblox_forward_nsg`                                                                  |
| `bloxone_dns_forward_zone`                   | `infoblox_zone_forward`                                                                 |
| `bloxone_dns_host`                           | `infoblox_dns_host`                                                                     |
| `bloxone_dns_server`                         | `infoblox_dns_server`                                                                   |
| `bloxone_dns_view`                           | `infoblox_view`                                                                         |
| `bloxone_federation_federated_realm`         | `infoblox_federated_realm`                                                              |
| `bloxone_infra_host`                         | `infoblox_infra_host`                                                                   |
| `bloxone_infra_join_token`                   | `infoblox_join_token`                                                                   |
| `bloxone_infra_service`                      | `infoblox_infra_service`                                                                |
| `bloxone_ipam_address`                       | `infoblox_address`                                                                      |
| `bloxone_ipam_address_block`                 | `infoblox_network_container` (IPv4) or `infoblox_ipv6_network_container` (IPv6)         |
| `bloxone_ipam_host`                          | `infoblox_ipam_host`                                                                    |
| `bloxone_ipam_ip_space`                      | `infoblox_network_view`                                                                 |
| `bloxone_ipam_range`                         | `infoblox_range` (IPv4)                                                                 |
| `bloxone_ipam_subnet`                        | `infoblox_network` (IPv4) or `infoblox_ipv6_network` (IPv6)                             |
| `bloxone_ipam_next_available_address_blocks` | `infoblox_next_available_address_blocks`                                                |
| `bloxone_ipam_next_available_ips`            | `infoblox_next_available_ips`                                                           |
| `bloxone_ipam_next_available_subnets`        | `infoblox_next_available_subnets`                                                       |
| `bloxone_keys_tsig`                          | `infoblox_tsig_key`                                                                     |
| `bloxone_td_access_code`                     | `infoblox_access_code`                                                                  |
| `bloxone_td_application_filter`              | `infoblox_application_filter`                                                           |
| `bloxone_td_category_filter`                 | `infoblox_category_filter`                                                              |
| `bloxone_td_custom_redirect`                 | `infoblox_custom_redirect`                                                              |
| `bloxone_td_internal_domain_list`            | `infoblox_internal_domain_list`                                                         |
| `bloxone_td_named_list`                      | `infoblox_named_list`                                                                   |
| `bloxone_td_network_list`                    | `infoblox_network_list`                                                                 |
| `bloxone_td_security_policy`                 | `infoblox_security_policy`                                                              |
