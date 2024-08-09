# Network view

The `infoblox_network_view` resource enables you to perform `create` and
`update` operations on network views in a NIOS appliance.
The resource represents the ‘networkview’ WAPI object in NIOS.

The following list describes the parameters you can define in the `infoblox_network_view` resource block:

* `name`: required, specifies the desired name of the network view as shown in the NIOS appliance. The name has the same requirements as the corresponding parameter in WAPI.
* `comment`: optional, describes the network view.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the network view.

!>  Once the network view is created, you cannot change the `name` parameter.

You can modify or even remove the `comment` and ext_attrs `parameters` from
the resource block.

### Example of Network View Resource

```hcl
resource "infoblox_network_view" "netview1234" {
  name    = "one_more_network_view"
  comment = "example network view"
  ext_attrs = jsonencode({
    "Location" = "the North pole"
  })
}
```

The minimal resource block required to create a network view is as follows:

```hcl
resource "infoblox_network_view" "nv1" {
  name = "network view 1"
}
```
