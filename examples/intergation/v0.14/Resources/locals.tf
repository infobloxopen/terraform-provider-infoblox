locals {
    res_prefix = "terraform_example1"
    tenant_id = "${local.res_prefix}_tenant"

    # net_view = "default"
    # ... or (as a non-standard example)
    net_view = "${local.res_prefix}_netview"
}
