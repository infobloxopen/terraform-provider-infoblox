# IPv6 Network Resource

The `infoblox_ipv6_network` resource enables you to perform `create`, `update` and `delete` operations
on IPv6 networks. Network resources support the next available network feature when you use
the `allocate_prefix_len` parameter in the below list.

The following list describes the parameters you can define in a `infoblox_ipv6_network` resource block:

* `network_view`: optional, specifies the network view in which to create the network; the default value is `default`.
* `cidr`: required only if `parent_cidr` is not set; specifies the network block to use for the network, in CIDR notation. Do not use an IPv4 CIDR for an IPv6 network. If you configure both `cidr` and `parent_cidr`, the value of `parent_cidr` is ignored.
* `parent_cidr`: required only if `cidr` is not set; specifies the network container from which the network must be dynamically allocated. The network container must exist in the NIOS database, but not necessarily as a Terraform resource.
* `allocate_prefix_len`: required only if `parent_cidr` is set; defines the length of the network part of the address for a network that should be allocated from a network container, which in turn is determined by `parent_cidr`.
* `gateway`: optional, defines the IP address of the gateway within the network block. If a value is not set, the first IP address of the allocated network is assigned as the gateway address. If the value of the gateway parameter is set as `none`, no value is assigned.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network.
* `reserve_ipv6`: optional, specifies the number of IPv6 addresses that you want to reserve in the IPv6 network. The default value is 0

!> Once a network object is created, the `reserve_ipv6` and `gateway` fields cannot be edited.

!> IP addresses that are reserved by setting the `reserve_ipv6` field are used for network maintenance by the cloud providers. Therefore, Infoblox does not recommend using these IP addresses for other purposes.

### Examples of an IPv6 Network Block

```hcl
// statically allocated IPv6 network, minimal set of parameters
resource "infoblox_ipv6_network" "net1" {
  cidr = "2002:1f93:0:3::/96"
}

// full set of parameters for statically allocated IPv6 network
resource "infoblox_ipv6_network" "net2" {
  cidr = "2002:1f93:0:4::/96"
  network_view = "nondefault_netview"
  reserve_ip = 10
  gateway = "2002:1f93:0:4::1"
  comment = "let's try IPv6"
  ext_attrs = jsonencode({
    "Site" = "somewhere in Antarctica"
  })
}

// full set of parameters for dynamically allocated IPv6 network
resource "infoblox_ipv6_network" "net3" {
  parent_cidr = infoblox_ipv6_network_container.v6net_c1.cidr // reference to the resource from another example
  allocate_prefix_len = 100 // 96 (existing network container) + 4 (new network), prefix
  network_view = "default" // we may omit this but it is not a mistake to specify explicitly
  reserve_ip = 20
  gateway = "none" // no gateway defined for this network
  comment = "the network for the Test Lab"
  ext_attrs = jsonencode({
    "Site" = "small inner cluster"
  })
}
```
