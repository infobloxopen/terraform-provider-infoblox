# IPv6 Network Data Source

The data source for the network object allows you to get the following parameters for an IPv6 network resource:

* `network_view`: the network view which the network container exists in. Example: `nondefault_netview`
* `cidr`: the network block which corresponds to the network, in CIDR notation. Example: `2002:1f93:0:4::/96`
* `comment`: a description of the network. This is a regular comment. Example: `Untrusted network`.
* `ext_attrs`: The set of extensible attributes, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\",\"Administrator\":\"unknown\"}"`.


To retrieve information about IPv6 network that match the specified filters, use the `filters` argument and specify the parameters mentioned in the below table. These are the searchable parameters of the corresponding object in Infoblox NIOS WAPI. If you do not specify any parameter, the data source retrieves information about all host records in the NIOS Grid.

The following table describes the parameters you can define in an `infoblox_ipv6_network` data source block:

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| network      | cidr         | string | yes        |
| network_view | network_view | string | yes        |
| comment      | comment      | string | yes        |

!> Aliases are the parameter names used in the prior releases of Infoblox IPAM Plug-In for Terraform. Do not use the alias names for parameters in the data source blocks. Using them can result in error scenarios.

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
