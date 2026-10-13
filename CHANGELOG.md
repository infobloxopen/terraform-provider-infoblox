# Changelog

## [v3.0.0-beta](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v3.0.0-beta) (2026-10-13)

- Infoblox Provider (v3.x.x) manages both NIOS and Universal DDI (UDDI) from a single provider. It replaces the Legacy Infoblox Provider (v2.x.x), the NIOS Provider (`infobloxopen/nios`), and the BloxOne Provider (`infobloxopen/bloxone`).
- Objects carried over from the Legacy Infoblox Provider (v2.x.x) have a new schema, and many have been renamed. For example, `infoblox_a_record` is now `infoblox_record_a`. Each resource and data source also has a `nios` or `uddi` block for its attributes. See the [Object Mapping](docs/guides/object-mapping.md) guide for the full list of renames.
- The provider now uses [infoblox-nios-go-client](https://github.com/infobloxopen/infoblox-nios-go-client) for NIOS API calls, replacing [infoblox-go-client](https://github.com/infobloxopen/infoblox-go-client).
- The provider configuration has changed. Credentials now go inside a `nios` or `uddi` block, and one provider block can use only one set of credentials at a time. To manage both, use two provider blocks with an alias.
- **v3.0.0 is not backward compatible with v2.x.** Existing state cannot be upgraded in place, so you need to remove the old resources from state and import them again. Your objects in NIOS and UDDI are not changed. See the [Migration Guide](MIGRATION.md) and the [Object Mapping](docs/guides/object-mapping.md) guide.

### Supported Resources and Datasources

All resources and data sources below are available in this release. Objects carried over from v2.x.x are listed under their new names.

#### ACL

- `infoblox_namedacl` : Manage Named Access Control Lists and retrieve existing Named Access Control Lists. _(NIOS, UDDI)_ ([#583](https://github.com/infobloxopen/terraform-provider-infoblox/pull/583))

#### Anycast

- `infoblox_anycast_config` : Manage Anycast Configurations and retrieve existing Anycast Configurations. _(UDDI)_ ([#670](https://github.com/infobloxopen/terraform-provider-infoblox/pull/670))
- `infoblox_anycast_host` : Manage Anycast Hosts. _(UDDI)_ ([#682](https://github.com/infobloxopen/terraform-provider-infoblox/pull/682))

#### Cloud

- `infoblox_aws_user` : Manage AWS Users and retrieve existing AWS Users. _(NIOS)_ ([#644](https://github.com/infobloxopen/terraform-provider-infoblox/pull/644))

#### Cloud Discovery

- `infoblox_cloud_discovery_provider` : Manage Cloud Discovery Providers and retrieve existing Cloud Discovery Providers. _(UDDI)_ ([#687](https://github.com/infobloxopen/terraform-provider-infoblox/pull/687))

#### DHCP

- `infoblox_dhcp_optiondefinition` : Manage DHCP Option Definitions and retrieve existing DHCP Option Definitions. _(NIOS, UDDI)_ ([#572](https://github.com/infobloxopen/terraform-provider-infoblox/pull/572))
- `infoblox_dhcp_optionspace` : Manage DHCP Option Spaces and retrieve existing DHCP Option Spaces. _(NIOS, UDDI)_ ([#572](https://github.com/infobloxopen/terraform-provider-infoblox/pull/572))
- `infoblox_filteroption` : Manage DHCP Filter Options and retrieve existing DHCP Filter Options. _(NIOS, UDDI)_ ([#544](https://github.com/infobloxopen/terraform-provider-infoblox/pull/544))
- `infoblox_fixed_address` : Manage DHCP Fixed Addresses (IPv4) and retrieve existing DHCP Fixed Addresses (IPv4). _(NIOS, UDDI)_ ([#645](https://github.com/infobloxopen/terraform-provider-infoblox/pull/645))
- `infoblox_ha_group` : Manage DHCP HA Groups and retrieve existing DHCP HA Groups. _(UDDI)_ ([#563](https://github.com/infobloxopen/terraform-provider-infoblox/pull/563))
- `infoblox_hardware_filter` : Manage DHCP Hardware Filters and retrieve existing DHCP Hardware Filters. _(UDDI)_ ([#675](https://github.com/infobloxopen/terraform-provider-infoblox/pull/675))
- `infoblox_ipv6_dhcp_optiondefinition` : Manage DHCP IPv6 Option Definitions and retrieve existing DHCP IPv6 Option Definitions. _(NIOS, UDDI)_ ([#576](https://github.com/infobloxopen/terraform-provider-infoblox/pull/576))
- `infoblox_ipv6_dhcp_optionspace` : Manage DHCP IPv6 Option Spaces and retrieve existing DHCP IPv6 Option Spaces. _(NIOS, UDDI)_ ([#576](https://github.com/infobloxopen/terraform-provider-infoblox/pull/576))
- `infoblox_ipv6_filteroption` : Manage DHCP IPv6 Filter Options and retrieve existing DHCP IPv6 Filter Options. _(NIOS, UDDI)_ ([#643](https://github.com/infobloxopen/terraform-provider-infoblox/pull/643))
- `infoblox_ipv6_fixed_address` : Manage DHCP IPv6 Fixed Addresses and retrieve existing DHCP IPv6 Fixed Addresses. _(NIOS, UDDI)_ ([#602](https://github.com/infobloxopen/terraform-provider-infoblox/pull/602))
- `infoblox_ipv6_fixed_address_template` : Manage DHCP IPv6 Fixed Address Templates and retrieve existing DHCP IPv6 Fixed Address Templates. _(NIOS)_ ([#609](https://github.com/infobloxopen/terraform-provider-infoblox/pull/609))
- `infoblox_ipv6_range_template` : Manage DHCP IPv6 Range Templates and retrieve existing DHCP IPv6 Range Templates. _(NIOS)_ ([#611](https://github.com/infobloxopen/terraform-provider-infoblox/pull/611))
- `infoblox_ipv6_sharednetwork` : Manage DHCP IPv6 Shared Networks and retrieve existing DHCP IPv6 Shared Networks. _(NIOS)_ ([#646](https://github.com/infobloxopen/terraform-provider-infoblox/pull/646))
- `infoblox_option_group` : Manage DHCP Option Groups and retrieve existing DHCP Option Groups. _(UDDI)_ ([#659](https://github.com/infobloxopen/terraform-provider-infoblox/pull/659))
- `infoblox_range` : Manage DHCP Ranges (IPv4) and retrieve existing DHCP Ranges (IPv4). _(NIOS, UDDI)_ ([#671](https://github.com/infobloxopen/terraform-provider-infoblox/pull/671))
- `infoblox_range_template` : Manage DHCP Range Templates (IPv4) and retrieve existing DHCP Range Templates (IPv4). _(NIOS)_ ([#623](https://github.com/infobloxopen/terraform-provider-infoblox/pull/623))
- `infoblox_sharednetwork` : Manage DHCP Shared Networks (IPv4) and retrieve existing DHCP Shared Networks (IPv4). _(NIOS)_ ([#599](https://github.com/infobloxopen/terraform-provider-infoblox/pull/599))

#### Discovery

- `infoblox_discovery_credential_group` : Manage Discovery Credential Groups and retrieve existing Discovery Credential Groups. _(NIOS)_ ([#663](https://github.com/infobloxopen/terraform-provider-infoblox/pull/663))

#### DNS

- `infoblox_auth_nsg` : Manage Authoritative DNS Server Groups and retrieve existing Authoritative DNS Server Groups. _(UDDI)_ ([#551](https://github.com/infobloxopen/terraform-provider-infoblox/pull/551))
- `infoblox_dns_host` : Manage Configures existing DNS Hosts and retrieve existing DNS Hosts. _(UDDI)_ ([#690](https://github.com/infobloxopen/terraform-provider-infoblox/pull/690))
- `infoblox_dns_server` : Manage DNS Servers and retrieve existing DNS Servers. _(UDDI)_ ([#566](https://github.com/infobloxopen/terraform-provider-infoblox/pull/566))
- `infoblox_forward_nsg` : Manage Forward DNS Server Groups and retrieve existing Forward DNS Server Groups. _(UDDI)_ ([#558](https://github.com/infobloxopen/terraform-provider-infoblox/pull/558))
- `infoblox_ip_association` : Manage an IP Association for a Host Record. _(NIOS)_ ([#630](https://github.com/infobloxopen/terraform-provider-infoblox/pull/630))
- `infoblox_nsgroup` : Manage DNS NS Groups and retrieve existing DNS NS Groups. _(NIOS)_ ([#552](https://github.com/infobloxopen/terraform-provider-infoblox/pull/552))
- `infoblox_nsgroup_delegation` : Manage DNS NS Group Delegations and retrieve existing DNS NS Group Delegations. _(NIOS)_ ([#618](https://github.com/infobloxopen/terraform-provider-infoblox/pull/618))
- `infoblox_nsgroup_forwardingmember` : Manage DNS NS Group Forwarding Members and retrieve existing DNS NS Group Forwarding Members. _(NIOS)_ ([#613](https://github.com/infobloxopen/terraform-provider-infoblox/pull/613))
- `infoblox_nsgroup_forwardstubserver` : Manage DNS NS Group Forward Stub Servers and retrieve existing DNS NS Group Forward Stub Servers. _(NIOS)_ ([#615](https://github.com/infobloxopen/terraform-provider-infoblox/pull/615))
- `infoblox_nsgroup_stubmember` : Manage DNS NS Group Stub Members and retrieve existing DNS NS Group Stub Members. _(NIOS)_ ([#614](https://github.com/infobloxopen/terraform-provider-infoblox/pull/614))
- `infoblox_record_a` : Manage DNS A Records and retrieve existing DNS A Records. _(NIOS, UDDI)_ ([#519](https://github.com/infobloxopen/terraform-provider-infoblox/pull/519))
- `infoblox_record_aaaa` : Manage DNS AAAA Records and retrieve existing DNS AAAA Records. _(NIOS, UDDI)_ ([#525](https://github.com/infobloxopen/terraform-provider-infoblox/pull/525))
- `infoblox_record_alias` : Manage DNS ALIAS Records and retrieve existing DNS ALIAS Records. _(NIOS)_ ([#573](https://github.com/infobloxopen/terraform-provider-infoblox/pull/573))
- `infoblox_record_caa` : Manage DNS CAA Records and retrieve existing DNS CAA Records. _(NIOS, UDDI)_ ([#527](https://github.com/infobloxopen/terraform-provider-infoblox/pull/527))
- `infoblox_record_cname` : Manage DNS CNAME Records and retrieve existing DNS CNAME Records. _(NIOS, UDDI)_ ([#526](https://github.com/infobloxopen/terraform-provider-infoblox/pull/526))
- `infoblox_record_dname` : Manage DNS DNAME Records and retrieve existing DNS DNAME Records. _(NIOS, UDDI)_ ([#530](https://github.com/infobloxopen/terraform-provider-infoblox/pull/530))
- `infoblox_record_host` : Manage DNS Host Records and retrieve existing DNS Host Records. _(NIOS)_ ([#630](https://github.com/infobloxopen/terraform-provider-infoblox/pull/630))
- `infoblox_record_https` : Manage DNS HTTPS Records and retrieve existing DNS HTTPS Records. _(UDDI)_ ([#580](https://github.com/infobloxopen/terraform-provider-infoblox/pull/580))
- `infoblox_record_mx` : Manage DNS MX Records and retrieve existing DNS MX Records. _(NIOS, UDDI)_ ([#541](https://github.com/infobloxopen/terraform-provider-infoblox/pull/541))
- `infoblox_record_naptr` : Manage DNS NAPTR Records and retrieve existing DNS NAPTR Records. _(NIOS, UDDI)_ ([#531](https://github.com/infobloxopen/terraform-provider-infoblox/pull/531))
- `infoblox_record_ns` : Manage DNS NS Records and retrieve existing DNS NS Records. _(NIOS, UDDI)_ ([#532](https://github.com/infobloxopen/terraform-provider-infoblox/pull/532))
- `infoblox_record_ptr` : Manage DNS PTR Records and retrieve existing DNS PTR Records. _(NIOS, UDDI)_ ([#540](https://github.com/infobloxopen/terraform-provider-infoblox/pull/540))
- `infoblox_record_srv` : Manage DNS SRV Records and retrieve existing DNS SRV Records. _(NIOS, UDDI)_ ([#536](https://github.com/infobloxopen/terraform-provider-infoblox/pull/536))
- `infoblox_record_svcb` : Manage DNS SVCB Records and retrieve existing DNS SVCB Records. _(UDDI)_ ([#595](https://github.com/infobloxopen/terraform-provider-infoblox/pull/595))
- `infoblox_record_txt` : Manage DNS TXT Records and retrieve existing DNS TXT Records. _(NIOS, UDDI)_ ([#528](https://github.com/infobloxopen/terraform-provider-infoblox/pull/528))
- `infoblox_sharedrecord_a` : Manage DNS Shared A Records and retrieve existing DNS Shared A Records. _(NIOS)_ ([#568](https://github.com/infobloxopen/terraform-provider-infoblox/pull/568))
- `infoblox_sharedrecord_aaaa` : Manage DNS Shared AAAA Records and retrieve existing DNS Shared AAAA Records. _(NIOS)_ ([#571](https://github.com/infobloxopen/terraform-provider-infoblox/pull/571))
- `infoblox_sharedrecord_cname` : Manage DNS Shared CNAME Records and retrieve existing DNS Shared CNAME Records. _(NIOS)_ ([#642](https://github.com/infobloxopen/terraform-provider-infoblox/pull/642))
- `infoblox_sharedrecord_group` : Manage DNS Shared Record Groups and retrieve existing DNS Shared Record Groups. _(NIOS)_ ([#608](https://github.com/infobloxopen/terraform-provider-infoblox/pull/608))
- `infoblox_sharedrecord_mx` : Manage DNS Shared MX Records and retrieve existing DNS Shared MX Records. _(NIOS)_ ([#626](https://github.com/infobloxopen/terraform-provider-infoblox/pull/626))
- `infoblox_sharedrecord_srv` : Manage DNS Shared SRV Records and retrieve existing DNS Shared SRV Records. _(NIOS)_ ([#641](https://github.com/infobloxopen/terraform-provider-infoblox/pull/641))
- `infoblox_sharedrecord_txt` : Manage DNS Shared TXT Records and retrieve existing DNS Shared TXT Records. _(NIOS)_ ([#603](https://github.com/infobloxopen/terraform-provider-infoblox/pull/603))
- `infoblox_view` : Manage DNS Views and retrieve existing DNS Views. _(NIOS, UDDI)_ ([#523](https://github.com/infobloxopen/terraform-provider-infoblox/pull/523))
- `infoblox_zone_auth` : Manage Authoritative DNS Zones and retrieve existing Authoritative DNS Zones. _(NIOS, UDDI)_ ([#524](https://github.com/infobloxopen/terraform-provider-infoblox/pull/524))
- `infoblox_zone_delegated` : Manage Delegated DNS Zones and retrieve existing Delegated DNS Zones. _(NIOS, UDDI)_ ([#579](https://github.com/infobloxopen/terraform-provider-infoblox/pull/579))
- `infoblox_zone_forward` : Manage Forwarding DNS Zones and retrieve existing Forwarding DNS Zones. _(NIOS, UDDI)_ ([#575](https://github.com/infobloxopen/terraform-provider-infoblox/pull/575))
- `infoblox_zone_rp` : Manage DNS Response Policy Zones and retrieve existing DNS Response Policy Zones. _(NIOS)_ ([#591](https://github.com/infobloxopen/terraform-provider-infoblox/pull/591))
- `infoblox_zone_stub` : Manage DNS Stub Zones and retrieve existing DNS Stub Zones. _(NIOS)_ ([#610](https://github.com/infobloxopen/terraform-provider-infoblox/pull/610))

#### DTC

- `infoblox_dtc_lbdn` : Manage DTC LBDNs and retrieve existing DTC LBDNs. _(NIOS, UDDI)_ ([#569](https://github.com/infobloxopen/terraform-provider-infoblox/pull/569))
- `infoblox_dtc_monitor_http` : Manage DTC HTTP Monitors and retrieve existing DTC HTTP Monitors. _(NIOS, UDDI)_ ([#634](https://github.com/infobloxopen/terraform-provider-infoblox/pull/634))
- `infoblox_dtc_monitor_icmp` : Manage DTC ICMP Monitors and retrieve existing DTC ICMP Monitors. _(NIOS, UDDI)_ ([#621](https://github.com/infobloxopen/terraform-provider-infoblox/pull/621))
- `infoblox_dtc_monitor_pdp` : Manage DTC PDP Monitors and retrieve existing DTC PDP Monitors. _(NIOS, UDDI)_ ([#620](https://github.com/infobloxopen/terraform-provider-infoblox/pull/620))
- `infoblox_dtc_monitor_snmp` : Manage DTC SNMP Monitors and retrieve existing DTC SNMP Monitors. _(NIOS, UDDI)_ ([#648](https://github.com/infobloxopen/terraform-provider-infoblox/pull/648))
- `infoblox_dtc_monitor_tcp` : Manage DTC TCP Monitors and retrieve existing DTC TCP Monitors. _(NIOS, UDDI)_ ([#619](https://github.com/infobloxopen/terraform-provider-infoblox/pull/619))
- `infoblox_dtc_pool` : Manage DTC Pools and retrieve existing DTC Pools. _(NIOS, UDDI)_ ([#559](https://github.com/infobloxopen/terraform-provider-infoblox/pull/559))
- `infoblox_dtc_server` : Manage DTC Servers and retrieve existing DTC Servers. _(NIOS, UDDI)_ ([#546](https://github.com/infobloxopen/terraform-provider-infoblox/pull/546))
- `infoblox_dtc_topology` : Manage DTC Topologies and retrieve existing DTC Topologies. _(NIOS, UDDI)_ ([#631](https://github.com/infobloxopen/terraform-provider-infoblox/pull/631))

#### Threat Defense (FW)

- `infoblox_access_code` : Manage Access Codes and retrieve existing Access Codes. _(UDDI)_ ([#616](https://github.com/infobloxopen/terraform-provider-infoblox/pull/616))
- `infoblox_application_filter` : Manage Application Filters and retrieve existing Application Filters. _(UDDI)_ ([#651](https://github.com/infobloxopen/terraform-provider-infoblox/pull/651))
- `infoblox_category_filter` : Manage Category Filters and retrieve existing Category Filters. _(UDDI)_ ([#683](https://github.com/infobloxopen/terraform-provider-infoblox/pull/683))
- `infoblox_internal_domain_list` : Manage Internal Domain Lists and retrieve existing Internal Domain Lists. _(UDDI)_ ([#685](https://github.com/infobloxopen/terraform-provider-infoblox/pull/685))
- `infoblox_named_list` : Manage Named Lists (Custom Lists) and retrieve existing Named Lists (Custom Lists). _(UDDI)_ ([#625](https://github.com/infobloxopen/terraform-provider-infoblox/pull/625))
- `infoblox_network_list` : Manage Network Lists (External Networks) and retrieve existing Network Lists (External Networks). _(UDDI)_ ([#650](https://github.com/infobloxopen/terraform-provider-infoblox/pull/650))
- `infoblox_security_policy` : Manage Security Policies and retrieve existing Security Policies. _(UDDI)_ ([#655](https://github.com/infobloxopen/terraform-provider-infoblox/pull/655))

#### Grid

- `infoblox_distribution_schedule` : Manage Configures the Grid Distribution Schedule and retrieve existing Retrieves information about the current Grid Distribution Schedule. _(NIOS)_ ([#691](https://github.com/infobloxopen/terraform-provider-infoblox/pull/691))
- `infoblox_extensible_attribute_def` : Manage Extensible Attribute Definitions and retrieve existing Extensible Attribute Definitions. _(NIOS)_ ([#567](https://github.com/infobloxopen/terraform-provider-infoblox/pull/567))
- `infoblox_nat_group` : Manage Grid NAT Groups and retrieve existing Grid NAT Groups. _(NIOS)_ ([#570](https://github.com/infobloxopen/terraform-provider-infoblox/pull/570))
- `infoblox_servicerestart_group` : Manage Grid Service Restart Groups and retrieve existing Grid Service Restart Groups. _(NIOS)_ ([#596](https://github.com/infobloxopen/terraform-provider-infoblox/pull/596))
- `infoblox_upgrade_group` : Manage Grid Upgrade Groups and retrieve existing Grid Upgrade Groups. _(NIOS)_ ([#585](https://github.com/infobloxopen/terraform-provider-infoblox/pull/585))

#### Infra

- `infoblox_infra_host` : Manage Infrastructure Hosts and retrieve existing Infrastructure Hosts. _(UDDI)_ ([#656](https://github.com/infobloxopen/terraform-provider-infoblox/pull/656))
- `infoblox_infra_service` : Manage Infrastructure Services and retrieve existing Infrastructure Services. _(UDDI)_ ([#658](https://github.com/infobloxopen/terraform-provider-infoblox/pull/658))
- `infoblox_join_token` : Manage Join Tokens and retrieve existing Join Tokens. _(UDDI)_ ([#689](https://github.com/infobloxopen/terraform-provider-infoblox/pull/689))

#### IPAM

- `infoblox_address` : Manage IPAM Addresses and retrieve existing IPAM Addresses. _(UDDI)_ ([#519](https://github.com/infobloxopen/terraform-provider-infoblox/pull/519))
- `infoblox_bulk_hostname_template` : Manage IPAM Bulk Hostname Templates and retrieve existing IPAM Bulk Hostname Templates. _(NIOS)_ ([#612](https://github.com/infobloxopen/terraform-provider-infoblox/pull/612))
- `infoblox_ipam_host` : Manage IPAM Hosts and retrieve existing IPAM Hosts. _(UDDI)_ ([#672](https://github.com/infobloxopen/terraform-provider-infoblox/pull/672))
- `infoblox_ipv6_network` : Manage IPAM IPv6 Networks and retrieve existing IPAM IPv6 Networks. _(NIOS, UDDI)_ ([#560](https://github.com/infobloxopen/terraform-provider-infoblox/pull/560))
- `infoblox_ipv6_network_container` : Manage IPAM IPv6 Network Containers and retrieve existing IPAM IPv6 Network Containers. _(NIOS, UDDI)_ ([#561](https://github.com/infobloxopen/terraform-provider-infoblox/pull/561))
- `infoblox_network` : Manage IPAM Networks and retrieve existing IPAM Networks. _(NIOS, UDDI)_ ([#545](https://github.com/infobloxopen/terraform-provider-infoblox/pull/545))
- `infoblox_network_container` : Manage IPAM Network Containers and retrieve existing IPAM Network Containers. _(NIOS, UDDI)_ ([#543](https://github.com/infobloxopen/terraform-provider-infoblox/pull/543))
- `infoblox_network_view` : Manage IPAM Network Views and retrieve existing IPAM Network Views. _(NIOS, UDDI)_ ([#557](https://github.com/infobloxopen/terraform-provider-infoblox/pull/557))
- `infoblox_superhost` : Manage IPAM Super Hosts and retrieve existing IPAM Super Hosts. _(NIOS)_ ([#598](https://github.com/infobloxopen/terraform-provider-infoblox/pull/598))
- `infoblox_vlan` : Manage IPAM VLANs and retrieve existing IPAM VLANs. _(NIOS)_ ([#649](https://github.com/infobloxopen/terraform-provider-infoblox/pull/649))
- `infoblox_vlan_range` : Manage IPAM VLAN Ranges and retrieve existing IPAM VLAN Ranges. _(NIOS)_ ([#661](https://github.com/infobloxopen/terraform-provider-infoblox/pull/661))
- `infoblox_vlan_view` : Manage IPAM VLAN Views and retrieve existing IPAM VLAN Views. _(NIOS)_ ([#649](https://github.com/infobloxopen/terraform-provider-infoblox/pull/649))

#### IPAM Federation

- `infoblox_federated_realm` : Manage Federated Realms and retrieve existing Federated Realms. _(UDDI)_ ([#590](https://github.com/infobloxopen/terraform-provider-infoblox/pull/590))

#### Keys

- `infoblox_tsig_key` : Manage TSIG Keys and retrieve existing TSIG Keys. _(UDDI)_ ([#589](https://github.com/infobloxopen/terraform-provider-infoblox/pull/589))

#### Miscellaneous

- `infoblox_bfd_template` : Manage BFD Templates and retrieve existing BFD Templates. _(NIOS)_ ([#564](https://github.com/infobloxopen/terraform-provider-infoblox/pull/564))
- `infoblox_ruleset` : Manage Rule Sets and retrieve existing Rule Sets. _(NIOS)_ ([#553](https://github.com/infobloxopen/terraform-provider-infoblox/pull/553))

#### Notification

- `infoblox_notification_rest_endpoint` : Manage Notification REST Endpoints and retrieve existing Notification REST Endpoints. _(NIOS)_ ([#629](https://github.com/infobloxopen/terraform-provider-infoblox/pull/629))
- `infoblox_notification_rule` : Manage Notification Rules and retrieve existing Notification Rules. _(NIOS)_ ([#688](https://github.com/infobloxopen/terraform-provider-infoblox/pull/688))

#### Redirect

- `infoblox_custom_redirect` : Manage Custom Redirects and retrieve existing Custom Redirects. _(UDDI)_ ([#657](https://github.com/infobloxopen/terraform-provider-infoblox/pull/657))

#### RIR

- `infoblox_rir_organization` : Manage RIR Organizations and retrieve existing RIR Organizations. _(NIOS)_ ([#673](https://github.com/infobloxopen/terraform-provider-infoblox/pull/673))

#### RPZ

- `infoblox_record_rpz_a` : Manage RPZ A Records and retrieve existing RPZ A Records. _(NIOS)_ ([#624](https://github.com/infobloxopen/terraform-provider-infoblox/pull/624))
- `infoblox_record_rpz_a_ipaddress` : Manage RPZ A IP Address Records and retrieve existing RPZ A IP Address Records. _(NIOS)_ ([#636](https://github.com/infobloxopen/terraform-provider-infoblox/pull/636))
- `infoblox_record_rpz_aaaa` : Manage RPZ AAAA Records and retrieve existing RPZ AAAA Records. _(NIOS)_ ([#632](https://github.com/infobloxopen/terraform-provider-infoblox/pull/632))
- `infoblox_record_rpz_aaaa_ipaddress` : Manage RPZ AAAA IP Address Records and retrieve existing RPZ AAAA IP Address Records. _(NIOS)_ ([#562](https://github.com/infobloxopen/terraform-provider-infoblox/pull/562))
- `infoblox_record_rpz_cname` : Manage RPZ CNAME Records and retrieve existing RPZ CNAME Records. _(NIOS)_ ([#617](https://github.com/infobloxopen/terraform-provider-infoblox/pull/617))
- `infoblox_record_rpz_cname_clientipaddress` : Manage RPZ CNAME Client IP Address Records and retrieve existing RPZ CNAME Client IP Address Records. _(NIOS)_ ([#633](https://github.com/infobloxopen/terraform-provider-infoblox/pull/633))
- `infoblox_record_rpz_cname_clientipaddressdn` : Manage RPZ CNAME Client IP Address DN Records and retrieve existing RPZ CNAME Client IP Address DN Records. _(NIOS)_ ([#578](https://github.com/infobloxopen/terraform-provider-infoblox/pull/578))
- `infoblox_record_rpz_cname_ipaddress` : Manage RPZ CNAME IP Address Records and retrieve existing RPZ CNAME IP Address Records. _(NIOS)_ ([#622](https://github.com/infobloxopen/terraform-provider-infoblox/pull/622))
- `infoblox_record_rpz_cname_ipaddressdn` : Manage RPZ CNAME IP Address DN Records and retrieve existing RPZ CNAME IP Address DN Records. _(NIOS)_ ([#635](https://github.com/infobloxopen/terraform-provider-infoblox/pull/635))
- `infoblox_record_rpz_naptr` : Manage RPZ NAPTR Records and retrieve existing RPZ NAPTR Records. _(NIOS)_ ([#574](https://github.com/infobloxopen/terraform-provider-infoblox/pull/574))
- `infoblox_record_rpz_ptr` : Manage RPZ PTR Records and retrieve existing RPZ PTR Records. _(NIOS)_ ([#637](https://github.com/infobloxopen/terraform-provider-infoblox/pull/637))
- `infoblox_record_rpz_txt` : Manage RPZ TXT Records and retrieve existing RPZ TXT Records. _(NIOS)_ ([#577](https://github.com/infobloxopen/terraform-provider-infoblox/pull/577))

#### Security

- `infoblox_admin_user` : Manage Admin Users and retrieve existing Admin Users. _(NIOS)_ ([#676](https://github.com/infobloxopen/terraform-provider-infoblox/pull/676))

### Features

- Every object has a list resource, so you can find existing objects with `terraform query` (Terraform v1.14.0 or later).
- Added retries with exponential backoff, limited by the provider-level `operation_timeout` setting. ([#537](https://github.com/infobloxopen/terraform-provider-infoblox/pull/537))
- Added support for managing a NIOS Grid through the Infoblox Portal with `uddi.enable_nios_passthru`. ([#555](https://github.com/infobloxopen/terraform-provider-infoblox/pull/555))
- Added the `proxy_search` provider setting and the `nios.proxy_url` setting for NIOS requests. ([#684](https://github.com/infobloxopen/terraform-provider-infoblox/pull/684))
- Added `ssl_verify`, `ca_cert_file`, and `ca_cert_pem` to configure TLS certificate verification for NIOS. ([#700](https://github.com/infobloxopen/terraform-provider-infoblox/pull/700))

### Maintenance

- Upgraded the Go version to 1.26.8. ([#713](https://github.com/infobloxopen/terraform-provider-infoblox/pull/713))

## [v2.13.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.13.0) (2026-04-02)

- Upgraded Go version to v1.25.8 and refreshed dependencies for better compatibility and security.

## [v2.12.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.12.0) (2026-03-10)

- Data source updated to support VLAN.

## [v2.11.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.11.0) (2025-10-20)

- Improved error handling in data sources by treating "not found" responses as valid empty results, enhancing stability.
- Upgraded Go to v1.24.7 and refreshed dependencies for better compatibility and security.
- Minor refactoring to simplify error handling and remove redundant checks.

## [v2.10.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.10.0) (2025-4-25)
- New Resources:
  - infoblox_alias_record
  - infoblox_ns_record
  - infoblox_ipv4_shared_network
  - infoblox_ipv4_fixed_address
  - infoblox_ipv4_range
  - infoblox_ipv4_range_template
- New Datasources:
  - infoblox_alias_record
  - infoblox_ns_record
  - infoblox_ipv4_shared_network
  - infoblox_ipv4_fixed_address
  - infoblox_ipv4_range
  - infoblox_ipv4_range_template

## [v2.9.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.9.0) (2025-1-31)
- Upgraded the base WAPI version to v2.12.3.
- New Resources:
  - infoblox_dtc_lbdn
  - infoblox_dtc_pool
  - infoblox_dtc_server
- New Datasources:
  - infoblox_dtc_lbdn
  - infoblox_dtc_pool
  - infoblox_dtc_server

## [v2.8.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.8.0) (2024-10-30)
- New Feature: 
  - Support for Creating Objects with the Next Available IP Address and Next Available Network Obtained Using the Extensible Attribute Filter.
  - Add support for aliases and disable flag in Host Records.
- New Resources:
  - infoblox_zone_delegated
- New Datasources:
  - infoblox_zone_delegated

## [v2.7.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.7.0) (2024-06-20)
- New Feature: Added support for Import block.
- New Resources:
  - infoblox_zone_forward
- New Datasources:
  - infoblox_zone_forward
  - infoblox_ipv6_network
  - infoblox_ipv6_network_container
  - infoblox_host_record

## [v2.6.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.6.0) (2024-04-17)
- New Feature: Ability to manage drift through "Terraform Internal ID" Extensible Attributes in resources
- Bugfixes
  - Fixed Host record import with empty MAC or DUID.
  
## [v2.5.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.5.0) (2023-10-30)
- Resources are reworked aganist the changes from auto generated objects in go-client
- New Feature: ability to search through Extensible Attributes in datasources
- Additionally, added Multi Value Extensible Attributes search support
- EA Inheritance issue fixed, where inherited EAs in NIOS were getting deleted for second apply
- Datasources are reworked to use `filters`, for fetching matching objects, refer to [Terraform Docs](https://github.com/infobloxopen/terraform-provider-infoblox/blob/master/docs/index.md)
- New Resources:
  - infoblox_dns_view
  - infoblox_zone_auth
- New Datasources:
  - infoblox_dns_view
  - infoblox_zone_auth

## [v2.4.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.4.1) (2023-06-20)
- A/AAAA Record resources reworked:
  - removed limitation on updating 'cidr' field
- Bugfixes

## [v2.4.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.4.0) (2023-05-29)
- IPV4/IPV6 Network Container resources reworked:
  - 'parent_cidr' and 'allocate_prefix_len' are added for dynamic allocation
  - both the resources now support the dynamic allocation determined by 'parent_cidr'
  - added examples for dynamic allocation in each IPV4 and IPV6 resources
- Bugfixes

## [v2.3.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.3.0) (2023-04-13)
- Minimal required Go-lang version is bumped up
- infoblox_ptr_record resource's behaviour changes (see the documentation changes for the details)
- 'dns_view' and 'network_view' fields are now optional for all the `resources and data sources
- New resources:
  - infoblox_mx_record
  - infoblox_txt_record
  - infoblox_srv_record
- New data sources:
  - infoblox_mx_record
  - infoblox_txt_record
  - infoblox_srv_record
- Bugfixes

## [v2.2.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.2.0) (2022-12-07)
- New feature: ability to import resources from existing NIOS objects
- New Data sources:
  - infoblox_aaaa_record
  - infoblox_ptr_record
  - infoblox_network_view
  - infoblox_ipv4_network_container
- Allocation/Association resources have been reworked
- Examples are reorganized
- Numerous bugfixes

## [v2.1.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.1.0) (2022-02-10)

- Moved to Terraform Plugin SDK v2
- Allocation/Association resources reworked:
  - new resources added: infoblox_ip_allocation and infoblox_ip_association;
    both IPv4 and IPv6 addresses may be allocated within a single resource in one go.
  - infoblox_ipv4_allocation, infoblox_ipv6_allocation, infoblox_ipv4_association and infoblox_ipv6_association
    are deprecated and unsupported from now on.
- Improvements in infoblox_ipv4_network and infoblox_ipv6_network resources: IP address reservation reworked.
- Numerous bugfixes

## [v2.0.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v2.0.1) (2021-07-27)

**List of changes since 1.0.6 release:**

- Existing resources have been changed:
  - Extensible attributes and comments are introduced
  - Some fields are removed, use Extensible Attributes instead.
  - IPv6 version of resources added
- New resources:
  - Network container (infoblox_ipv4_network_container, infoblox_ipv6_network_container)
  - Network view (infoblox_network_view)
  - AAAA-record (infoblox_aaa_record)
- Data sources: extensible attributes and comments are introduced
  - infoblox_a_record
  - infoblox_cname_record
  - infoblox_ipv4_network
- New features:
  - IP address auto-allocation feature
  - Ability to update a resource, in addition to 'create', 'read' and 'delete' operations.
  - All the resources have 'comment' and 'ext_attrs' fields.
  - DNS-related fields have 'ttl' field. TTL=0 means TTL inherited from the parent zone.
- Examples of how to integrate with different cloud environments

## [v1.1.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.1.1) (2021-04-23)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.1.0...v1.1.1)

**Implemented enhancements:**

- Resources are not truly destroyed upon destroy [\#57](https://github.com/infobloxopen/terraform-provider-infoblox/issues/57)
- Can't dynamically request a network [\#27](https://github.com/infobloxopen/terraform-provider-infoblox/issues/27)

**Closed issues:**

- Feature request: ability to change extended attributes [\#77](https://github.com/infobloxopen/terraform-provider-infoblox/issues/77)
- Commit changes? [\#69](https://github.com/infobloxopen/terraform-provider-infoblox/issues/69)
- DHCP Host records won't store in state [\#68](https://github.com/infobloxopen/terraform-provider-infoblox/issues/68)
- Terraform Provider Development Program - Second Review [\#66](https://github.com/infobloxopen/terraform-provider-infoblox/issues/66)
- Passing credentials without using environment variables [\#65](https://github.com/infobloxopen/terraform-provider-infoblox/issues/65)
- inconsistent vendoring [\#59](https://github.com/infobloxopen/terraform-provider-infoblox/issues/59)
- make build fails [\#34](https://github.com/infobloxopen/terraform-provider-infoblox/issues/34)
- Error accessing infoblox thru https proxy [\#28](https://github.com/infobloxopen/terraform-provider-infoblox/issues/28)

**Merged pull requests:**

- Remove old release workflow [\#111](https://github.com/infobloxopen/terraform-provider-infoblox/pull/111) ([cgroschupp](https://github.com/cgroschupp))
- Display travis build status of master and develop branch [\#109](https://github.com/infobloxopen/terraform-provider-infoblox/pull/109) ([somashekhar](https://github.com/somashekhar))
- Add provider to terraform registry [\#107](https://github.com/infobloxopen/terraform-provider-infoblox/pull/107) ([cgroschupp](https://github.com/cgroschupp))
- Revert "doc change" [\#91](https://github.com/infobloxopen/terraform-provider-infoblox/pull/91) ([somashekhar](https://github.com/somashekhar))
- doc change [\#90](https://github.com/infobloxopen/terraform-provider-infoblox/pull/90) ([somashekhar](https://github.com/somashekhar))
- doc change [\#85](https://github.com/infobloxopen/terraform-provider-infoblox/pull/85) ([somashekhar](https://github.com/somashekhar))
- Example NIOS and AWS/VMWare/Azure tf files to support Next Available Network and DNS/DHCP Records support for AWS. [\#84](https://github.com/infobloxopen/terraform-provider-infoblox/pull/84) ([somashekhar](https://github.com/somashekhar))
- Check that parent exists before network allocation. [\#83](https://github.com/infobloxopen/terraform-provider-infoblox/pull/83) ([AliaksandrDziarkach](https://github.com/AliaksandrDziarkach))
- Migration from terraform providers to our own repository [\#78](https://github.com/infobloxopen/terraform-provider-infoblox/pull/78) ([AvRajath](https://github.com/AvRajath))

## [v1.0.6](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.6) (2021-04-23)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.1.1...v1.0.6)

## [v1.0.5](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.5) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.4...v1.0.5)

## [v1.0.4](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.4) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.3...v1.0.4)

## [v1.0.3](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.3) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.2...v1.0.3)

**Closed issues:**

- Are you able to make releases through CI? [\#33](https://github.com/infobloxopen/terraform-provider-infoblox/issues/33)

## [v1.0.2](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.2) (2020-05-14)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.1...v1.0.2)

**Closed issues:**

- a\_record generated with duplicate domain name [\#71](https://github.com/infobloxopen/terraform-provider-infoblox/issues/71)
- Error creating A Record from network block [\#62](https://github.com/infobloxopen/terraform-provider-infoblox/issues/62)
- Too many errors [\#61](https://github.com/infobloxopen/terraform-provider-infoblox/issues/61)
- Resource: infoblox\_cname\_record by default appending zone information to alias record [\#60](https://github.com/infobloxopen/terraform-provider-infoblox/issues/60)
- Improved documentation [\#55](https://github.com/infobloxopen/terraform-provider-infoblox/issues/55)
- terraform doesn't change a HostRecord name when the vm\_name is changed in the block, only the EA vm\_name is changed, not the actual name [\#54](https://github.com/infobloxopen/terraform-provider-infoblox/issues/54)
- Add HostRecord function [\#53](https://github.com/infobloxopen/terraform-provider-infoblox/issues/53)
- terraform doesn't like it when a managed IP gets removed from infoblox by hand [\#52](https://github.com/infobloxopen/terraform-provider-infoblox/issues/52)
- infoblox\_ip\_association causing crash [\#49](https://github.com/infobloxopen/terraform-provider-infoblox/issues/49)
- infoblox\_ip\_allocation fails - doesn't know what to do with vm\_name field [\#48](https://github.com/infobloxopen/terraform-provider-infoblox/issues/48)
- Terraform Provider Development Program - Review [\#45](https://github.com/infobloxopen/terraform-provider-infoblox/issues/45)
- "stock" go make build fails [\#43](https://github.com/infobloxopen/terraform-provider-infoblox/issues/43)
- API error [\#32](https://github.com/infobloxopen/terraform-provider-infoblox/issues/32)
- Official terraform provider status [\#30](https://github.com/infobloxopen/terraform-provider-infoblox/issues/30)
- Does this provider work with Terraform 0.12.1 [\#29](https://github.com/infobloxopen/terraform-provider-infoblox/issues/29)
- Tenant ID Parameter? [\#21](https://github.com/infobloxopen/terraform-provider-infoblox/issues/21)
- Error using your Terraform provider to access Infoblox api [\#20](https://github.com/infobloxopen/terraform-provider-infoblox/issues/20)

**Merged pull requests:**

- New datasource infoblox network [\#74](https://github.com/infobloxopen/terraform-provider-infoblox/pull/74) ([pearcec](https://github.com/pearcec))
- Delete Network View fix issue \#57 [\#73](https://github.com/infobloxopen/terraform-provider-infoblox/pull/73) ([pearcec](https://github.com/pearcec))
- Leverage Github Actions for Building, Tagging, and Releasing the provider [\#70](https://github.com/infobloxopen/terraform-provider-infoblox/pull/70) ([NickLarsenNZ](https://github.com/NickLarsenNZ))
- Modify CNAME resource create function to remove zone [\#64](https://github.com/infobloxopen/terraform-provider-infoblox/pull/64) ([elangoganesan](https://github.com/elangoganesan))
- Updated Readme for use case without CNA license [\#63](https://github.com/infobloxopen/terraform-provider-infoblox/pull/63) ([AvRajath](https://github.com/AvRajath))
- Terraform review program [\#47](https://github.com/infobloxopen/terraform-provider-infoblox/pull/47) ([saiprasannasastry](https://github.com/saiprasannasastry))
- merge conflicts [\#46](https://github.com/infobloxopen/terraform-provider-infoblox/pull/46) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Moving to go mod and bumping terraform version [\#42](https://github.com/infobloxopen/terraform-provider-infoblox/pull/42) ([saiprasannasastry](https://github.com/saiprasannasastry))

## [v1.0.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.1) (2019-09-05)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.0...v1.0.1)

**Merged pull requests:**

- Fix build [\#39](https://github.com/infobloxopen/terraform-provider-infoblox/pull/39) ([mikecook](https://github.com/mikecook))
- Dsbrng25b remove license check [\#38](https://github.com/infobloxopen/terraform-provider-infoblox/pull/38) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Addressing some internal QA bugs [\#37](https://github.com/infobloxopen/terraform-provider-infoblox/pull/37) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Remove license check entirely [\#36](https://github.com/infobloxopen/terraform-provider-infoblox/pull/36) ([dvob](https://github.com/dvob))

## [v1.0.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.0) (2019-08-09)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/ec354d7410947945fe01f876269ded480f18029d...v1.0.0)

**Implemented enhancements:**

- Support for CNAME or A Records [\#22](https://github.com/infobloxopen/terraform-provider-infoblox/issues/22)

**Closed issues:**

- Error on plan [\#26](https://github.com/infobloxopen/terraform-provider-infoblox/issues/26)

**Merged pull requests:**

- CHANGELOG.md updated to v1.0.0 [\#35](https://github.com/infobloxopen/terraform-provider-infoblox/pull/35) ([jkraj](https://github.com/jkraj))
- Minor readability changes to README.md [\#31](https://github.com/infobloxopen/terraform-provider-infoblox/pull/31) ([scottsuch](https://github.com/scottsuch))
- This commit contains entire DNS changes [\#25](https://github.com/infobloxopen/terraform-provider-infoblox/pull/25) ([saiprasannasastry](https://github.com/saiprasannasastry))
- This commit contains patch to rebase issue [\#19](https://github.com/infobloxopen/terraform-provider-infoblox/pull/19) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Update README.md [\#18](https://github.com/infobloxopen/terraform-provider-infoblox/pull/18) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Example [\#17](https://github.com/infobloxopen/terraform-provider-infoblox/pull/17) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Update README.md [\#16](https://github.com/infobloxopen/terraform-provider-infoblox/pull/16) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Added block of code to reserve n number of IP's [\#14](https://github.com/infobloxopen/terraform-provider-infoblox/pull/14) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Added an example tf for Vmware Vshpere [\#12](https://github.com/infobloxopen/terraform-provider-infoblox/pull/12) ([saiprasannasastry](https://github.com/saiprasannasastry))
- UT's for Infoblox-Provider [\#11](https://github.com/infobloxopen/terraform-provider-infoblox/pull/11) ([saiprasannasastry](https://github.com/saiprasannasastry))
- This commit contains a new resource IPAssociation [\#10](https://github.com/infobloxopen/terraform-provider-infoblox/pull/10) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Validation of Cloud License [\#8](https://github.com/infobloxopen/terraform-provider-infoblox/pull/8) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Validation of gateway ip existence [\#5](https://github.com/infobloxopen/terraform-provider-infoblox/pull/5) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Modified makefile to build and gofmtcheck [\#3](https://github.com/infobloxopen/terraform-provider-infoblox/pull/3) ([jkraj](https://github.com/jkraj))
- Support for network and ip allocation [\#1](https://github.com/infobloxopen/terraform-provider-infoblox/pull/1) ([saiprasannasastry](https://github.com/saiprasannasastry))



\* *This Changelog was automatically generated       by [github_changelog_generator]      (https://github.com/github-changelog-generator/github-changelog-generator)*
