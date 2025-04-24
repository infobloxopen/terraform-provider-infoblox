goog_cm_deployment_name = "nios-906-byol"


/* 
# ------------------------------------------------------------------------------------------------------
# NETWORK CONFIGURATION FOR MULTI-NIC DEPLOYMENTS
# ------------------------------------------------------------------------------------------------------

# List of VPC network names to be attached to each NIC of the VM.
# - The first value (e.g., "vnet1") will be used for NIC0 (typically used for management).
# - The second value (e.g., "vnet2") will be used for NIC1 (typically used for LAN access).
# networks = ["vnet1", "vnet2"]

# Add subnet names corresponding to each VPC listed above.
# - Example: "vnet1-subnet1" for NIC0, "vnet2-subnet2" for NIC1.
variable "sub_networks" {
  description = "The subnetwork names associated with each network above."
  type        = list(string)
  default     = ["vnet1-subnet1", "vnet2-subnet2"]
}

 
# -------------------------------------------------------------------------------------------------------
# EXTERNAL IP CONFIGURATION
# -------------------------------------------------------------------------------------------------------

# List of external IP types assigned to each NIC.
# - "NONE" means no external IP which is assigned to NIC0(usually for management).
# - "EPHEMERAL" will create a temporary external IP which is assigned to NIC1 (for GUI/web access).
# - First entry applies to NIC0, second entry applies to NIC1.
variable "external_ips" {
  description = "External IP assignment for each NIC. Use 'NONE' or 'EPHEMERAL'."
  type        = list(string)
  default     = ["NONE", "EPHEMERAL"]
}



# ---------------------------------------------------------------------------------------------------------
# OPTIONAL METADATA (can be uncommented if needed)
# ---------------------------------------------------------------------------------------------------------

# For "metadata" block this will apply temp_license for DNS, DHCP,grid, cloud and NIOS V926.

# variable "metadata" {
#   description = "(Optional) Key-value metadata pairs injected into the VM."
#   type        = map(string)
#   default     = {
#     user-data                  = "temp_License: nios IB-V926 enterprise dns dhcp cloud\ndefault_admin_password: Infoblox*123"
#     google-logging-enable      = "0"
#     google-monitoring-enable   = "0"
#   }
# }

*/