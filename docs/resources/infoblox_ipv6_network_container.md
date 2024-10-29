# IPv6 Network Container

The `infoblox_ipv6_network_container` resource, enables you to create, update,
or delete an IPv6 network container in a NIOS appliance.

The following list describes the parameters you can define in the network container
resource block:

* `network_view`: optional, specifies the network view in which to create the network container; if a value is not specified, the name `default` is used as the network view.
* `cidr`: required only if `parent_cidr` is not set, specifies the network block to use for the network container; do not use an IPv4 CIDR for an IPv6 network container.
* `parent_cidr`: required only if `cidr` is not set, specifies the network container from which next available network container must be allocated.
* `allocate_prefix_len`: required only if `parent_cidr` is set, defines length of netmask for a network container that should be allocated from network container, determined by `parent_cidr`.
* `comment`: optional, describes the network container.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network container.
* `filter_params`: required for dynamic allocation when `parent_cidr` is not used, specifies the extensible attributes of the parent network container that must be used as filters to retrieve the next available network for creating the network container object. Example: `jsonencode({"*Site": "Turkey"})`.

* !> Once the network container is created, the `network_view` and `cidr` parameter values cannot be changed by performing an `update` operation.

!> Once the network container is created dynamically, the `parent_cidr`, `filter_params` and `allocate_prefix_len` parameter values cannot be changed.

### Examples of the Network Container Resource

```hcl
// statically allocated IPv6 network container, minimal set of parameters
resource "infoblox_ipv6_network_container" "v6net_c1" {
  cidr = "2002:1f93:0:1::/96"
}

// full set of parameters for statically allocated IPv6 network container
resource "infoblox_ipv6_network_container" "v6net_c2" {
  cidr = "2002:1f93:0:2::/96"
  network_view = "nondefault_netview"
  comment = "new generation network segment"
  ext_attrs = jsonencode({
    "Site" = "space station"
    "Country" = "Earth orbit"
  })
}

// full set of parameters for dynamic allocation of network containers
resource "infoblox_ipv6_network_container" "v6net_c3" {
  parent_cidr = infoblox_ipv6_network_container.v6net_c2.cidr
  allocate_prefix_len = 97
  network_view = infoblox_ipv6_network_container.v6net_c2.network_view
  comment = "dynamic allocation of network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}

// dynamic allocation of IPv6 network container resource using filter_params
resource "infoblox_ipv6_network_container" "network_container_ipv6" {
  allocate_prefix_len = 68
  comment = "IPv6 network container created with next available network"
  filter_params = jsonencode({
    "*Site": "Blr"
  })
}
```
