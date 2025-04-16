goog_cm_deployment_name = "nios-906-byol"


/* 
# For multiNIC deployments
variable "networks" {
description = "The network name to attach the VM instance."
  type        = list(string)
  default     = ["vnet1","vnet2"]
}
#add vnets as per NIC 

variable "sub_networks" {
  description = "The sub network name to attach the VM instance."
  type        = list(string)
  default     = ["vnet1-subnet1","vnet2-subnet2"]
}
#Add subnets as per vnets
 
variable "external_ips" {
  description = "The external IPs assigned to the VM for public access."
  type        = list(string)
  default     = ["NONE","EPHEMERAL"]
}
# For externalip: "NONE" is assigned to NIC0(mgmt) 
# "EPHEMERAL is assigned to NIC1 (used for LAN accessing GUI)


For "metadata" block this will apply temp_license for DNS, DHCP,grid, cloud and NIOS V926
#variable "metadata" {
#  description = "(Optional) Metadata key/value pairs to make available from within the VM instance."
#  type        = map(string)
#  default     = {
#      user-data = "temp_license: nios IB-V926 enterprise dns dhcp cloud\ndefault_admin_password: Infoblox*123"
#      google-logging-enable = "0"
#      google-monitoring-enable = "0"
# }
#}

*/