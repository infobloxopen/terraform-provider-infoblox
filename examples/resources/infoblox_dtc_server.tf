// creating a DTC server record with minimal set of parameters
resource "infoblox_dtc_server" "server1" {
  name = "server1"
  host = "12.12.1.1"
}

// creating a DTC Server record with full set of parameters
resource "infoblox_dtc_server" "server2" {
  name                    = "server2"
  host                    = "11.11.1.1"
  comment                 = "test DTC server"
  auto_create_host_record = false
  use_sni_hostname        = true
  sni_hostname            = "test.com"
  disable                 = true

  monitors {
    monitor_name = "https"
    monitor_type = "http"
    host         = "22.21.1.2"
  }

  ext_attrs = jsonencode({
    "Site" = "CA"
  })
}
