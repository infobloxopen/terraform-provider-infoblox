# TsigKey — uddi list test cases
# No legacy list test exists (the BloxOne provider predates list resources);
# modelled on the shipped dns/view list cases.
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
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
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
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
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      tags   = { tag1 = "{{random}}" }
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
