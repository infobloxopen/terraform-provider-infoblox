# Network View Data Source

Use the data source to retrieve the following information for a network view resource from the corresponding object in NIOS:

* `name`: the name of the network view to be specified. Example: `custom_netview`
* `comment`: a description of the network view. This is a regular comment. Example: `From the outside`.
* `ext_attrs`: the set of extensible attributes of the network view, if any. The content is formatted string of JSON map. Example: `"{\"Administrator\":\"jsw@telecom.ca\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----
| Field   | Alias   | Type   | Searchable |
|---------|---------|--------|------------|
| name    | name    | string | yes        |
| comment | comment | string | yes        |

!> Either you can fetch with both `name` and `comment` or just with `name` field.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_network_view" "nview_filter" {
    filters = {
        name = "nondefault_netview"
    }
 }
 ```

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_network_view` will be fetched in results.

### Example of a Network View Data Source Block

```hcl
resource "infoblox_network_view" "inet_nv" {
  name = "inet_visible_nv"
  comment = "Internet-facing networks"
  ext_attrs = jsonencode({
    "Location" = "the North pole"
  })
}

data "infoblox_network_view" "inet_nv" {
  filters = {
    name = "inet_visible_nv"
  }

  // This is just to ensure that the network view has been be created
  // using 'infoblox_network_view' resource block before the data source will be queried.
  depends_on = [infoblox_network_view.inet_nv]
}

output "nview_res" {
  value = data.infoblox_network_view.inet_nv
}

// accessing individual field in results
output "nview_name" {
  value = data.infoblox_network_view.inet_nv.results.0.name //zero represents index of json object from results list
}

// accessing IPv4 network through EA's
data "infoblox_network_view" "nview_ea" {
  filters = {
    "*Administrator" = "jsw@telecom.ca"
  }
}

// throws matching Network Views with EA, if any
output "nview_ea_out" {
  value = data.infoblox_network_view.nview_ea
}
```
