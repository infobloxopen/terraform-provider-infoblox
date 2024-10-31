# IPv4 Network Container Data Source

Use the data source to retrieve the following information for an IPv4 network container resource from the corresponding
object in NIOS:
* `network_view`: the network view which the network container exists in. Example: `nondefault_netview`
* `cidr`: the IPv4 network block of the network container. Example: `19.17.0.0/16`
* `comment`: a description of the network container. This is a regular comment. Example: `Tenant 1 network container`.
* `ext_attrs`: the set of extensible attributes of the network view, if any. The content is formatted as stirng of JSON map. Example: `"{\"Administrator\":\"jsw@telecom.ca\"}"`.

As there is new feature filters , the previous usage of combination of Network view and address of the network block in CIDR format has been removed.

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
 data "infoblox_ipv4_network_container" "nc_filter" {
    filters = {
        network = "10.11.0.0/16"
        network_view = "nondefault_netview"
    }
 }
 ```

!> From the above example, if the 'network_view' value is not specified, if same network container exists in one or more different network views, those
all network containers will be fetched in results.

!> If `null` or empty filters are passed, then all the network containers or objects associated with datasource like here `infoblox_ipv4_network_container`, will be fetched in results.

### Example of an IPv4 Network Container Data Source Block

```hcl
resource "infoblox_ipv4_network_container" "nearby_org" {
  cidr = "192.168.128.0/17"
  network_view = "separate_tenants"
  comment = "one of our clients"
  ext_attrs = jsonencode({
    "Site" = "remote office"
    "Country" = "Australia"
  })
}

data "infoblox_ipv4_network_container" "nearby_nc" {
  filters = {
    network_view = "separate_tenants"
    network = "192.168.128.0/17"
  }

  # This is just to ensure that the network container has been be created
  # using 'infoblox_ipv4_network_container' resource block before the data source will be queried.
  depends_on = [infoblox_ipv4_network_container.nearby_org]
}

output "nc_res" {
  value = data.infoblox_ipv4_network_container.nearby_nc
}

# Accessing individual field in results
output "nc_cidr_out" {
  value = data.infoblox_ipv4_network_container.nearby_nc.results.0.cidr # zero represents index of json object from results list
}

# Accessing IPv4 Network Container through EA's
data "infoblox_ipv4_network_container" "nc_ea" {
  filters = {
    "*Site" = "GMC Site"
  }
}

# Throws matching IPv4 Network Containers with EA, if any
output "nc_ea_out" {
  value = data.infoblox_ipv4_network_container.nc_ea
}
```
