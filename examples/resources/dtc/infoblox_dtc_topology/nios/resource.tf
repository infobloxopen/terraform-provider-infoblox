// Create DTC Servers (required as rule destinations)
resource "infoblox_dtc_server" "example_server_us" {
  nios = {
    name = "example-server-us"
    host = "2.2.2.2"
  }
}

// Create DTC Servers (required as rule destinations)
resource "infoblox_dtc_server" "example_server_default" {
  nios = {
    name = "example-server-default"
    host = "3.3.3.3"
  }
}
// Create a DTC Topology with required fields only
resource "infoblox_dtc_topology" "example_basic" {
  nios = {
    name = "example-topology-basic_1"
  }
}

// Create a DTC Topology with routing rules (server destination)
resource "infoblox_dtc_topology" "example_with_rules" {
  nios = {
    name    = "example-topology-rules_2"
    comment = "Topology with geographic rules"
    rules = [
      {
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.example_server_default.id
        return_type      = "REGULAR"
      },
      {
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.example_server_us.id
        return_type      = "REGULAR"
        sources = [
          {
            source_op    = "IS"
            source_type  = "SUBNET"
            source_value = "10.0.0.0/8"
          },
        ]
      },
    ]
    ext_attrs = {
      Site = "us-east-1"
    }
  }
  depends_on = [infoblox_dtc_server.example_server_default, infoblox_dtc_server.example_server_us]
}
