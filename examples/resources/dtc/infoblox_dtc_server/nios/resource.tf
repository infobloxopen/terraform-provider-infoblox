// Create a DTC Server with required fields only
resource "infoblox_dtc_server" "example" {
  nios = {
    name = "dtc-server-basic"
    host = "192.168.1.10"
  }
}

// Create a DTC Server with all optional fields
resource "infoblox_dtc_server" "example_all_fields" {
  nios = {
    name                    = "dtc-server-full"
    host                    = "192.168.1.20"
    comment                 = "Primary DTC server for web traffic"
    disable                 = false
    auto_create_host_record = true
    sni_hostname            = "server.example.com"
    ext_attrs = {
      Site = "us-east-1"
    }
  }
}

// Create a DTC Server with health monitors attached
resource "infoblox_dtc_server" "example_with_monitors" {
  nios = {
    name = "dtc-server-monitored"
    host = "192.168.1.30"
    monitors = [
      {
        host    = "192.168.1.30"
        monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      },
      {
        host    = "192.168.1.31"
        monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"
      },
    ]
  }
}
