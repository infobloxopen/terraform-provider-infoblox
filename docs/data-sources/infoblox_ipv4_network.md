# IPv4 Network Datasource

The data source `infoblox_ipv4_network` for the network object allows you to get the following parameters for an IPv4 network resource:

* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: a set of extensible attributes, if any. The content is formatted as a JSON map. Example: `{"Owner": "public internet caffe", "Administrator": "unknown"}`.

To get information about a network, you must specify a combination of the network view and network address in the CIDR format. The following list describes the parameters you must define in an `infoblox_ipv4_network` data source block:

* `network_view`: the network view in which the network is to be created. The default value is `default`. If a value is not specified, the default network view defined in NIOS is considered.
* `cidr`: the network block in the CIDR notation that is used for the network. Do not use the IPv4 CIDR for an IPv6 network or the IPv6 CIDR for an IPv4 network.

### Example of a Network Data Source Block

```hcl
data "infoblox_ipv4_network" "nearby_network" {
  network_view = "default"
  cidr = "192.168.128.0/20"
}
```
