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

- For creation of Fixed Address or if you're just using the provider without dns,
  follow the templates used in fixed address folder
- For creation of Host Record use the example in Host Record Folder.
  The enableDns flag in the templates decides if it has to be used or not 
  for dns purposes
- Provisioning of Multiple VM's and Creation of A Record's for those multiple
  VM's is shown in Multiple Folder
- Example of other records such as A,PTR and CNAME are shown in 
  infoblox.tf of Fixed Adrress folder
- To get next available network from a given parent CIDR of a prefix length use 
  templates from NextAvailableNetwork.

### Note
```
Need to create forward-mapping and reverse-mapping zones manually for creation of DNS records in DNS View.
A parent network container has to be in existence before requesting next available network from it.
```

# Running the Resource

- terraform [init](https://www.terraform.io/docs/commands/init.html)
- terraform plan
- terraform apply


# Destroying the Resource
- terraform destroy
