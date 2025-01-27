# DTC Server Data Source

Use the `infoblox_dtc_server` data source to retrieve the following information for a DTC Server if any, which is managed by a NIOS server:

* `name`: The name of th DTC Server. Example: `test-server`.
* `auto_create_host_record`: Flag to enable the auto-creation of a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes.
* `host`: The address or FQDN of the server.
* `monitors`: List of IP/FQDN and monitor pairs to be used for additional monitoring.
* `disable`: Flag to determine whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled.
* `sni_hostname`: The hostname for Server Name Indication (SNI) in FQDN format.
* `use_sni_hostname`: Flag to enable the use of SNI hostname.
* `health`: The DTC Server health information.
* `comment`: The description of the DTC Server. This is a regular comment. Example `this is some text`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":\"Kapu\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `comment`, `host`, `sni_hostname` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field        | Alias        | Type   | Searchable |
|--------------|--------------|--------|------------|
| name         | name         | string | yes        |
| comment      | comment      | string | yes        |
| host         | host         | string | yes        |
| sni_hostname | sni_hostname | string | yes        |


!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_dtc_server" "server_filter" {
    filters = {
        name = "test-server"
        comment = "sample Server"
    }
 }
 ```

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_dtc_server` will be fetched in results.

### Example of DTC Server Data Source Block

```hcl
resource "infoblox_dtc_server" "server" {
  name = "server2"
  host = "11.11.1.1"
  comment = "test DTC server"
  ext_attrs = jsonencode({
    "Site" = "CA"
  })
  disable = true
  auto_create_host_record = false
  use_sni_hostname = true
  sni_hostname = "test.com"
  monitors {
    monitor_name = "https"
    monitor_type = "http"
    host = "22.21.1.2"
  }
}

data "infoblox_dtc_server" "server" {
  filters = {
    name = "testserver"
    comment = "test Server"
    host = "11.11.1.1"
  }
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_dtc_server' resource block before the data source will be queried.
  depends_on = [infoblox_dtc_server.server]
}

output "server_res" {
  value = data.infoblox_dtc_server.server
}

// accessing individual field in results
output "server_name" {
  value = data.infoblox_dtc_server.server_res.results.0.name //zero represents index of json object from results list
}

// accessing DTC Server through EA's
data "infoblox_dtc_server" "server_ea" {
  filters = {
    "*Site" = "CA"
  }
}

output "server_out" {
  value = data.infoblox_dtc_server.server_ea
}
```