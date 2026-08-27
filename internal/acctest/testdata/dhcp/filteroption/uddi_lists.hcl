# Filteroption — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.
// An option code has to be created before running the test cases.

case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }]
      }
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }]
      }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random2}}" }
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "dhcp/option_code/de50b0db-01cc-4da8-8213-aefd0880340f"
          option_value = "value1"
        }]
      }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
