# IPv4 Network Container Data Source

Use the data source to retrieve the following information for an IPv4 network container resource from the corresponding
object in NIOS:

* `comment`: a description of the network container. This is a regular comment. Example: `Tenant 1 network container`.
* `ext_attrs`: the set of extensible attributes of the network view, if any. The content is formatted as a JSON map. Example: `{"Administrator": "jsw@telecom.ca"}`.

To get information about a network container, specify a combination of
the network view and the address of the network block in CIDR format.
The following list describes the parameters you must define
in an `infoblox_ipv4_network_container` data source block (all of them are required):

* `network_view`: specifies the network view in which the network container exists.
* `cidr`: specifies the IPv4 network block of the network container.

### Example of an IPv4 Network Container Data Source Block

```hcl
data "infoblox_ipv4_network_container" "nearby_org" {
  network_view = "separate_tenants"
  cidr = "192.168.128.0/16"
}

output "nearby_org_comment" {
  value = data.infoblox_ipv4_network_container.nearby_org.comment
}

output "nearby_org_ext_attrs" {
  value = data.infoblox_ipv4_network_container.nearby_org.ext_attrs
}
```
