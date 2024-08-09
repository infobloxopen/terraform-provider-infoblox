# DNS View Resource

The `infoblox_dns_view` resource enables you to perform `create`, `update` and
`delete` operations on DNS views in a NIOS appliance.
The resource represents the ‘view’ WAPI object in NIOS.

The following list describes the parameters you can define in the `infoblox_dns_view` resource block:

* `name`: required, specifies the name of the DNS View. Example: `nondefault_dnsview`.
* `network_view`: optional, specifies the name of the Network View in which DNS View exists. If value is not specified ,the `default`
will be considered as default networkview. Example: `custom_netview`.
* `comment`: optional, describes the DNS view. Example: `example DNS view`.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to DNS view. Example: `jsonencode({})`.

You can update 'name' of the DNS view created in resource block, as it can be modified in NIOS.

### Examples of an DNS View Block

```hcl
//creating DNS view resource with minimal set of parameters
resource "infoblox_dns_view" "view1" {
  name = "test_view"
}

//creating DNS view resource with full set of parameters
resource "infoblox_dns_view" "view2" {
  name         = "customview"
  network_view = "default"
  comment      = "test dns view example"
  ext_attrs = jsonencode({
    "Site" = "Main test site"
  })
}

// creating DNS View under non default network view
resource "infoblox_dns_view" "view3" {
  name         = "custom_view"
  network_view = "non_defaultview"
  comment      = "example under custom network view"
  ext_attrs = jsonencode({
    "Site" = "Cal Site"
  })
}
```
