# Rangetemplate — nios list cases
# No legacy list test was found for this object.
# Add list cases here manually.

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      ext_attrs = { "Tenant ID" = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "ext_attr_filters"
      values = {
        "Tenant ID" = "nios.ext_attrs.Tenant ID"
      }
    }
  }

}
