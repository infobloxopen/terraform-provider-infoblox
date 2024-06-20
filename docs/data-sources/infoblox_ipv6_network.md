# IPv6 Network Data Source

The data source for the network object allows you to get the following parameters for an IPv6 network resource:

* `network_view`: the network view which the network container exists in. Example: `nondefault_netview`
* `cidr`: the network block which corresponds to the network, in CIDR notation. Example: `2002:1f93:0:4::/96`
* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: The set of extensible attributes, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\",\"Administrator\":\"unknown\"}"`.


For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `network`, `network_view` corresponding to object.
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
data "infoblox_ipv6_network" "readNet1" {
  filters = {
    network = "2002:1f93:0:4::/96"
    network_view = "nondefault_netview"
  }
  depends_on = [infoblox_ipv6_network.ipv6net1]
}
 ```

!> From the above example, if the 'network_view' value is not specified, if same network exists in one or more different network views, those
all networks will be fetched in results.

!> If `null` or empty filters are passed, then all the networks or objects associated with datasource like here `infoblox_ipv6_network`, will be fetched in results.

### Example of a Network Data Source Block

```hcl
// This is just to ensure that the network has been be created
resource "infoblox_ipv6_network" "ipv6net1" {
  cidr = "2002:1f93:0:4::/96"
  reserve_ipv6 = 10
  gateway = "2002:1f93:0:4::1"
  comment = "let's try IPv6"
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

data "infoblox_ipv6_network" "readNet1" {
  filters = {
    network = "2002:1f93:0:4::/96"
  }
  // using 'infoblox_ipv6_network' resource block before the data source will be queried.
  depends_on = [infoblox_ipv6_network.ipv6net1]
}

// accessing IPv6 network through EA's
data "infoblox_ipv6_network" "readnet2" {
  filters = {
    "*Site" = "Antarctica"
  }
  depends_on = [infoblox_ipv6_network.ipv6net1]
}

// throws matching IPv6 network.
output "ipv6net_res" {
  value = data.infoblox_ipv6_network.readNet1
}

// throws matching IPv4 networks with EA, if any
output "ipv6net_res1" {
  value = data.infoblox_ipv6_network.readnet2
}
```
