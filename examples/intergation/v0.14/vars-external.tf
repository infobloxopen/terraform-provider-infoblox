locals {
  external_netview_name = "${local.nios_name_prefix}_external_netview"
  external_default_dnsview = "default.${local.external_netview_name}"
  external_ipv4_pool = "1.0.2.128/25"
  external_ipv4_pool_reserve_ip = 5 # for network IP and a few special servers
  external_ipv4_pool_gateway = "1.0.2.250"
  external_ipv4_pool_mailserver = "1.0.2.249"

  external_ipv6_pool = "2a00:1148:1c9d::/48"
  external_ipv6_pool_reserve_ip = 25 # for network IP and a few special servers
  external_ipv6_pool_gateway = "2a00:1148:1c9d::fe"
  external_ipv6_pool_mailserver = "2a00:1148:1c9d::fd"

  external_vm_mail_nic1_mac = "12:17:0f:20:a2:3c"
  external_vm_mail_nic1_duid = "63:2d:ff:89:ab:43:c9:90:6f"
}
