# IPv4 Network Resource

The `infoblox_ipv4_network` resource enables you to perform create, update, and delete operations on IPv4 networks. Network resources support the next available network feature by using the `allocate_prefix_len` parameter from the list the below.

The following list describes the parameters you can define in a network resource block:

* `network_view`: optional, specifies the network view in which to create the network; the default value is `default`.
* `cidr`: required only if `parent_cidr` is not set, specifies the network block to use for the network, in CIDR notation; do not use the IPv4 CIDR for IPv6 network and vice versa.
* `parent_cidr`: required only if `cidr` is not set, specifies the network container from which the network must be dynamically allocated; the network container must exist in the NIOS database, but not necessarily as a Terraform resource.
* `allocate_prefix_len`: required only if `parent_cidr` is set, defines the length of the network part of the address for a network that should be allocated from a network container, which in turn is determined by `parent_cidr`.
* `gateway`: optional, represents the IP address of the gateway within the network block; if you do not specify an address, by default, the first IP address is set as the gateway address. For more information, see [Limitations] (#limitations).
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network.
* `reserve_ip`: optional, specifies the number of IPv4 addresses that you want to reserve in the IPv4 network; the default value is 0.

-> Either `cidr` or the combination of `parent_cidr` and `allocate_prefix_len` is required. The rest of the parameters are optional.

-> IP addresses that are reserved by setting the `reserve_ip` field are used for network maintenance by the cloud providers. Therefore, Infoblox does not recommend using these IP addresses for other purposes.

!> Once a network object is created, the `reserve_ip` and `gateway` fields cannot be edited.

### Examples of the Network Block

```hcl
resource "infoblox_ipv4_network_container" "nc1" {
  network_view = "default"
  cidr = "192.168.30.0/24"
}

// A statically allocated network in a non-default network view.
resource "infoblox_ipv4_network" "nw1" {
  network_view = "very_special_network_view"
  cidr = "10.1.2.128/25"
  comment = "mockup network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Site" = "Nevada" 
  })
}

resource "infoblox_ipv4_network" "nw2" {
  // The 'network_view' attribute is omitted,
  // thus is implied to be 'default'
  parent_cidr = infoblox_ipv4_network_container.nc1.cidr
  allocate_prefix_len = 30
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
     "Custom EA 1" = "category 14"
  })
}
```
