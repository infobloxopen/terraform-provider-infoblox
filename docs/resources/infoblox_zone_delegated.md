# Resource Zone Delegated

A Zone Delegated resource creates NS records for a subdomain, pointing to one or more external authoritative name servers. The `infoblox_zone_delegated` resource allow managing such delegations. The parent zone must already exist 

The following list describes the parameters you can define in the `infoblox_zone_delegated` resource block:

## Argument Reference
* `fqdn`: (Required) The subdomain name to be delegated
* `delegate_to`: (Required) Nested block(s)s for the delegated name servers
    * `name`: (Required) The FQDN of the name server
* `ext_attrs`: (Optional) A set of NIOS extensible attributes that are attached to the record, using jsonencode. Currently only "Tenant ID" is supported

## Attribute Reference
* `delegate_to`:
    * `address`: The computed IP address for each delegated name server

## Example Usage

```hcl
resource "infoblox_zone_delegated" "subdomain" {

  fqdn = "subdomain.test.com"

  delegate_to {
    name = "ns-1488.awsdns-58.org"
  }

  delegate_to {
    name = "ns-2034.awsdns-62.co.uk"
  }

}
```

## Import
Zone Delegated resources can be imported by using either the object reference or the subdomain fqdn, for example:
```shell script
# terraform import infoblox_zone_delegated.subdomain zone_delegated/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5jb2xsZWdlY2hvaWNldHJhbnNpdGlvbi5nc2xi:subdomain.test.com/default
# terraform import infoblox_zone_delegated.subdomain subdomain.test.com
```
