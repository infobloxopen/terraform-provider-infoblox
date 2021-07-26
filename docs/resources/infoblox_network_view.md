# Network view

This resource represents the 'networkview' WAPI object in NIOS. It has
an attribute (in addition to the basic attributes 'comment' and
'ext_attrs', this note will be omitted thereafter) called 'name', which
is required to create a network view. It is a textual name, with the
same requirements as for the same attribute in WAPI.

To create a network view, you should use the following resource block in
terraform file (TF file):

    resource "infoblox_network_view" "nv1" {
      name = "network view 1"
      comment = "this is an example of network view"
      ext_attrs = jsonencode({
        "Site" = "Nevada"
      })
    }

Once the network view is created, you may change 'comment' and
'ext_attrs' (even remove them, or leave empty) but 'name' cannot be
changed. The minimal resource block to create a network view looks like
this:


    resource "infoblox_network_view" "nv1" {
      name = "network view 1"
    }
