---
page_title: "Resources and Data Sources"
subcategory: "Guides"
description: |-
  Lists all resources and data sources and the backends (NIOS, UDDI) each one supports.
---

# Resources and Data Sources

The tables below list all available resources and data sources, and the backends each one supports.

- **NIOS**: works with an Infoblox NIOS Grid (WAPI).
- **UDDI**: works with Infoblox Universal DDI (Infoblox Portal).
- **NIOS, UDDI**: works with both. The backend in use comes from the provider configuration.

Every object that has a data source also has a list resource of the same name, for use with `terraform query`.

### ACL

| Name                | Resource Description               | Data Source Description                                         | Supported Backends |
|---------------------|------------------------------------|-----------------------------------------------------------------|--------------------|
| `infoblox_namedacl` | Manages Named Access Control Lists | Retrieves information about existing Named Access Control Lists | NIOS, UDDI         |

### ANYCAST

| Name                      | Resource Description           | Data Source Description                                     | Supported Backends |
|---------------------------|--------------------------------|-------------------------------------------------------------|--------------------|
| `infoblox_anycast_config` | Manages Anycast Configurations | Retrieves information about existing Anycast Configurations | UDDI               |
| `infoblox_anycast_host`   | Manages Anycast Hosts          |                                                             | UDDI               |

### CLOUD

| Name                | Resource Description | Data Source Description                        | Supported Backends |
|---------------------|----------------------|------------------------------------------------|--------------------|
| `infoblox_aws_user` | Manages AWS Users    | Retrieves information about existing AWS Users | NIOS               |

### CLOUD DISCOVERY

| Name                                | Resource Description              | Data Source Description                                        | Supported Backends |
|-------------------------------------|-----------------------------------|----------------------------------------------------------------|--------------------|
| `infoblox_cloud_discovery_provider` | Manages Cloud Discovery Providers | Retrieves information about existing Cloud Discovery Providers | UDDI               |

### DHCP

| Name                                   | Resource Description                      | Data Source Description                                                | Supported Backends |
|----------------------------------------|-------------------------------------------|------------------------------------------------------------------------|--------------------|
| `infoblox_dhcp_optiondefinition`       | Manages DHCP Option Definitions           | Retrieves information about existing DHCP Option Definitions           | NIOS, UDDI         |
| `infoblox_dhcp_optionspace`            | Manages DHCP Option Spaces                | Retrieves information about existing DHCP Option Spaces                | NIOS, UDDI         |
| `infoblox_filteroption`                | Manages DHCP Filter Options               | Retrieves information about existing DHCP Filter Options               | NIOS, UDDI         |
| `infoblox_fixed_address`               | Manages DHCP Fixed Addresses (IPv4)       | Retrieves information about existing DHCP Fixed Addresses (IPv4)       | NIOS, UDDI         |
| `infoblox_ha_group`                    | Manages DHCP HA Groups                    | Retrieves information about existing DHCP HA Groups                    | UDDI               |
| `infoblox_hardware_filter`             | Manages DHCP Hardware Filters             | Retrieves information about existing DHCP Hardware Filters             | UDDI               |
| `infoblox_ipv6_dhcp_optiondefinition`  | Manages DHCP IPv6 Option Definitions      | Retrieves information about existing DHCP IPv6 Option Definitions      | NIOS, UDDI         |
| `infoblox_ipv6_dhcp_optionspace`       | Manages DHCP IPv6 Option Spaces           | Retrieves information about existing DHCP IPv6 Option Spaces           | NIOS, UDDI         |
| `infoblox_ipv6_filteroption`           | Manages DHCP IPv6 Filter Options          | Retrieves information about existing DHCP IPv6 Filter Options          | NIOS, UDDI         |
| `infoblox_ipv6_fixed_address`          | Manages DHCP IPv6 Fixed Addresses         | Retrieves information about existing DHCP IPv6 Fixed Addresses         | NIOS, UDDI         |
| `infoblox_ipv6_fixed_address_template` | Manages DHCP IPv6 Fixed Address Templates | Retrieves information about existing DHCP IPv6 Fixed Address Templates | NIOS               |
| `infoblox_ipv6_range_template`         | Manages DHCP IPv6 Range Templates         | Retrieves information about existing DHCP IPv6 Range Templates         | NIOS               |
| `infoblox_ipv6_sharednetwork`          | Manages DHCP IPv6 Shared Networks         | Retrieves information about existing DHCP IPv6 Shared Networks         | NIOS               |
| `infoblox_option_group`                | Manages DHCP Option Groups                | Retrieves information about existing DHCP Option Groups                | UDDI               |
| `infoblox_range`                       | Manages DHCP Ranges (IPv4)                | Retrieves information about existing DHCP Ranges (IPv4)                | NIOS, UDDI         |
| `infoblox_range_template`              | Manages DHCP Range Templates (IPv4)       | Retrieves information about existing DHCP Range Templates (IPv4)       | NIOS               |
| `infoblox_sharednetwork`               | Manages DHCP Shared Networks (IPv4)       | Retrieves information about existing DHCP Shared Networks (IPv4)       | NIOS               |

### DISCOVERY

| Name                                  | Resource Description                | Data Source Description                                          | Supported Backends |
|---------------------------------------|-------------------------------------|------------------------------------------------------------------|--------------------|
| `infoblox_discovery_credential_group` | Manages Discovery Credential Groups | Retrieves information about existing Discovery Credential Groups | NIOS               |

### DNS

| Name                                 | Resource Description                        | Data Source Description                                                | Supported Backends |
|--------------------------------------|---------------------------------------------|------------------------------------------------------------------------|--------------------|
| `infoblox_auth_nsg`                  | Manages Authoritative DNS Server Groups     | Retrieves information about existing Authoritative DNS Server Groups   | UDDI               |
| `infoblox_dns_host`                  | Configures existing DNS Hosts               | Retrieves information about existing DNS Hosts                         | UDDI               |
| `infoblox_dns_server`                | Manages DNS Servers                         | Retrieves information about existing DNS Servers                       | UDDI               |
| `infoblox_forward_nsg`               | Manages Forward DNS Server Groups           | Retrieves information about existing Forward DNS Server Groups         | UDDI               |
| `infoblox_ip_association`            | Manages an IP Association for a Host Record |                                                                        | NIOS               |
| `infoblox_nsgroup`                   | Manages DNS NS Groups                       | Retrieves information about existing DNS NS Groups                     | NIOS               |
| `infoblox_nsgroup_delegation`        | Manages DNS NS Group Delegations            | Retrieves information about existing DNS NS Group Delegations          | NIOS               |
| `infoblox_nsgroup_forwardingmember`  | Manages DNS NS Group Forwarding Members     | Retrieves information about existing DNS NS Group Forwarding Members   | NIOS               |
| `infoblox_nsgroup_forwardstubserver` | Manages DNS NS Group Forward Stub Servers   | Retrieves information about existing DNS NS Group Forward Stub Servers | NIOS               |
| `infoblox_nsgroup_stubmember`        | Manages DNS NS Group Stub Members           | Retrieves information about existing DNS NS Group Stub Members         | NIOS               |
| `infoblox_record_a`                  | Manages DNS A Records                       | Retrieves information about existing DNS A Records                     | NIOS, UDDI         |
| `infoblox_record_aaaa`               | Manages DNS AAAA Records                    | Retrieves information about existing DNS AAAA Records                  | NIOS, UDDI         |
| `infoblox_record_alias`              | Manages DNS ALIAS Records                   | Retrieves information about existing DNS ALIAS Records                 | NIOS               |
| `infoblox_record_caa`                | Manages DNS CAA Records                     | Retrieves information about existing DNS CAA Records                   | NIOS, UDDI         |
| `infoblox_record_cname`              | Manages DNS CNAME Records                   | Retrieves information about existing DNS CNAME Records                 | NIOS, UDDI         |
| `infoblox_record_dname`              | Manages DNS DNAME Records                   | Retrieves information about existing DNS DNAME Records                 | NIOS, UDDI         |
| `infoblox_record_host`               | Manages DNS Host Records                    | Retrieves information about existing DNS Host Records                  | NIOS               |
| `infoblox_record_https`              | Manages DNS HTTPS Records                   | Retrieves information about existing DNS HTTPS Records                 | UDDI               |
| `infoblox_record_mx`                 | Manages DNS MX Records                      | Retrieves information about existing DNS MX Records                    | NIOS, UDDI         |
| `infoblox_record_naptr`              | Manages DNS NAPTR Records                   | Retrieves information about existing DNS NAPTR Records                 | NIOS, UDDI         |
| `infoblox_record_ns`                 | Manages DNS NS Records                      | Retrieves information about existing DNS NS Records                    | NIOS, UDDI         |
| `infoblox_record_ptr`                | Manages DNS PTR Records                     | Retrieves information about existing DNS PTR Records                   | NIOS, UDDI         |
| `infoblox_record_srv`                | Manages DNS SRV Records                     | Retrieves information about existing DNS SRV Records                   | NIOS, UDDI         |
| `infoblox_record_svcb`               | Manages DNS SVCB Records                    | Retrieves information about existing DNS SVCB Records                  | UDDI               |
| `infoblox_record_txt`                | Manages DNS TXT Records                     | Retrieves information about existing DNS TXT Records                   | NIOS, UDDI         |
| `infoblox_sharedrecord_a`            | Manages DNS Shared A Records                | Retrieves information about existing DNS Shared A Records              | NIOS               |
| `infoblox_sharedrecord_aaaa`         | Manages DNS Shared AAAA Records             | Retrieves information about existing DNS Shared AAAA Records           | NIOS               |
| `infoblox_sharedrecord_cname`        | Manages DNS Shared CNAME Records            | Retrieves information about existing DNS Shared CNAME Records          | NIOS               |
| `infoblox_sharedrecord_group`        | Manages DNS Shared Record Groups            | Retrieves information about existing DNS Shared Record Groups          | NIOS               |
| `infoblox_sharedrecord_mx`           | Manages DNS Shared MX Records               | Retrieves information about existing DNS Shared MX Records             | NIOS               |
| `infoblox_sharedrecord_srv`          | Manages DNS Shared SRV Records              | Retrieves information about existing DNS Shared SRV Records            | NIOS               |
| `infoblox_sharedrecord_txt`          | Manages DNS Shared TXT Records              | Retrieves information about existing DNS Shared TXT Records            | NIOS               |
| `infoblox_view`                      | Manages DNS Views                           | Retrieves information about existing DNS Views                         | NIOS, UDDI         |
| `infoblox_zone_auth`                 | Manages Authoritative DNS Zones             | Retrieves information about existing Authoritative DNS Zones           | NIOS, UDDI         |
| `infoblox_zone_delegated`            | Manages Delegated DNS Zones                 | Retrieves information about existing Delegated DNS Zones               | NIOS, UDDI         |
| `infoblox_zone_forward`              | Manages Forwarding DNS Zones                | Retrieves information about existing Forwarding DNS Zones              | NIOS, UDDI         |
| `infoblox_zone_rp`                   | Manages DNS Response Policy Zones           | Retrieves information about existing DNS Response Policy Zones         | NIOS               |
| `infoblox_zone_stub`                 | Manages DNS Stub Zones                      | Retrieves information about existing DNS Stub Zones                    | NIOS               |

### DTC

| Name                        | Resource Description      | Data Source Description                                | Supported Backends |
|-----------------------------|---------------------------|--------------------------------------------------------|--------------------|
| `infoblox_dtc_lbdn`         | Manages DTC LBDNs         | Retrieves information about existing DTC LBDNs         | NIOS, UDDI         |
| `infoblox_dtc_monitor_http` | Manages DTC HTTP Monitors | Retrieves information about existing DTC HTTP Monitors | NIOS, UDDI         |
| `infoblox_dtc_monitor_icmp` | Manages DTC ICMP Monitors | Retrieves information about existing DTC ICMP Monitors | NIOS, UDDI         |
| `infoblox_dtc_monitor_pdp`  | Manages DTC PDP Monitors  | Retrieves information about existing DTC PDP Monitors  | NIOS, UDDI         |
| `infoblox_dtc_monitor_snmp` | Manages DTC SNMP Monitors | Retrieves information about existing DTC SNMP Monitors | NIOS, UDDI         |
| `infoblox_dtc_monitor_tcp`  | Manages DTC TCP Monitors  | Retrieves information about existing DTC TCP Monitors  | NIOS, UDDI         |
| `infoblox_dtc_pool`         | Manages DTC Pools         | Retrieves information about existing DTC Pools         | NIOS, UDDI         |
| `infoblox_dtc_server`       | Manages DTC Servers       | Retrieves information about existing DTC Servers       | NIOS, UDDI         |
| `infoblox_dtc_topology`     | Manages DTC Topologies    | Retrieves information about existing DTC Topologies    | NIOS, UDDI         |

### FW

| Name                            | Resource Description                      | Data Source Description                                                | Supported Backends |
|---------------------------------|-------------------------------------------|------------------------------------------------------------------------|--------------------|
| `infoblox_access_code`          | Manages Access Codes                      | Retrieves information about existing Access Codes                      | UDDI               |
| `infoblox_application_filter`   | Manages Application Filters               | Retrieves information about existing Application Filters               | UDDI               |
| `infoblox_category_filter`      | Manages Category Filters                  | Retrieves information about existing Category Filters                  | UDDI               |
| `infoblox_internal_domain_list` | Manages Internal Domain Lists             | Retrieves information about existing Internal Domain Lists             | UDDI               |
| `infoblox_named_list`           | Manages Named Lists (Custom Lists)        | Retrieves information about existing Named Lists (Custom Lists)        | UDDI               |
| `infoblox_network_list`         | Manages Network Lists (External Networks) | Retrieves information about existing Network Lists (External Networks) | UDDI               |
| `infoblox_security_policy`      | Manages Security Policies                 | Retrieves information about existing Security Policies                 | UDDI               |

### GRID

| Name                                | Resource Description                      | Data Source Description                                               | Supported Backends |
|-------------------------------------|-------------------------------------------|-----------------------------------------------------------------------|--------------------|
| `infoblox_distribution_schedule`    | Configures the Grid Distribution Schedule | Retrieves information about the current Grid Distribution Schedule    | NIOS               |
| `infoblox_extensible_attribute_def` | Manages Extensible Attribute Definitions  | Retrieves information about existing Extensible Attribute Definitions | NIOS               |
| `infoblox_nat_group`                | Manages Grid NAT Groups                   | Retrieves information about existing Grid NAT Groups                  | NIOS               |
| `infoblox_servicerestart_group`     | Manages Grid Service Restart Groups       | Retrieves information about existing Grid Service Restart Groups      | NIOS               |
| `infoblox_upgrade_group`            | Manages Grid Upgrade Groups               | Retrieves information about existing Grid Upgrade Groups              | NIOS               |

### INFRA

| Name                     | Resource Description            | Data Source Description                                      | Supported Backends |
|--------------------------|---------------------------------|--------------------------------------------------------------|--------------------|
| `infoblox_infra_host`    | Manages Infrastructure Hosts    | Retrieves information about existing Infrastructure Hosts    | UDDI               |
| `infoblox_infra_service` | Manages Infrastructure Services | Retrieves information about existing Infrastructure Services | UDDI               |
| `infoblox_join_token`    | Manages Join Tokens             | Retrieves information about existing Join Tokens             | UDDI               |

### IPAM

| Name                                     | Resource Description                 | Data Source Description                                                        | Supported Backends |
|------------------------------------------|--------------------------------------|--------------------------------------------------------------------------------|--------------------|
| `infoblox_address`                       | Manages IPAM Addresses               | Retrieves information about existing IPAM Addresses                            | UDDI               |
| `infoblox_bulk_hostname_template`        | Manages IPAM Bulk Hostname Templates | Retrieves information about existing IPAM Bulk Hostname Templates              | NIOS               |
| `infoblox_ipam_host`                     | Manages IPAM Hosts                   | Retrieves information about existing IPAM Hosts                                | UDDI               |
| `infoblox_ipv6_network`                  | Manages IPAM IPv6 Networks           | Retrieves information about existing IPAM IPv6 Networks                        | NIOS, UDDI         |
| `infoblox_ipv6_network_container`        | Manages IPAM IPv6 Network Containers | Retrieves information about existing IPAM IPv6 Network Containers              | NIOS, UDDI         |
| `infoblox_network`                       | Manages IPAM Networks                | Retrieves information about existing IPAM Networks                             | NIOS, UDDI         |
| `infoblox_network_container`             | Manages IPAM Network Containers      | Retrieves information about existing IPAM Network Containers                   | NIOS, UDDI         |
| `infoblox_network_view`                  | Manages IPAM Network Views           | Retrieves information about existing IPAM Network Views                        | NIOS, UDDI         |
| `infoblox_next_available_address_blocks` |                                      | Retrieves the next available address blocks in an address block                | UDDI               |
| `infoblox_next_available_ips`            |                                      | Retrieves the next available IP addresses in an address block, subnet or range | UDDI               |
| `infoblox_next_available_subnets`        |                                      | Retrieves the next available subnets in an address block                       | UDDI               |
| `infoblox_superhost`                     | Manages IPAM Super Hosts             | Retrieves information about existing IPAM Super Hosts                          | NIOS               |
| `infoblox_vlan`                          | Manages IPAM VLANs                   | Retrieves information about existing IPAM VLANs                                | NIOS               |
| `infoblox_vlan_range`                    | Manages IPAM VLAN Ranges             | Retrieves information about existing IPAM VLAN Ranges                          | NIOS               |
| `infoblox_vlan_view`                     | Manages IPAM VLAN Views              | Retrieves information about existing IPAM VLAN Views                           | NIOS               |

### IPAM FEDERATION

| Name                       | Resource Description     | Data Source Description                               | Supported Backends |
|----------------------------|--------------------------|-------------------------------------------------------|--------------------|
| `infoblox_federated_realm` | Manages Federated Realms | Retrieves information about existing Federated Realms | UDDI               |

### KEYS

| Name                | Resource Description | Data Source Description                        | Supported Backends |
|---------------------|----------------------|------------------------------------------------|--------------------|
| `infoblox_tsig_key` | Manages TSIG Keys    | Retrieves information about existing TSIG Keys | UDDI               |

### MISCELLANEOUS

| Name                    | Resource Description  | Data Source Description                            | Supported Backends |
|-------------------------|-----------------------|----------------------------------------------------|--------------------|
| `infoblox_bfd_template` | Manages BFD Templates | Retrieves information about existing BFD Templates | NIOS               |
| `infoblox_ruleset`      | Manages Rule Sets     | Retrieves information about existing Rule Sets     | NIOS               |

### NOTIFICATION

| Name                                  | Resource Description                | Data Source Description                                          | Supported Backends |
|---------------------------------------|-------------------------------------|------------------------------------------------------------------|--------------------|
| `infoblox_notification_rest_endpoint` | Manages Notification REST Endpoints | Retrieves information about existing Notification REST Endpoints | NIOS               |
| `infoblox_notification_rule`          | Manages Notification Rules          | Retrieves information about existing Notification Rules          | NIOS               |

### REDIRECT

| Name                       | Resource Description     | Data Source Description                               | Supported Backends |
|----------------------------|--------------------------|-------------------------------------------------------|--------------------|
| `infoblox_custom_redirect` | Manages Custom Redirects | Retrieves information about existing Custom Redirects | UDDI               |

### RIR

| Name                        | Resource Description      | Data Source Description                                | Supported Backends |
|-----------------------------|---------------------------|--------------------------------------------------------|--------------------|
| `infoblox_rir_organization` | Manages RIR Organizations | Retrieves information about existing RIR Organizations | NIOS               |

### RPZ

| Name                                          | Resource Description                           | Data Source Description                                                     | Supported Backends |
|-----------------------------------------------|------------------------------------------------|-----------------------------------------------------------------------------|--------------------|
| `infoblox_record_rpz_a`                       | Manages RPZ A Records                          | Retrieves information about existing RPZ A Records                          | NIOS               |
| `infoblox_record_rpz_a_ipaddress`             | Manages RPZ A IP Address Records               | Retrieves information about existing RPZ A IP Address Records               | NIOS               |
| `infoblox_record_rpz_aaaa`                    | Manages RPZ AAAA Records                       | Retrieves information about existing RPZ AAAA Records                       | NIOS               |
| `infoblox_record_rpz_aaaa_ipaddress`          | Manages RPZ AAAA IP Address Records            | Retrieves information about existing RPZ AAAA IP Address Records            | NIOS               |
| `infoblox_record_rpz_cname`                   | Manages RPZ CNAME Records                      | Retrieves information about existing RPZ CNAME Records                      | NIOS               |
| `infoblox_record_rpz_cname_clientipaddress`   | Manages RPZ CNAME Client IP Address Records    | Retrieves information about existing RPZ CNAME Client IP Address Records    | NIOS               |
| `infoblox_record_rpz_cname_clientipaddressdn` | Manages RPZ CNAME Client IP Address DN Records | Retrieves information about existing RPZ CNAME Client IP Address DN Records | NIOS               |
| `infoblox_record_rpz_cname_ipaddress`         | Manages RPZ CNAME IP Address Records           | Retrieves information about existing RPZ CNAME IP Address Records           | NIOS               |
| `infoblox_record_rpz_cname_ipaddressdn`       | Manages RPZ CNAME IP Address DN Records        | Retrieves information about existing RPZ CNAME IP Address DN Records        | NIOS               |
| `infoblox_record_rpz_naptr`                   | Manages RPZ NAPTR Records                      | Retrieves information about existing RPZ NAPTR Records                      | NIOS               |
| `infoblox_record_rpz_ptr`                     | Manages RPZ PTR Records                        | Retrieves information about existing RPZ PTR Records                        | NIOS               |
| `infoblox_record_rpz_txt`                     | Manages RPZ TXT Records                        | Retrieves information about existing RPZ TXT Records                        | NIOS               |

### SECURITY

| Name                  | Resource Description | Data Source Description                          | Supported Backends |
|-----------------------|----------------------|--------------------------------------------------|--------------------|
| `infoblox_admin_user` | Manages Admin Users  | Retrieves information about existing Admin Users | NIOS               |
