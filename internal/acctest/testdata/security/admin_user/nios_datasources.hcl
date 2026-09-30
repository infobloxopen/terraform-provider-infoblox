# Auto-generated datasource acceptance-test cases for Adminuser.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.auth_method", "nios.auth_type", "nios.disable", "nios.enable_certificate_authentication", "nios.name", "nios.time_zone"]

  step {
    nios {
      name         = "{{random}}"
      password     = tostring("Example-Admin123!")
      admin_groups = ["admin-group"]
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

  pair_checks = ["nios.auth_method", "nios.auth_type", "nios.comment", "nios.disable", "nios.enable_certificate_authentication", "nios.name", "nios.time_zone"]

  step {
    nios {
      name         = "{{random2}}"
      password     = tostring("Example-Admin123!")
      admin_groups = ["admin-group"]
      ext_attrs    = { Site = "admin-group" }
    }
  }

}
