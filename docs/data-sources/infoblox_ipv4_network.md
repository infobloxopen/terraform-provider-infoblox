# IPv4 Network Data Source

The data source for the network object allows you to get the following parameters for an IPv4 network resource:

* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: The set of extensible attributes, if any. The content is formatted as a JSON map. Example: `{"Owner": "public internet caffe", "Administrator": "unknown"}`.

To get information about a network, you must specify a combination of the network view and
network address in the CIDR format.
The following list describes the parameters you must define in an `infoblox_ipv4_network` data source block (all of them are required):

* `network_view`: optional, specifies the network view which the network container exists in. If a value is not specified, the name `default` is used as the network view.
* `cidr`: specifies the network block which correcponds to the network, in CIDR notation. Do not use the IPv6 CIDR for an IPv4 network.

### Example of a Network Data Source Block

```hcl
data "infoblox_ipv4_network" "nearby_network" {
  network_view = "default"
  cidr = "192.168.128.0/20"
}

output "nearby_network_comment" {
  value = data.infoblox_ipv4_network.nearby_network.comment
}

output "nearby_network_ext_attrs" {
  value = data.infoblox_ipv4_network.nearby_network.ext_attrs
}
```
