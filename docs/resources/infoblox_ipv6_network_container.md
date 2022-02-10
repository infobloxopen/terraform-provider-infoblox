# IPv6 Network Container

The `infoblox_ipv6_network_container` resource, enables you to create, update, or delete an IPv6 network container in a NIOS appliance.

The following list describes the parameters you can define in the network container resource block:

* `network_view`: required, specifies the network view in which to create the network container; if a value is not specified, the network container will be created in the default network view defined in NIOS.
* `cidr`: required, specifies the network block to use for the network container; do not use an IPv4 CIDR for an IPv6 network and vice versa.
* `comment`: optional, describes the network container.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network container.

!> Once the network container is created, the network_view and cidr parameter values cannot be changed by performing an update operation.

### Examples of the Network Container Resource

```hcl
resource "infoblox_ipv6_network_container" "nc2" {
  network_view = "very_special_network_view"
  cidr = "2a00:1148::/32"
  comment = "this is an example of network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Site" = "Nevada" 
  })
}
```
