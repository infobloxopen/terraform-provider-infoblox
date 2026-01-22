# VLAN

The `infoblox_vlan` resource enables you to perform `create`, `read`, `update`, and `delete` operations on VLANs in a NIOS appliance.
The resource represents the 'vlan' WAPI object in NIOS.

The following list describes the parameters you can define in the `infoblox_vlan` resource block:

* `parent`: required, specifies the VLAN View or VLAN Range to which this VLAN belongs (reference string). This field cannot be changed after creation.
* `name`: required, specifies the name of the VLAN.
* `vlan_id`: optional, specifies the VLAN ID value (typically 1-4094). If not specified, the next available VLAN ID will be automatically allocated from the parent range.
* `comment`: optional, describes the VLAN with a descriptive comment.
* `description`: optional, provides a description for the VLAN object, may be potentially used for longer VLAN names.
* `department`: optional, specifies the department where the VLAN is used.
* `contact`: optional, specifies the contact information for the person/team managing or using the VLAN.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that will be attached to the VLAN.

You can modify the `name`, `vlan_id`, `comment`, `description`, `department`, `contact`, and `ext_attrs` parameters after the VLAN is created.

### Finding the Parent Reference

To find the parent reference for your VLAN View or VLAN Range, you can use the Infoblox WAPI:

```bash
# List all VLAN views
curl -k -u "${INFOBLOX_USERNAME}:${INFOBLOX_PASSWORD}" \
  "https://${INFOBLOX_SERVER}/wapi/v2.12/vlanview"

# List all VLAN ranges
curl -k -u "${INFOBLOX_USERNAME}:${INFOBLOX_PASSWORD}" \
  "https://${INFOBLOX_SERVER}/wapi/v2.12/vlanrange"
```

### Example of VLAN Resource with Explicit VLAN ID

```hcl
resource "infoblox_vlan" "vlan100" {
  parent      = "vlanview/ZG5zLnZsYW5fdmlldyRkZWZhdWx0:default/1/4094"
  name        = "production-vlan"
  vlan_id     = 100
  comment     = "Production environment VLAN"
  description = "VLAN for production workloads"
  department  = "IT Operations"
  contact     = "ops-team@example.com"
  ext_attrs = jsonencode({
    "Site"        = "Datacenter 1"
    "Environment" = "Production"
  })
}
```

### Example of VLAN Resource with Auto-Allocated VLAN ID

When `vlan_id` is not specified, the next available VLAN ID will be automatically allocated from the parent range:

```hcl
resource "infoblox_vlan" "auto_vlan" {
  parent      = "vlanview/ZG5zLnZsYW5fdmlldyRkZWZhdWx0:default/1/4094"
  name        = "auto-allocated-vlan"
  comment     = "VLAN with auto-allocated ID"
  # vlan_id is omitted - will be auto-allocated from the parent range
}
```

After creation, you can reference the allocated VLAN ID using `infoblox_vlan.auto_vlan.vlan_id`.

### Minimal Resource Block

The minimal resource block required to create a VLAN is as follows:

```hcl
resource "infoblox_vlan" "simple_vlan" {
  parent = "vlanview/ZG5zLnZsYW5fdmlldyRkZWZhdWx0:default/1/4094"
  name   = "my-vlan"
  # vlan_id will be auto-allocated
}
```

Or with an explicit VLAN ID:

```hcl
resource "infoblox_vlan" "simple_vlan" {
  parent  = "vlanview/ZG5zLnZsYW5fdmlldyRkZWZhdWx0:default/1/4094"
  name    = "my-vlan"
  vlan_id = 100
}
```

### Import

You can import existing VLANs using their NIOS reference:

```shell
terraform import infoblox_vlan.vlan100 vlan/ZG5zLnZsYW4kMTAw:default/my-vlan/100
```

**Note:** After importing, you must specify the `parent` field in your Terraform configuration, as it cannot be automatically read from the API due to a limitation in the Infoblox Go client.
