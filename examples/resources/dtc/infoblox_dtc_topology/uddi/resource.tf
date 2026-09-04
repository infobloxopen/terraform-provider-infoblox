// Create a DTC Topology with a subnet source
resource "infoblox_dtc_topology" "example_basic" {
  uddi = {
    name = "example-topology-basic"
    sources = [
      {
        name    = "subnet-source"
        source  = "subnet"
        subnets = ["10.0.0.0/8"]
      }
    ]
  }
}

// Create a DTC Topology with multiple sources and optional fields
resource "infoblox_dtc_topology" "example_advanced" {
  uddi = {
    name     = "example-topology-advanced"
    comment  = "Topology with geographic routing"
    disabled = false
    sources = [
      {
        name    = "us-east"
        source  = "subnet"
        subnets = ["10.0.0.0/8", "192.168.0.0/16"]
      },
      {
        name    = "us-west"
        source  = "subnet"
        subnets = ["172.16.0.0/12"]
      }
    ]
    tags = {
      Site = "us-east-1"
    }
  }
}

// Create a DTC Topology with tag-rule-based sources
resource "infoblox_dtc_topology" "example_tag_rules" {
  uddi = {
    name = "example-topology-tag-rules"
    sources = [
      {
        name   = "production-source"
        source = "tag_rule"
        tag_rules = [
          {
            key   = "env"
            op    = "EQUALS"
            value = "production"
          },
          {
            key   = "region"
            op    = "NOT_EQUALS"
            value = "us-east-1"
          }
        ]
      }
    ]
  }
}
