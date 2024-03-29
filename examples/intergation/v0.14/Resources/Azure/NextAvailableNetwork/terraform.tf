terraform {
  # Required providers block for Terraform v0.14
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.50.0"
    }
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}
