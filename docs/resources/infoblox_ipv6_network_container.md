# IPv6 Network Container

The `infoblox_ipv6_network_container` resource, enables you to create, update,
or delete an IPv6 network container in a NIOS appliance.

The following list describes the parameters you can define in the network container
resource block:

* `network_view`: optional, specifies the network view in which to create the network container; if a value is not specified, the name `default` is used as the network view.
* `cidr`: required, specifies the network block to use for the network container; do not use an IPv4 CIDR for an IPv6 network container.
* `comment`: optional, describes the network container.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network container.

!> Once the network container is created, the `network_view` and `cidr` parameter values cannot be changed by performing an `update` operation.

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

// so far, we do not support dynamic allocation of network containers
```
