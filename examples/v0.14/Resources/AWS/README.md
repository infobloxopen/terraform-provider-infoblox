# Integration of AWS resources with Infoblox Provider

### Prerequisites:
```
export INFOBLOX_PASSWORD="${password}"
export INFOBLOX_SERVER="${server}"
export INFOBLOX_USERNAME="${username}"

Install AWS CLI. Use AWS Configure command to configure Access Key ID and Secret Access Key.
```

#Using the templates for below use cases.
- NetworkContainer     : Create IPv4/IPv6 Network Containers
- Network              : Create IPv4/IPv6 Network
- NextAvailableNetwork : Get next available network from a given parent CIDR of a prefix length.
- AllocationAndAssociation : Assign an IPv4 and IPv6 address to an AWS instance and get its MAC address synced at NIOS
- PTRRecord            : Create PTR record in forward/reverse mapping zones 
- CNAMERecord          : Create CNAME record in a zone

### Note
```
A parent network container has to be in existence before requesting next available network from it.
A forward/reverse mapping zone has to be in existence for a PTR record creation.
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
