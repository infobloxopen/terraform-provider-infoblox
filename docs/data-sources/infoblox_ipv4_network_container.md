# IPv4 Network Container Data Source

The data source `infoblox_ipv4_network_container` allows you to get the following
parameters for an IPv4 network container resource from the corresponding NIOS object:

* `comment`: a description of the network container. This is a regular comment. Example: `Tenant 1 network container`.
* `ext_attrs`: a set of extensible attributes, if any. The content is formatted as a JSON map. Example: `{"Administrator": "jsw@telecom.ca"}`.

To get information about a network container, you must specify a combination of the network view and
network block address in the CIDR format. The following list describes the parameters you must define
in an `infoblox_ipv4_network_container` data source block:

* `network_view`: the network view in which the network is to be created. The default value is `default`. If a value is not specified, the default network view defined in NIOS is considered.
* `cidr`: the network block in the CIDR notation that is used for the network container.

### Example of an IPv4 Network Container Data Source Block

```hcl
data "infoblox_ipv4_network_container" "nearby_org" {
  network_view = "separate_tenants"
  cidr = "192.168.128.0/16"
}

output "comment" {
  value = data.infoblox_ipv4_network_container.nearby_org.comment
}

output "ext_attrs" {
  value = data.infoblox_ipv4_network_container.nearby_org.ext_attrs
}
```
