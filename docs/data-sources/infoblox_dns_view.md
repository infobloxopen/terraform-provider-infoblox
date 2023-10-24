# DNS View Data Source

Use the `infoblox_dns_view` data source to retrieve the following information for a DNS View if any, which is managed by a NIOS server:

* `name`: The name of th DNS View. Example: `custom_dnsview`.
* `network_view`: The name of the network view object associated with this DNS view. Example: `nondefault_netview`.
* `comment`: The description of the DNS View. This is a regular comment. Example `this is some text`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\", \"Expires\":\"never\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----
| Field        | Alias | Type   | Searchable |
|--------------|-------|--------|------------|
| name         | ---   | string | yes        |
| network_view | ---   | string | yes        |
| comment      | ---   | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_dns_view" "view_filter" {
    filters = {
        name = "custom_dnsview"
        network_view = "default"
    }
 }
 ```

!> If `null` or empty filters are passed, then all the views or objects associated with datasource like here `infoblox_dns_view` will be fetched in results.

### Example of DNS View Data Source Block

```hcl
resource "infoblox_dns_view" "view1" {
  name = "customview"
  network_view = "nondefault_netview"
  comment = "sample custom dns view"
  ext_attrs = jsonencode({
    "Location" = "NewYork"
  })
}

data "infoblox_dns_view" "dsview" {
  filters = {
    name = "customview"
    network_view = "nondefault_netview"
  }
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_dns_view' resource block before the data source will be queried.
  depends_on = [infoblox_dns_view.view1]
}

output "dsview_res" {
  value = data.infoblox_dns_view.dsview
}

// accessing individual field in results
output "dsview_name" {
  value = data.infoblox_dns_view.dsview_res.results.0.name //zero represents index of json object from results list
}

// accessing DNS Views through EA's
data "infoblox_dns_view" "dsview_ea" {
  filters = {
    "*TestEA" = "SampleEA"
  }
}

output "dsview_out" {
  value = data.infoblox_dns_view.dsview_ea
}
```