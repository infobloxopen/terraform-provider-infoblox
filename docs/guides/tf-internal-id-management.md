---
page_title: "Management of Provider Operations with Terraform Internal ID"
subcategory: "Guides"
description: |-
  How the Infoblox provider uses the Terraform Internal ID extensible attribute to track NIOS resources and handle drift. Applies to the NIOS backend only.
---

# Management of Provider Operations with Terraform Internal ID

> **Note:** This guide applies to the **NIOS backend only**. Resources managed through the `uddi` backend do not use the Terraform Internal ID extensible attribute and need no extra setup.

The Infoblox provider is a unified provider that works with both NIOS and UDDI. When the provider is configured with a `nios` block, it uses **Terraform Internal ID**, an extensible attribute created in NIOS, to manage operations performed on the resource objects it supports.

As a prerequisite for the NIOS backend, the extensible attribute definition for Terraform Internal ID must exist in NIOS. The extensible attribute is used to associate a unique ID to each Terraform resource to prevent any sort of drifts which may occur when external updates happen and the WAPI reference gets changed. For methods that you can use to create the extensible attribute, see [Creating the Terraform Internal ID Extensible Attribute](#creating-the-terraform-internal-id-extensible-attribute).

The extensible attribute is used for the following:
* To manage drift state in Terraform. For more information, see [Management of Drift State in Terraform](#management-of-drift-state-in-terraform).
* To create and manage NIOS resources supported by the provider.

This behavior is controlled by the provider-level `manage_internal_id_ea` setting, which defaults to `true`:

```terraform
provider "infoblox" {
  manage_internal_id_ea = true

  nios = {
    host_url = "<NIOS_HOST_URL>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}
```

When `manage_internal_id_ea` is `false`, the provider does not validate, create, update, or otherwise manage the attribute.

According to the operation that you perform in Terraform, the behavior exhibited by the provider on the NIOS backend is as follows:

**Creating a resource:**
* If you are creating a resource in Terraform and the provider is able to find the Terraform Internal ID extensible attribute, it attaches it to the resource in NIOS and saves it for that resource in the state file.

**Modifying an existing resource:**
* If the provider finds a match for the reference ID and Terraform Internal ID, it completes the update operation.
* If the provider does not find a match for the reference ID, but finds a match for the Terraform Internal ID, it proceeds with the update operation and also retrieves the changed reference ID from NIOS and updates it in the state file. For more information, see [Management of Drift State in Terraform](#management-of-drift-state-in-terraform).
* If the provider does not find a match for either the reference ID or the Terraform Internal ID, the provider clears the resource from the state file and tries to recreate the resource on a subsequent run of the `terraform apply` command.

**Importing a resource from NIOS:**
* When you import an existing resource from NIOS to Terraform using the resource reference, the provider creates the Terraform Internal ID extensible attribute and attaches it to the resource in NIOS and saves it in the state file in Terraform.
* If a Terraform Internal ID extensible attribute already exists on the NIOS object during import, it will be overwritten with a newly generated value.

## Creating the Terraform Internal ID Extensible Attribute

Only a NIOS admin with superuser privileges is authorized to create extensible attributes in NIOS. For more information about NIOS admin accounts, refer to the [Infoblox NIOS Documentation](https://docs.infoblox.com/).

Use one of the following methods to create the Terraform Internal ID extensible attribute:

1. **Create the extensible attribute manually in Infoblox NIOS Grid Manager.** For steps, refer to the Adding Extensible Attributes topic in the [Infoblox NIOS Documentation](https://docs.infoblox.com/).
   * If the user you want to manage is a cloud member, then enable the following option for the extensible attribute:
     * In Grid Manager, on the Administration tab > Extensible Attributes tab, edit the extensible attribute.
     * On the Additional Properties tab, enable "Allow cloud members to have the following access to this extensible attribute" and select "Read/Write (and disallow Write access from the GUI and the standard API)".

2. **Use the following cURL command to create the extensible attribute as a read-only attribute in NIOS:**
```bash
   curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/extensibleattributedef -d '{"name": "Terraform Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Terraform Resource"}'
```
   * If the user you want to manage is a cloud member, then include the flag `C` for cloud API.
   * If you are using multiple flags in the command, ensure that the flags are written in correct order. For more information about flags, refer to the Extensible Attribute Definition object in the [Infoblox WAPI documentation](https://docs.infoblox.com/space/NIOS/35400616/NIOS).

3. **Let the provider create the extensible attribute automatically** by configuring the `nios` block with credentials of a NIOS admin user with superuser privileges and leaving `manage_internal_id_ea` set to `true`.

### Notes

* With `manage_internal_id_ea = true`, either the Terraform Internal ID extensible attribute definition must be present in NIOS or the `nios` credentials must have superuser access so the provider can create it. If not, the connection to NIOS will fail.
* If you choose to create the Terraform Internal ID extensible attribute manually or by using the cURL command, the creation of the extensible attribute is not managed by the provider.
* You must not modify the Terraform Internal ID for a resource under any circumstances. If it is modified, the resource will no longer be managed by Terraform.

## Management of Drift State in Terraform

On the NIOS backend, the provider can track and manage drift state that is caused due to a mismatch in the reference ID of a resource saved in the Terraform state file with that of its counterpart in Infoblox NIOS. To detect and resolve the drift state, the provider uses two levels of validation to identify a resource. First, with a reference ID issued by Infoblox NIOS WAPI, and then with the extensible attribute, Terraform Internal ID. The reference ID of a resource might change, but its Terraform Internal ID does not, so the provider can still find the resource. If a mismatch is detected, the provider takes appropriate measure to fix it.
