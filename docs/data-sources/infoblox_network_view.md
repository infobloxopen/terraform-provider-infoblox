# Network View Data Source

Use the data source to retrieve the following information for a network view resource from the corresponding object in NIOS:

* `comment`: a description of the network view. This is a regular comment. Example: `From the outside`.
* `ext_attrs`: the set of extensible attributes of the network view, if any. The content is formatted as a JSON map. Example: `{"Administrator": "jsw@telecom.ca"}`.

To get information about a network view, you must specify a name of the network view.

### Example of a Network View Data Source Block

```hcl
data "infoblox_network_view" "inet_nv" {
  name = "inet_visible_nv"
}

output "inet_nv_ext_attrs" {
  value = data.infoblox_network_view.inet_nv.ext_attrs
}

output "inet_nv_comment" {
  value = data.infoblox_network_view.inet_nv.comment
}
```
