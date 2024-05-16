# Infoblox IPAM Driver for Terraform

## Resources

There are resources for the following objects, supported by the plugin:

* Network view (`infoblox_network_view`)
* Network container (`infoblox_ipv4_network_container`, `infoblox_ipv6_network_container`)
* Network (`infoblox_ipv4_network`, `infoblox_ipv6_network`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* DNS View (`infoblox_dns_view`)
* PTR-record (`infoblox_ptr_record`)
* CNAME-record (`infoblox_cname_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)
* Zone Auth (`infoblox_zone_auth`)
* Host record (`infoblox_ip_allocation` / `infoblox_ip_association`)
* Zone Delegated (`infoblox_zone_delegated`)

Network and network container resources have two versions: IPv4 and IPv6. In
addition, there are two operations which are implemented as resources:
IP address allocation and IP address association with a network host
(ex. VM in a cloud environment); they have three versions: IPv4
and IPv6 separately, and IPv4/IPv6 combined.

The recommendation is to avoid using separate IPv4 and IPv6 versions of
IP allocation and IP association resources.
This recommendation does not relate to network container and network-related resources.

To work with DNS records a user must ensure that appropriate DNS zones
exist on the NIOS side, because currently the plugin does not support
creating a DNS zone.

Every resource has common attributes: 'comment' and 'ext_attrs'.
'comment' is text which describes the resource. 'ext_attrs' is a set of
NIOS Extensible Attributes attached to the resource.

For DNS-related resources there is 'ttl' attribute as well, it specifies
TTL value (in seconds) for appropriate record. There is no default
value, zone's TTL is used by NIOS, if the value is omitted.
In this case, the field 'ttl' will be set to a negative value in the Terraform's state
for the resource, just to indicate the absence of a value.
TTL value of 0 (zero) means caching should be disabled for this record.

Please note that anywhere in the documents about this plugin, the default DNS view or
the default network view means: a view with the name `default`.
Usually, this is the name for the view which is marked as the default view on NIOS side, but this may be overridden.
But the plugin does use the name `default` for the view, despite which view is marked as the default on NIOS side.

## Data sources

There are data sources for the following objects:

* Network View (`infoblox_network_view`)
* IPv4 Network (`infoblox_ipv4_network`)
* IPv4 Network Container (`infoblox_ipv4_network_container`)
* A-record (`infoblox_a_record`)
* AAAA-record (`infoblox_aaaa_record`)
* CNAME-record (`infoblox_cname_record`)
* DNS View (`infoblox_dns_view`)
* PTR-record (`infoblox_ptr_record`)
* MX-record (`infoblox_mx_record`)
* TXT-record (`infoblox_txt_record`)
* SRV-record (`infoblox_srv_record`)
* Zone Auth (`infoblox_zone_auth`)

!> From version 2.5.0, new feature filters are introduced. Now the data sources support to populate more than one
matching NIOS objects.

* `filters`: the schema, with passing combination of searchable fields are supported by NIOS server, which
returns one or more matching objects from the NIOS server.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.

### Example for using filters:
```hcl
resource "infoblox_a_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ip_addr = "10.3.1.65"
  comment = "special host"
  dns_view = "nondefault_dnsview2"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}


data "infoblox_a_record" "a_rec_temp" {
  filters = {
    name = "very-interesting-host.example.com"
    ipv4addr = "10.3.1.65" //alias is ip_addr
    view = "nondefault_dnsview2"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_a_record' resource block before the data source will be queried.
  depends_on = [infoblox_a_record.vip_host]
}

output "a_rec_res" {
  value = data.infoblox_a_record.a_rec_temp
}

// accessing individual field in results
output "a_rec_name" {
  value = data.infoblox_a_record.a_rec_temp.results.0.fqdn //zero represents index of json object from results list
}
```

The list of matching objects as JSON format returned in output under results, with fields or arguments that are passed in the filters.

Filters will support `EA Search` i.e, fetches matching objects or records associated with the EAs' corresponding to provided data source, if any.

### Example for using filters for EA Search:
 ```hcl
 data "infoblox_a_record" "a_rec1" {
    filters = {
        "*TestEA" = "Acceptance Test"
    }
 }
 ```
Filters will also support Multi Value EA Search, where if the EA has more than one value, to be passed as comma seperated
value as a string. In here EAs' can have multiple or multi values of types like 'string', 'integer', etc..

### Example for using Multi Value EA Search:
```hcl
data "infoblox_a_record" "a_rec2" {
    filters = {
        "*tf_multi_val" = "test,test2,demo"
    }
 }

// for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
data "infoblox_a_record" "a_rec3" {
  filters = {
    "*tf_multi_val" = "test"
    "*tf_multi_val!" = "dummy"
  }
}
```

## Importing existing resources

There is a possibility to import existing resources, enabling them to be managed by Terraform.
As of now, there is a limitation: you have to write full resource's definition yourself.

In general, the process of importing an existing resource looks like this:

- write a new Terraform file (ex. a-record-imported.tf) with the content:
  ```
  resource "infoblox_a_record" "a_rec_1_imported" {
    fqdn = "rec-a-1.imported.test.com"
    ip_addr = "192.168.1.2"
    ttl = 10
    comment = "A-record to be imported"
    ext_attrs = jsonencode({
      "Location" = "New office"
    })
  }
  ```
- get a reference for the resource you want to import (ex. by using `curl` tool)
- issue a command of the form `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_REFERENCE`.
  Example: `terraform import infoblox_a_record.a_rec_1_imported record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQub3JnLmV4YW1wbGUsc3RhdGljMSwxLjIuMy40:rec-a-1.imported.test.com/default`

Please, note that if some of resource's properties (supported by the Infoblox provider plugin) is not defined or
is empty for the object on NIOS side, then appropriate resource's property must be empty or not defined.
Otherwise, you will get a difference in the resource's actual state and resource's description you specified,
and thus you will get a resource's update performed on the next `terraform apply` command invocation,
which will actually set the value of the property to the one which you defined (ex. empty value).

To import a host record (represented by the `infoblox_ip_allocation` and
`infoblox_ip_association` resources in Terraform), add the `Terraform Internal ID` extensible attribute
with a randomly generated value in the form of a UUID to the record.
- For steps to add the extensible attribute, refer to the [Infoblox NIOS Documentation](https://docs.infoblox.com).
- You may use the command-line tool `uuid` for Linux-based systems to generate a UUID.

-> The `Terraform Internal ID` extensible attribute is not shown in the terraform.tfstate file. Use it to create
   or import the `infoblox_ip_allocation` and `infoblox_ip_association` resources.
   You must not add it in a resource block with other extensible attributes.

!> You must not delete (ex. with 'terraform destroy' command) an `infoblox_ip_association` resource right after importing, but you may do this after 'terraform apply'.
   The reason: after 'terraform import' the dependency between `infoblox_ip_association` and respective `infoblox_ip_allocation` is not established by Terraform.
