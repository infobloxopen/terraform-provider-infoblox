# IPv4 Network Data Source

The data source for the network object allows you to get the following parameters for an IPv4 network resource:

* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: The set of extensible attributes, if any. The content is formatted as a JSON map. Example: `{"Owner": "public internet caffe", "Administrator": "unknown"}`.

To get information about a network, you must specify a combination of the network view and
network address in the CIDR format.
The following list describes the parameters you must define in an `infoblox_ipv4_network` data source block (all of them are required):

* `network_view`: optional, specifies the network view which the network container exists in. If a value is not specified, the name `default` is used as the network view.
* `cidr`: specifies the network block which correcponds to the network, in CIDR notation. Do not use the IPv6 CIDR for an IPv4 network.

### Supported Arguments for filters

-----
| Field        | Alias | Type   | Searchable |
|--------------|-------|--------|------------|
| network      | cidr  | string | yes        |
| network_view | ---   | string | yes        |
| comment      | ---   | string | yes        |

Note: Please consider using only fields as the keys in terraform datasource, kindly don't use alias names as keys from the above table.

### Example of a Network Data Source Block

```hcl
resource "infoblox_ipv4_network" "net2" {
  cidr = "192.168.128.0/20"
  network_view = "nondefault_netview"
  reserve_ip = 5
  gateway = "192.168.128.254"
  comment = "small network for testing"
  ext_attrs = jsonencode({
    "Site" = "bla-bla-bla... testing..."
  })
}

data "infoblox_ipv4_network" "nearby_network" {
  network_view = "nondefault_netview"
  cidr = "192.168.128.0/20"

  // This is just to ensure that the network has been be created
  // using 'infoblox_ipv4_network' resource block before the data source will be queried.
  depends_on = [infoblox_ipv4_network.net2]
}

output "nearby_network_comment" {
  value = data.infoblox_ipv4_network.nearby_network.comment
}

output "nearby_network_ext_attrs" {
  value = data.infoblox_ipv4_network.nearby_network.ext_attrs
}
```
