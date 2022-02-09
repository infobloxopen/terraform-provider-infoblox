# IPv6 Network Resource

The `infoblox_ipv6_network` resource enables you to perform create, update, and delete operations on IPv6 networks. Network resources support the next available network feature by using the `allocate_prefix_len` parameter from the list the below.

The following list describes the parameters you can define in a network resource block:

* `network_view`: optional, specifies the network view in which to create the network; the default value is `default`.
* `cidr`: required only if `parent_cidr` is not set, specifies the network block to use for the network, in CIDR notation; do not use the IPv4 CIDR for IPv6 network and vice versa.
* `parent_cidr`: required only if `cidr` is not set, specifies the network container from which the network must be dynamically allocated; the network container must exist in the NIOS database, but not necessarily as a Terraform resource.
* `allocate_prefix_len`: required only if `parent_cidr` is set, defines the length of the network part of the address for a network that should be allocated from a network container, which in turn is determined by `parent_cidr`.
* `gateway`: optional, represents the IP address of the gateway within the network block; if you do not specify an address, by default, the first IP address is set as the gateway address. For more information, see [Limitations] (#limitations).
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network.
* `reserve_ipv6`: optional, specifies the number of IPv6 addresses that you want to reserve in the IPv6 network. The default value is 0.

-> Either `cidr` or the combination of `parent_cidr` and `allocate_prefix_len` is required. The rest of the parameters are optional.

-> IPv6 addresses that are reserved by setting the `reserve_ipv6` field are used for network maintenance by the cloud providers. Therefore, Infoblox does not recommend using these IP addresses for other purposes.

-> When creating an IPv6 network using the `reserve_ipv6` flag, the DUIDs assigned to the reserved IPv6 addresses that you choose, may not be in standard format.

-> Tenant ID does not display in the IPv6 network container on NIOS.

!> Once a network object is created, the `reserve_ip` and `gateway` fields cannot be edited.

### Examples of the Network Block

```hcl
resource "infoblox_ipv6_network_container" "nc1" {
  network_view = "default"
  cidr = "2a00:1228:32bf:22ad:/64"
}

// Static allocation of a network
resource "infoblox_ipv6_network" "nw1" {
  network_view = "very_special_network_view"
  cidr = "2a00:1148::/32"
  comment = "mockup network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Site" = "Nevada" 
  })
}

// Dynamic allocation of a network
resource "infoblox_ipv6_network" "nw2" {
  // The 'network_view' attribute is omitted,
  // thus is implied to be 'default'
  parent_cidr = infoblox_ipv6_network_container.nc1.cidr
  allocate_prefix_len = 96
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Custom EA 1" = "category 14"
  })
}

// "2a00:1148::/32" network container in the
// example below is supposed to exist in 'default'
// network view. We have not created it here so it is
// implied that it was created by other means.

resource "infoblox_ipv6_network" "nw3" {
  // we want to create a network with /64 hosts
  allocate_prefix_len = 64

     // inside the network container "2a00:1148::/32"
  parent_cidr = "2a00:1148::/32"

  // in 'default' network view
  network_view = "default"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Site" = "Nevada"
  })
}
```
