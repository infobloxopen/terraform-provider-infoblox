# Integration of AWS resources with Infoblox Provider

### Prerequisites:
```
export INFOBLOX_PASSWORD="${password}"
export INFOBLOX_SERVER="${server}"
export INFOBLOX_USERNAME="${username}"

Install AWS CLI. Use AWS Configure command to configure Access Key ID and Secret Access Key.
```

#Using the templates for below use cases.

- FixedAddress: Create AWS instances with IP being managed from Infoblox NIOS grid
- ARecord     : Create AWS instances with IP being managed from Infoblox NIOS grid and DNS A record creation automation
- PTRRecord   : Create AWS instances with IP being managed from Infoblox NIOS grid and DNS PTR record creation automation
- CNAMERecord : Create AWS instances with IP being managed from Infoblox NIOS grid and DNS CNAME record creation automation
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
- terraform state
- terraform state list
- terraform state show <item>
# Destroying the Resource
- terraform destroy
