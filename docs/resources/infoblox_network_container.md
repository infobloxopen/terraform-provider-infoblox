# Network container

A network container has two similar resources for different versions of
NIOS network container: 'infoblox_ipv4_network_container' and
'infoblox_ipv6_network_container'. They are almost identical. The
difference is about which type of network range IP address you may use:
IPv4 or IPv6.

A network container has attributes:

| Attribute | Required/optional | Description | Example |
| --- | --- | --- | --- |
| network_view | required | which network view to create the network container within; 'default' value is used when the attribute is omitted | ext_org_temporary |
| cidr | required | which network block to use for the network container, in CIDR notation; it is an error to use IPv4 CIDR for IPv6 network container and vice versa. | 2001:db8::/64 |

Both attributes are mandatory and cannot be changed (by UPDATE
operation) once the network container is created.

Examples of resource blocks:

    resource "infoblox_ipv4_network_container" "nc1" {
      network_view = "default"
      cidr = "10.20.30.192/28"
      comment = "this is an example of network container"
      ext_attrs = jsonencode({
        "Tenant ID" = "ISP 1"
      })
    }
    
    resource "infoblox_ipv6_network_container" "nc2" {
      network_view = "very_special_network_view"
      cidr = "2a00:1148::/32"
    }
