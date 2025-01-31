# DTC Server Data Source

Use the `infoblox_dtc_server` data source to retrieve the following information for a DTC Server if any, which is managed by a NIOS server:

* `name`: The name of th DTC Server. Example: `test-server`.
* `auto_create_host_record`: Flag to enable the auto-creation of a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes. Example: `true`.
* `host`: The address or FQDN of the server. Example: `10.1.1.1`.
* `monitors`: List of IP/FQDN and monitor pairs to be used for additional monitoring. `monitors` has the following three fields `monitor_name`, `monitor_type` and `host`. The description of the fields of `monitors` is as follows:
  
  `monitor_name`: required, specifies the name of the monitor used for monitoring. Example: `https`.

  `monitor_type`: required, specifies the type of the monitor used for monitoring. Example: `https`.

  `host`: required, specifies the IP address or FQDN of the server used for monitoring. Example: `12.1.1.10`

```terraform
monitors {
    monitor_name = "https"
    monitor_type = "https"
    host = "12.12.1.1"
  }
```
* `disable`: Flag to determine whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled. Example: `true`.
* `sni_hostname`: The hostname for Server Name Indication (SNI) in FQDN format. Example: `test.com`.
* `use_sni_hostname`: Flag to enable the use of SNI hostname. Example: `true`.
* `health`: The DTC Server health information. The description of the fields of `health` is as follows:

  `availability`: The availability color status. Default value: `NONE`. Valid values are one of these: `BLUE`, `GREEN`, `GRAY`, `NONE`, `RED` and `YELLOW`.
  
  `description`: The textual description of the DTC Server object’s status. Default value: `""`. Example: `test health`.

  `enabled_state`: The enabled state of the DTC Server. Default value: `ENABLED`. Valid values are one of these: `DISABLED`, `DISABLED_BY_PARENT`, `ENABLED` and `NONE`.
```terraform
health { 
  availability = "NONE"
  description = ""
  enabled_state = "DISABLED"
}
```
* `comment`: The description of the DTC Server. This is a regular comment. Example: `test Dtc Server`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"*Site\":\"Antarctica\"}"`

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

This field used only for searching. This is not an actual member of the object and therefore the server does not return
this field with this name unless it is a nested return field.
-----
| Field   | Alias         | Type   | Searchable |
|---------|---------------|--------|------------|
| -       | status_member | string | yes        |

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

```hcl
data "infoblox_dtc_lbdn" "lbdn_filter" {   
  filters = {
    status_member = "infoblox.localdomain"
  }
}
```

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_dtc_server` will be fetched in results.

### Example of DTC Server Data Source Block

```hcl
resource "infoblox_dtc_server" "server_record" {
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

data "infoblox_dtc_server" "server_read" {
  filters = {
    name = infoblox_dtc_server.server_record.name
    comment = infoblox_dtc_server.server_record.comment
    host = infoblox_dtc_server.server_record.host
  }
}

output "server_res" {
  value = data.infoblox_dtc_server.server_read
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