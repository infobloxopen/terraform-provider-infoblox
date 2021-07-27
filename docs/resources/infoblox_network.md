# Network

Similarly to Network container, network resource has two versions:
'infoblox_ipv4_network' and 'infoblox_ipv6_network'. The set of
attributes is the following:

| Attribute | Required/optional | Description | Example |
| --- | --- | --- | --- |
| network_view | optional | which network view to create the network within, the default value is 'default' | external |
| cidr| required for static network allocation | which network block to use for the network, in CIDR notation; it is an error to use IPv4 CIDR for IPv6 network container and vice versa | 10.3.2.0/24 | 
| parent_cidr | required for dynamic network allocation | if 'cidr' is not defined, this attribute is mandatory, it denotes a network container (which must exist in NIOS DB, not necessarily as a TF resource) which the network is to be dynamically allocated from. | 10.0.0.0/8 |
| allocate_prefix_len | required for dynamic network allocation | defines the length of the network part of the address for the network to be allocated from a network container, determined by 'parent_cidr'. | 30 |
| reserve_ip | optional | makes sense only for an IPv4 network, determines the number of IP addresses you want to reserve in the IPv4 Network; the default value is 0. | 5 |
| gateway | optional | for IPv4 networks represents the gateway's IP address within the network block; by default, the first IPv4 address is set as gateway address. | 192.168.32.254 |
| reserve_ipv6 | optional | makes sense only for an IPv6 network, determines the number of IP addresses you want to reserve in the IPv6 Network; the default value is 0. | 7 |

The main point here is: **either** 'cidr' **or** the pair of
'parent_cidr' and 'allocate_prefix_len' is **mandatory**. The rest of
the attributes are optional.

Examples of Network resource block:

    resource "infoblox_ipv4_network_container" "nc1" {
      network_view = "default"
      cidr = "192.168.30.0/24"
    }
    
    resource "infoblox_ipv6_network" "nw1" {
      network_view = "very_special_network_view"
      cidr = "2a00:1148::/32"
      comment = "just some dummy network"
    }
    
    // this is to dynamically allocate the
    // 192.168.30.0/30 network within
    // 192.168.30.0/24 network container
    resource "infoblox_ipv4_network" "nw2" {
      // The 'network_view' attribute is omitted,
      // thus is implied to be 'default'
      parent_cidr = "192.168.30.0/24"
      allocate_prefix_len = 30
      ext_attrs = jsonencode({
        "Custom EA 1" = "category 14"
      })
    }
    
    // "2a00:1148::/32" network container is supposed
    // to exist in 'default' network view.
    // We have not created it here so it is
    // implied that it was created by other means.
    resource "infoblox_ipv6_network" "nw3" {
      // we want to create a network with 2^64 hosts
      allocate_prefix_len = 64
     
      // inside the network container "2a00:1148::/32"
      parent_cidr = "2a00:1148::/32"
    
      // in 'default' network view
      network_view = "default"
    }
