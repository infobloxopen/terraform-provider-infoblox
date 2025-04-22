# IPv4 Network Data Source

The data source for the network object allows you to get the following parameters for an IPv4 network resource:

* `network_view`: the network view which the network container exists in. Example: `nondefault_netview`
* `cidr`: the network block which corresponds to the network, in CIDR notation. Example: `192.0.17.0/24`
* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: The set of extensible attributes, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\",\"Administrator\":\"unknown\"}"`.
* `options`: An array of DHCP option structs that lists the DHCP options associated with the object.
```terraform
options {
  name         = "dhcp-lease-time"
  value        = "43200"
  vendor_class = "DHCP"
  num          = 51
  use_option   = true
}
```
* `utilization`: The network utilization in percentage. Example: `0`

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| network      | cidr         | string | yes        |
| network_view | network_view | string | yes        |
| comment      | comment      | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_ipv4_network" "network_filter" {
    filters = {
        network = "10.11.0.0/16"
        network_view = "nondefault_netview"
    }
 }
 ```

!> From the above example, if the 'network_view' value is not specified, if same network exists in one or more different network views, those
all networks will be fetched in results.

!> If `null` or empty filters are passed, then all the networks or objects associated with datasource like here `infoblox_ipv4_network`, will be fetched in results.

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
  filters = {
    network = "192.168.128.0/20"
    network_view = "nondefault_netview"
  }
  // This is just to ensure that the network has been be created
  // using 'infoblox_ipv4_network' resource block before the data source will be queried.
  depends_on = [infoblox_ipv4_network.net2]
}

output "ipv4_net1" {
  value = data.infoblox_ipv4_network.nearby_network
}

// accessing individual field in results
output "ipv4_net2" {
  value = data.infoblox_ipv4_network.nearby_network.results.0.cidr //zero represents index of json object from results list
}

// accessing IPv4 network through EA's
data "infoblox_ipv4_network" "ipv4_net_ea" {
  filters = {
    "*Site" = "Custom network site"
  }
}

// throws matching IPv4 networks with EA, if any
output "net_ea_out" {
  value = data.infoblox_ipv4_network.ipv4_net_ea
}
```
