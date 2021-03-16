# Integration of vSphere with Infoblox Provider

### Prerequisites
```
export VSPHERE_USER="${user_name}"
export VSPHERE_PASSWORD="${password}"
export VSPHERE_SERVER="${server}"
export VSPHERE_ALLOW_UNVERIFIED_SSL=true
export INFOBLOX_PASSWORD="${password}"
export INFOBLOX_SERVER="${server}"
export INFOBLOX_USERNAME="${username}"
```

#Using the templates for different combinations.

- To get next available network from a given parent CIDR of a prefix length use 
  templates from NextAvailableNetwork.

### Note
```
A parent network container has to be in existence before requesting next available network from it.
```

# Running the Resource

- terraform [init](https://www.terraform.io/docs/commands/init.html)
- terraform plan
- terraform apply


# Destroying the Resource
- terraform destroy

