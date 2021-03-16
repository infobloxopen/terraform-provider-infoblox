locals {
    res_prefix = "terraform_example1"
    tenant_id = "${local.res_prefix}_tenant"

    # net_view = "default"
    # ... or (as a non-standard example)
    net_view = "${local.res_prefix}_netview"
    # This network view must exist in the Grid.

    # This CIDR must belong to an existing network *container*
    # in the network view specified above.
    parent_cidr = "10.0.0.0/16"
}
