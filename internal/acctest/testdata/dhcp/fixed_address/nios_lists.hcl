case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ipv4addr         = "15.0.0.111"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
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
      ipv4addr         = "15.0.0.112"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        ipv4addr = "nios.ipv4addr"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ipv4addr         = "15.0.0.113"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ext_attrs        = { Site = "{{random}}" }
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
