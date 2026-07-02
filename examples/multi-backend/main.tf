terraform {
  required_providers {
    unified = {
      source  = "infobloxopen/unified"
      version = "0.0.1"
    }
  }
}

# Default provider - NIOS (no alias mandatory)
# Resources without explicit "provider" attribute will use this
provider "unified" {
  nios {
    host_url = "https://172.28.83.238"
    username = "admin"
    password = "Infoblox@123"
  }
}

# Aliased provider - UDDI
# Resources must explicitly specify: provider = unified.uddi
provider "unified" {
  alias = "uddi"
  uddi {
    csp_url = "https://stage.csp.infoblox.com"
    api_key = "4a815e1e1c86a208efab3a5bfcc6f1f73259c009c43d30b13b337786ca9b3328"
  }
}

resource "unified_dns_record_a" "nios_test" {
  # No provider attribute = uses default (NIOS)
  name    = "test-rec-23.example.com"
  ipv4    = "10.0.0.1"
  comment = "This is a test A record nios"

  nios = {
    creator = "DYNAMIC"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

resource "unified_dns_record_a" "uddi_test" {
  provider = unified.uddi # Required for UDDI resources

  name    = "test-rec-23"
  ipv4    = "10.0.0.1"
  zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9" // example.com
  comment = "This is a test A record uddi"

  uddi = {
    tags = {
      Site = "location-1"
    }
  }
}
