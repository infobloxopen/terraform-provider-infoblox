# Fixed Address Resource

A fixed address is a specific IP address that a DHCP server always assigns when a lease request comes from a particular MAC address of the client.

The `infoblox_ipv4_fixed_address` resource, enables you to allocate, update, or delete an fixed address within a network in a NIOS appliance.

* `agent_circuit_id`: optional, The agent circuit ID for the fixed address. The field is required only when match_client is set to CIRCUIT_ID. Example: `23`
* `agent_remote_id`: optional, The agent remote ID for the fixed address. The field is required only when match_client is set to REMOTE_ID. Example: `34`
* `client_identifier_prepend_zero`: optional, This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address. Example: `false`
* `comment`: optional, Comment for the fixed address; maximum 256 characters. Example: `fixed address`
* `dhcp_client_identifier`: optional, The DHCP client ID for the fixed address. The field is required only when match_client is set to CLIENT_ID. Example: `20`
* `disable`: optional, Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled. Example: `false`
* `ext_attrs`: optional, Extensible attributes associated with the object. Example: `"{\"*Site\":\"Antarctica\"}"`
* `ipv4addr`: optional, The IPv4 Address of the fixed address. If the `ipv4addr` field is not provided and the `network` field is set, the next available IP address in the network will be allocated. Example: `10.0.0.34`
* `mac`: optional, The MAC address value for this fixed address. The field is required only when match_client is set to its default value - MAC_ADDRESS. Example: `00-1A-2B-3C-4D-5E`
* `match_client`: optional, The match client for the fixed address.Valid values are CIRCUIT_ID, CLIENT_ID , MAC_ADDRESS, REMOTE_ID and RESERVED. Default value is MAC_ADDRESS. Example: `CLIENT_ID`
* `name`: optional, This field contains the name of this fixed address. Example: `fixedAddressName`
* `network`: optional, The network to which this fixed address belongs, in IPv4 Address/CIDR format. Example: `10.0.0.0/24`
* `network_view`: optional, The name of the network view in which this fixed address resides. The default value is The default network view. Example: `default`
* `options`: optional, specifies an array of DHCP option structs that lists the DHCP options associated with the object. The description of the fields of `options` is as follows:
    * `name`: required, specifies the Name of the DHCP option. Example: `domain-name-servers`.
    * `num`: required, specifies the code of the DHCP option. Example: `6`.
    * `value`: required, specifies the value of the option. Example: `11.22.33.44`.
    * `vendor_class`: optional, specifies the name of the space this DHCP option is associated to. Default value is `DHCP`.
    * `use_option`: optional, only applies to special options that are displayed separately from other options and have a use flag. These options are `router`,
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
* `use_options`: optional, Use option is a flag that indicates whether the options field are used or not. The default value is false. Example: `false`

## Example for Fixed Address Block 

```hcl
//example for fixed address with maximal parameters and using next available ip function 
//ipv4addr not specified and network is given so next available ip in the network will be allocated
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
    use_option   = false  
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
resource "infoblox_ipv4_network" "net2" {
  cidr = "18.0.0.0/24"
}
//creates a fixed address by explicitly providing the `ipv4addr` value instead of using the next available IP function.
resource "infoblox_ipv4_fixed_address" "fix3"{
  ipv4addr        = "17.0.0.9"
  mac = "00:0C:24:2E:8F:2A"
  options {
    name         = "dhcp-lease-time"
    value        = "43200"
    vendor_class = "DHCP"
    num          = 51
    use_option   = false  
  }
  depends_on=[infoblox_ipv4_network.net3]
}
resource "infoblox_ipv4_network" "net3" {
  cidr = "17.0.0.0/24"
}
```