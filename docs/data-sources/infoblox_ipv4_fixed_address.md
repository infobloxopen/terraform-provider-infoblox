# Fixed Address Data Source

The `infoblox_ipv4_fixed_address` data source to retrieve the following information for an fixed address in the network if any, which is managed by a NIOS server:

* `agent_circuit_id`: The agent circuit ID for the fixed address. The field is required only when match_client is set to CIRCUIT_ID. Example: `23`
* `agent_remote_id`: The agent remote ID for the fixed address. The field is required only when match_client is set to REMOTE_ID. Example: `34`
* `client_identifier_prepend_zero`: This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address. Example: `false`
* `comment`: Comment for the fixed address; maximum 256 characters. Example: `fixed address`
* `dhcp_client_identifier`: The DHCP client ID for the fixed address. The field is required only when match_client is set to CLIENT_ID. Example: `20`
* `disable`: Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled. Example: `false`
* `ext_attrs`: Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `ipv4addr`: The IPv4 Address of the fixed address. If the `ipv4addr` field is not provided and the `network` field is set, the next available IP address in the network will be allocated. Example: `10.0.0.34`
* `mac`: The MAC address value for this fixed address. The field is required only when match_client is set to its default value - MAC_ADDRESS. Example: `00-1A-2B-3C-4D-5E`
* `match_client`: The match client for the fixed address.Valid values are CIRCUIT_ID, CLIENT_ID , MAC_ADDRESS, REMOTE_ID and RESERVED. Default value is MAC_ADDRESS. Example: `CLIENT_ID`
* `name`: This field contains the name of this fixed address. Example: `fixedAddressName`
* `network`: The network to which this fixed address belongs, in IPv4 Address/CIDR format. Example: `10.0.0.0/24`
* `network_view`: The name of the network view in which this fixed address resides. The default value is The default network view. Example: `default`
* `options`: An array of DHCP option structs that lists the DHCP options associated with the object. The description of the fields of `options` is as follows:
    * `name`: The Name of the DHCP option. Example: `domain-name-servers`.
    * `num`: The code of the DHCP option. Example: `6`.
    * `value`: The value of the option. Example: `11.22.33.44`.
    * `vendor_class`: The name of the space this DHCP option is associated to. Default value is `DHCP`.
    * `use_option`:Only applies to special options that are displayed separately from other options and have a use flag. These options are `router`,
      `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, `broadcast-address-offset`, `dhcp-lease-time`, and `dhcp6.name-servers`.

```terraform
options {
name         = "dhcp-lease-time"
value        = "43200"
vendor_class = "DHCP"
num          = 51
use_option   = false
}
```
* `cloud_info`: Structure containing all cloud API related information for this object. Example: `"{\"authority_type\":\"GM\",\"delegated_scope\":\"NONE\",\"owned_by_adaptor\":false}"`
* `use_options`: Use option is a flag that indicates whether the options field are used or not. The default value is false. Example: `false`

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `comment`, `ipv4addr` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| network_view | network_view | string | yes        |
| network      | network      | string | yes        |
| match_client | match_client | string | yes        |
| mac          | mac          | string | yes        |
| comment      | comment      | string | yes        |
| ipv4addr     | ipv4addr     | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
```hcl
data "infoblox_ipv4_fixed_address" "testFixedAddress_read1" {
  filters = {
    ipv4addr = "10.0.0.8"
  }
}
```
!> If `null` or empty filters are passed, then all the fixed address or objects associated with datasource like here `infoblox_ipv4_fixed_address` will be fetched in results.


### Example of an fixed address Data Source Block

This example defines a data source of type `infoblox_ipv4_fixed_address` and the name "fa_rec_temp1", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.
```hcl
resource "infoblox_ipv4_network" "net2" {
  cidr = "18.0.0.0/24"
}
resource "infoblox_ipv4_fixed_address" "fix4"{
  client_identifier_prepend_zero=true
  comment= "fixed address"
  dhcp_client_identifier="23"
  disable= true
  ext_attrs = jsonencode({
    "Site": "Blr"
  })
  match_client = "CLIENT_ID"
  name = "fixed_address_1"
  network = "18.0.0.0/24"
  network_view = "default"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = true
  }
  options {
    name = "routers"
    num = "3"
    use_option = true
    value = "18.0.0.2"
    vendor_class = "DHCP"
  }
  use_options = true
  depends_on=[infoblox_ipv4_network.net2]
}
data "infoblox_ipv4_fixed_address" "testFixedAddress_read1" {
  filters = {
    "*Site" = "Blr"
  }
  depends_on = [infoblox_ipv4_fixed_address.fix4]
}
output "fa_rec_temp1" {
  value = data.infoblox_ipv4_fixed_address.testFixedAddress_read1
}
```