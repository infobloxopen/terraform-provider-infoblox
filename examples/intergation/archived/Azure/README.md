# Integration of azure Provider with Infoblox Provider

### Prerequisites
```
Install Azure CLI and configure subscription ID, Client ID and Secret.

export INFOBLOX_PASSWORD="${password}"
export INFOBLOX_SERVER="${server}"
export INFOBLOX_USERNAME="${username}"
```

#Using the templates for below use cases.
- NextAvailableNetwork : Get next available network from a given parent CIDR of a prefix length.

### Note
```
Need to create forward-mapping and reverse-mapping zones manually for creation of DNS records in DNS View.
A parent network container has to be in existence before requesting next available network from it.
```

# Running the Resource

- terraform [init](https://www.terraform.io/docs/commands/init.html)
- terraform plan
- terraform apply

#Disclaimer

 Once the resource is created, run terraform apply again to update NIOS Appliance with vm properties

# Destroying the Resource
 terraform destroy
