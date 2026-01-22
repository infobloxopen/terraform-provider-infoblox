resource "infoblox_vlan" "vlan100" {
  name        = "production-vlan"
  vlan_id     = 100
  comment     = "Production environment VLAN"
  description = "VLAN for production workloads"
  department  = "IT Operations"
  contact     = "ops-team@example.com"
  ext_attrs = jsonencode({
    "Site"        = "Datacenter 1"
    "Environment" = "Production"
    "Tenant ID"   = "terraform_test_tenant"
  })
}

# Minimal example
resource "infoblox_vlan" "simple_vlan" {
  name    = "my-vlan"
  vlan_id = 200
}

# Example with multiple attributes
resource "infoblox_vlan" "dev_vlan" {
  name        = "dev-vlan"
  vlan_id     = 150
  comment     = "Development VLAN"
  description = "VLAN for development environment"
  department  = "Engineering"
  contact     = "dev-team@example.com"
}
