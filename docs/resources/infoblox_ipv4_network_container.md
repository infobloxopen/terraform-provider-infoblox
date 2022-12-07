# IPv4 Network Container

The `infoblox_ipv4_network_container` resource, enables you to create, update,
or delete an IPv4 network container in a NIOS appliance.

The following list describes the parameters you can define in the network container
resource block:

* `network_view`: optional, specifies the network view in which to create the network container; if a value is not specified, the name `default` is used as the network view.
* `cidr`: required, specifies the network block to use for the network container; do not use an IPv6 CIDR for an IPv4 network container.
* `comment`: optional, describes the network container.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network container.

!> Once the network container is created, the `network_view` and `cidr` parameter values cannot be changed by performing an `update` operation.

### Examples of the Network Container Resource

```hcl
// statically allocated IPv4 network container, minimal set of parameters
resource "infoblox_ipv4_network_container" "v4net_c1" {
  cidr = "10.2.0.0/24"
}

// full set of parameters for statically allocated IPv4 network container
resource "infoblox_ipv4_network_container" "v4net_c2" {
  cidr = "10.2.0.0/24" // we may allocate the same IP address range but in another network view
  network_view = "nondefault_netview"
  comment = "one of our clients"
  ext_attrs = jsonencode({
    "Site" = "remote office"
    "Country" = "Australia"
  })
}

// so far, we do not support dynamic allocation of network containers
```
