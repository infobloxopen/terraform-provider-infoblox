# IPv4 network

For a network resource appropriate data source allows to get the
following attributes:

| Attribute | Description | Example |
| --- | --- | --- |
| comment | A text describing the network, a regular comment. | Untrusted network |
| ext_attrs | A set of extensible attributes of the record, if any. The content is in JSON map format. | {"Owner": "public internet caffe", "Administrator": "unknown"} |

To get information about a network, you have to specify a selector which
uniquely identifies it: a combination of network view ('network_view'
field) and a network's address, in CIDR format ('cidr' field). All the
fields are required. Currently, only IPv4 networks are supported, not
IPv6.

## Example

    data "infoblox_ipv4_network" "nearby_network" {
      network_view = "default"
      cidr = "192.168.128.0/20"
    }
