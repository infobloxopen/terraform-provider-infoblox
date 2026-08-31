# Auto-generated datasource acceptance-test cases for Nsgroup.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.is_grid_default", "nios.is_multimaster", "nios.name", "nios.use_external_primary"]

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.is_grid_default", "nios.is_multimaster", "nios.name", "nios.use_external_primary"]

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs    = { Site = "{{random2}}" }
    }
  }

}
