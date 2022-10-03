# Network View Data Source

The data source `infoblox_network_view` allows you to get the following
parameters for a network view resource from the corresponding NIOS object:

* `comment`: a description of the network container. This is a regular comment. Example: `From the outside`.
* `ext_attrs`: a set of extensible attributes, if any. The content is formatted as a JSON map. Example: `{"Administrator": "jsw@telecom.ca"}`.

To get information about a network view, you must specify a name of the network view.

### Example of a Network View Data Source Block

```hcl
data "infoblox_network_view" "inet_nv" {
  name = "inet_visible_nv"
}

output "ext_attrs" {
  value = data.infoblox_network_view.inet_nv.ext_attrs
}

output "comment" {
  value = data.infoblox_network_view.inet_nv.comment
}
```
