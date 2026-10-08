# Roadmap

The Terraform Provider for Infoblox brings NIOS and Universal DDI under a single provider, and continues to expand across both backends. The objects and capabilities below are planned for upcoming releases.

## NIOS

| Category | Objects |
|---|---|
| Grid and member management | Grid members, Grid join, upgrade schedules |
| Security and authentication | Admin groups and roles, permissions, LDAP, RADIUS, SAML, TACACS+ and certificate authentication services, SNMP and FTP users |
| DHCP | MAC, NAC, relay agent and fingerprint filters, fingerprints, MAC filter addresses, roaming hosts, failover, IPv6 ranges |
| Object templates | Network templates, IPv6 network templates, IPv4 fixed address templates |
| DNS records | TLSA records, Unknown records for record types such as SPF and RP |
| Response Policy Zones | RPZ MX and SRV records |
| DTC | DTC A, AAAA, CNAME, NAPTR and SRV records, SIP monitors |
| Microsoft integration | Microsoft servers, Active Directory sites, superscopes |
| Discovery | vDiscovery tasks, AWS Route 53 task groups |
| Parental control | AVPs, blocking policies, subscriber records and sites |
| Smart folders | Global and personal smart folders |
| External endpoints | DXL, syslog and TFTP file distribution endpoints |

## Universal DDI

| Category | Objects |
|---|---|
| DHCP | DHCP servers, DHCP hosts |
| DNS | A generic record type, for record types without a dedicated resource |
| IPAM Federation | Federated blocks, extending the existing IPAM Federation support |
| Infrastructure services | DNS Forwarding Proxy (DFP) |

## Modules

Terraform modules for deploying NIOS on AWS and Azure are planned for later releases. Deployment on GCP through the Google Cloud Marketplace is already available in [nios-public-cloud-marketplace/gcp](nios-public-cloud-marketplace/gcp/).

## Requesting an Object or Feature

To request support for something that is not listed here, open an issue on the [GitHub Issues page](https://github.com/infobloxopen/terraform-provider-infoblox/issues).
