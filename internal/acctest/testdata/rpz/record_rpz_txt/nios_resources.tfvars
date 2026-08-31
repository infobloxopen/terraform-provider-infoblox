# Auto-generated resource acceptance-test cases for RecordRpzTxt.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : test-rpz.com  (view: default)
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
    check = {
      "nios.text"    = "Record Text"
      "nios.name"    = "{{random2}}.test-rpz.com"
      "nios.rp_zone" = "test-rpz.com"
      "nios.view"    = "default"
      "nios.disable" = "false"
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
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      comment = "This is a new rpz txt record"
    }
    check = {
      "nios.comment" = "This is a new rpz txt record"
    }
  }

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      comment = "This is an updated rpz txt record"
    }
    check = {
      "nios.comment" = "This is an updated rpz txt record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      disable = true
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
      text      = "Record Text"
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
      text      = "Record Text"
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
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
    check = {
      "nios.name" = "{{random2}}.test-rpz.com"
    }
  }

  step {
    nios {
      name    = "{{random3}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
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
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
    check = {
      "nios.rp_zone" = "test-rpz.com"
    }
  }

}

case "text" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
    check = {
      "nios.text" = "Record Text"
    }
  }

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Updated Record Text"
      rp_zone = "test-rpz.com"
    }
    check = {
      "nios.text" = "Updated Record Text"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      ttl     = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
      ttl     = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}
