# Ipv4 Shared Network Data Source

The `infoblox_ipv4_shared_network` resource allows you to create SharedNetwork-record on NIOS side,
The following list describes the parameters you can define for the `infoblox_ipv4_shared_network` resource block:

* `name`: required, specifies the name of the IPv4 shared network object. Example: `shared-network1`
* `networks`: required, specifies the list of networks belonging to the shared network. Example: `["11.11.1.0/24", "12.12.1.0/24"]`
* `network_view`: optional, specifies the name of the network view in which this shared network resides. Example: `view2`. Default value is `default`.
* `disable`: optional, specifies the disable flag for the IPv4 shared network object. Example: `true`. Default value is `false`.
* `use_options`: optional, specifies the use flag for options. Example: `true`. Default value is `false`.
* `options`: optional, specifies an array of DHCP option structs that lists the DHCP options associated with the object. The description of the fields of `options` is as follows:
    * `name`: required, specifies the Name of the DHCP option. Example: `domain-name-servers`.
    * `num`: required, specifies the code of the DHCP option. Example: `6`.
    * `value`: required, specifies the value of the option. Example: `11.22.33.44`.
    * `vendor_class`: optional, specifies the name of the space this DHCP option is associated to. Default value is `DHCP`.
    * `use_option`: optional, only applies to special options that are displayed separately from other options and have a use flag. These options are `router`, 
  `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, `broadcast-address-offset`, `dhcp-lease-time`, and `dhcp6.name-servers`.

Example for options field:
```terraform
options { 
    name = "domain-name-servers"
    use_option = true
    value = "11.22.33.44"
  }
```
Default value for options is:
```terraform
options { 
    name = "dhcp-lease-time"
    num = 51
    use_option = false  
    value = "43200"
    vendor_class = "DHCP"
  }
```
* `comment`: optional, specifies the description of the record. This is a regular comment. Example: `Temporary Ipv4 Shared Network`.
* `ext_attrs`: optional, specifies the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":"Vancouver"}"`

### Example of an Ipv4 Shared Network Resource Block:
 ```hcl
// shared network with minimum set of parameters
resource "infoblox_ipv4_shared_network" "shared_network_min_parameters" {
  name = "shared-network1"
  networks = ["37.12.3.0/24"]
}

// shared network with full set of parameters
resource "infoblox_ipv4_shared_network" "shared_network_full_parameters" {
  name = "shared-network2"
  comment = "test ipv4 shared network record"
  networks = ["31.12.3.0/24","31.13.3.0/24"]
  network_view = "view2"
  disable = false
  ext_attrs = jsonencode({
    "Site" = "Tokyo"
  })
  use_options = false
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
  options {
    name = "dhcp-lease-time"
    num = 51
    use_option = false  
    value = "43200"
    vendor_class = "DHCP"
  }
}
 ```