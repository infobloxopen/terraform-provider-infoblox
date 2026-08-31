// Create a DTC Topology with required fields only
resource "infoblox_dtc_topology" "example_basic" {
  nios = {
    name = "example-topology-basic"
  }
}

// Create a DTC Topology with a comment
resource "infoblox_dtc_topology" "example_with_comment" {
  nios = {
    name    = "example-topology-comment"
    comment = "DTC topology for geo-based routing"
  }
}

// Create a DTC Topology with routing rules (server destination)
resource "infoblox_dtc_topology" "example_with_rules" {
  nios = {
    name    = "example-topology-rules"
    comment = "Topology with geographic rules"
    rules = [
      {
        dest_type        = "SERVER"
        destination_link = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHNlcnZlcjE:server1"
        return_type      = "REGULAR"
        sources = [
          {
            source_type  = "COUNTRY"
            source_op    = "IS"
            source_value = "US"
          }
        ]
      },
      {
        # Default rule (no sources = catch-all)
        dest_type        = "SERVER"
        destination_link = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHNlcnZlcjI:server2"
        return_type      = "REGULAR"
      }
    ]
    ext_attrs = {
      Site = "us-east-1"
    }
  }
}
