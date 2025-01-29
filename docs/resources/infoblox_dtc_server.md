# DTC Server Resource

The `infoblox_dtc_server` resource enables you to perform `create`, `update` and `delete` operations on DTC Server in a NIOS appliance.
The resource represents the ‘dtc:server’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the DTC Server object:

* `name`: required, specifies the display name of the DTC Server. Example: `test-server`.
* `auto_create_host_record`: optional, specifies the flag to enable the auto-creation of a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes. Default value: `true`.
* `host`: required, specifies the address or FQDN of the server. Example: `11.1.1.2`.
* `disable`: optional, specifies whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled. Default value: `false`.
* `sni_hostname`: optional, specifies the hostname for Server Name Indication (SNI) in FQDN format. Example: `test.example.com`.
* `use_sni_hostname`: optional, specifies the flag to enable the use of SNI hostname. Default value: `false`.
* `comment`: optional, description of the DTC Server. Example: `custom DTC Server`.
* `ext_attrs`: optional, set of the Extensible attributes of the Server, as a map in JSON format. Example: `jsonencode({\"Site\":\"Kapu\"})`.
* `monitors`: optional, specifies the List of IP/FQDN and monitor pairs to be used for additional monitoring. `monitors` has the following three fields `monitor_name`, `monitor_type` and `host`. Example:
```terraform
monitors {
    monitor_name = "https"
    monitor_type = "https"
    host = "12.12.1.1"
  }
```
* `monitor_name`: required, specifies the name of the monitor used for monitoring. Example: `https`.
* `monitor_type`: required, specifies the type of the monitor used for monitoring. Example: `https`.
* `host`: required, specifies the IP address or FQDN of the server used for monitoring. Example: `12.1.1.10`

### Examples of a DTC Server Block

```hcl
// creating a DTC server record with minimal set of parameters
resource "infoblox_dtc_server" "server1" {
  name = "server1"
  host = "12.12.1.1"
}

// creating a DTC Server record with full set of parameters
resource "infoblox_dtc_server" "server2" {
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
```

