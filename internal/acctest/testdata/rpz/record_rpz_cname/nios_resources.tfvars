# Auto-generated resource acceptance-test cases for RecordRpzCname.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : test-rpz.com  (view: default)
#   - RPZ zone : tf-acc-rpz.com  (view: tf-acc-rpz-view)
#   - DNS view : tf-acc-rpz-view
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.name"      = "{{random2}}.test-rpz.com"
      "nios.canonical" = "*"
      "nios.rp_zone"   = "test-rpz.com"
      "nios.view"      = "default"
      "nios.disable"   = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
  }

}

case "canonical" {
  backend  = "nios"
  parallel = true

  # "*" is the RPZ NODATA rule; a FQDN is a substitution rule.
  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.canonical" = "*"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "{{random3}}.com"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.canonical" = "{{random3}}.com"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "sub.{{random4}}.com"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.canonical" = "sub.{{random4}}.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      comment   = "This is a new rpz cname record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname record"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      comment   = "This is an updated rpz cname record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.name" = "{{random2}}.test-rpz.com"
    }
  }

  step {
    nios {
      name      = "{{random3}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.name" = "{{random3}}.test-rpz.com"
    }
  }

}

case "rp_zone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
    check = {
      "nios.rp_zone" = "test-rpz.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
      ttl       = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

  # ttl removed from config after being set: Optional+Computed must absorb
  # whatever the API reports instead of raising a "was null, but now N" diff.
  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      canonical = "*"
      rp_zone   = "test-rpz.com"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random2}}.tf-acc-rpz.com"
      canonical = "*"
      rp_zone   = "tf-acc-rpz.com"
      view      = "tf-acc-rpz-view"
    }
    check = {
      "nios.view"    = "tf-acc-rpz-view"
      "nios.rp_zone" = "tf-acc-rpz.com"
    }
  }

}
