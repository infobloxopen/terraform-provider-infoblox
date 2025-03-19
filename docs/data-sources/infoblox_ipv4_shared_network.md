# Ipv4 Shared Network Data Source

Use the `infoblox_ipv4_shared_network` data source to retrieve the following information for an Ipv4 Shared Network if any, which is managed by a NIOS server:

* `name`: The name of the IPv4 shared network object. Example: `shared-network1`
* `networks`: The list of networks belonging to the shared network. Example: `["11.11.1.0/24"]`
* `network_view`: The name of the network view in which this shared network resides. Example: `default`
* `disable`: The disable flag for the IPv4 shared network object. Example: `true`
* `use_options`: Use flag for options. Example: `true`.
* `options`: An array of DHCP option structs that lists the DHCP options associated with the object. Example:
```terraform
option { 
    name = "domain-name-servers"
    value = "11.22.33.44"
    use_option = true
  }
```
* `comment`: The description of the record. This is a regular comment. Example: `Temporary Ipv4 Shared Network`.
* `ext_attrs`: The set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":"Vancouver"}"`

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `network_view` and `comment` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| name         | fqdn         | string | yes        |
| network_view | network_view | string | yes        |
| comment      | zone         | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_ipv4_shared_network" "shared_network_filter" {
    filters = {
        name = "shared-network1"
        network_view = "default" // associated Network view
    }
 }
 ```

!> From the above example, if the 'network_view' value is not specified, if same record exists in one or more different network views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_ipv4_shared_network` will be fetched in results.

### Example of an Ipv4 Shared Network Data Source Block

This example defines a data source of type `infoblox_ipv4_shared_network` and the name "shared_network_read", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_ipv4_shared_network" "shared_network" {
  name = "shared-network"
  comment = "test ipv4 shared network record"
  networks = ["31.12.3.0/24","31.13.3.0/24"]
  network_view = "default"
  disable = false
  ext_attrs = jsonencode({
    "Site" = "Yokohama"
  })
  use_options = false
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
}


data "infoblox_ipv4_shared_network" "shared_network_read" {
  filters = {
    name = "shared-network"
    network_view = "default"
  }
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_ipv4_shared_network' resource block before the data source will be queried.
  depends_on = [infoblox_ipv4_shared_network.shared_network]
}

output "shared_network_res" {
  value = data.infoblox_ipv4_shared_network.shared_network_read
}

// accessing individual field in results
output "shared_network_name" {
  value = data.infoblox_ipv4_shared_network.shared_network_read.results.0.name //zero represents index of json object from results list
}

// accessing Ipv4 Shared Network through EA's
data "infoblox_ipv4_shared_network" "shared_network_ea" {
  filters = {
    "*Site" = "Yokohama"
  }
}

output "shared_network_ea_res" {
  value = data.infoblox_ipv4_shared_network.shared_network_ea
}
```

