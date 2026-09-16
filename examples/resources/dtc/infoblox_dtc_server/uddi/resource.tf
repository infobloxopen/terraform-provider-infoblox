// Create a DTC Server with an IP address endpoint
resource "infoblox_dtc_server" "example" {
  uddi = {
    name          = "dtc-server-basic"
    address       = "192.168.1.10"
    endpoint_type = "address"
  }
}

// Create a DTC Server with an FQDN endpoint
resource "infoblox_dtc_server" "example_fqdn_endpoint" {
  uddi = {
    name          = "dtc-server-fqdn"
    fqdn          = "server.example.com."
    endpoint_type = "fqdn"
    comment       = "DTC server resolved via FQDN"
  }
}

// Create a DTC Server with all optional fields
resource "infoblox_dtc_server" "example_all_fields" {
  uddi = {
    name                         = "dtc-server-full"
    address                      = "192.168.1.20"
    endpoint_type                = "address"
    comment                      = "Primary DTC server for web traffic"
    disabled                     = false
    auto_create_response_records = true
    tags = {
      env  = "production"
      Site = "us-east-1"
    }
  }
}

// Create a DTC Server with associated DNS records
resource "infoblox_dtc_server" "example_with_records" {
  uddi = {
    name          = "dtc-server-with-records"
    address       = "192.168.1.30"
    endpoint_type = "address"
    records = [
      {
        type  = "A"
        rdata = { address = "192.168.1.30" }
      },
      {
        type  = "AAAA"
        rdata = { address = "2001:db8::1" }
      },
    ]
  }
}
