# Integration of azure Provider with Infoblox Provider

### Prerequisites
```
export subscription_id="${subscription_id}"
export client_id="${client_id}"
export client_secret="${client_secret}"
export tenant_id="${tenant_id}"
export INFOBLOX_PASSWORD="${password}"
export INFOBLOX_SERVER="${server}"
export INFOBLOX_USERNAME="${username}"
```
# Running the Resource

- terraform [init](https://www.terraform.io/docs/commands/init.html)
- terraform plan
- terraform apply

#Disclaimer

 Once the resource is created, run terraform apply again to update NIOS Appliance with vm properties

# Destroying the Resource
 terraform destroy
