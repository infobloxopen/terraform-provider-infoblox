# CNAME-record

CNAME-record maps one domain name to another (canonical) one.
CNAME-record resource allows managing such mappings at a NIOS server.
The resource has the following attributes: dns_view, canonical, alias,
ttl. 'canonical' and 'alias' are required, the rest are optional.
Consider the following table about the attributes:

| Attribute | Required/optional | Description | Example |
| --- | --- | --- | --- |
| alias | required | The alias name in FQDN format. | alias1.example.com |
| canonical | required | The canonical name in FQDN format. | main.example.com |
| dns_view | optional | DNS view which the zone does exist within. If omitted, the value 'default' is used. | tenant 1 view |

## Example CNAME-record resource definition

    resource "infoblox_cname_record" "ib_cname_record"{
      dns_view = "default" # the same as not specifying the attribute
      canonical = "CanonicalTestName.xyz.com"
      alias = "AliasTestName.xyz.com"
      ttl = 3600
    
      comment = "an example CNAME record"
      ext_attrs = jsonencode({
        "Tenant ID" = "tf-plugin"
        "CMP Type" = "Terraform"
        "Cloud API Owned" = "True"
      })
    }
