locals {
  internal_ipv4_pool = "172.18.2.128/25"
  internal_ipv4_pool_reserve_ip = 4 // for network IP and a few special servers
  internal_ipv4_pool_gateway = "172.18.2.1"

  internal_ipv6_pool = "2000::/48"
  internal_ipv6_pool_reserve_ip = 16 // for network IP and a few special servers
  internal_ipv6_pool_gateway = "2000::fe"

  internal_reception_1_nic1_mac = "32:5f:d6:98:00:ba"
  internal_reception_1_nic1_duid = "33:21:6d:dd:f1:a4:cb:b6:c9"

  internal_reception_2_nic1_mac = "32:5f:d6:98:00:bb"
  internal_reception_2_nic1_duid = "33:21:6d:dd:f1:a4:cb:b6:ca"

  internal_user_1_nic1_mac = "32:5f:d6:98:30:bd"
  internal_user_1_nic1_duid = "33:21:6d:dd:f1:a4:db:b6:cd"

  internal_user_2_nic1_mac = "32:5f:d6:98:30:bc"
  internal_user_2_nic1_duid = "33:21:6d:dd:f1:a4:db:b6:cb"

  internal_web_server = local.default_internal_dns_zone
  internal_web_server_cname = "www.${local.default_internal_dns_zone}"

  internal_web_server_ipv4_addr = "172.18.2.2"
}
