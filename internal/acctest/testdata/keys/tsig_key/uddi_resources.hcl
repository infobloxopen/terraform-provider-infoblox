# TsigKey — uddi resource test cases
# NOTE: mirrors the legacy BloxOne TSIG coverage; `secret` is Required because the API rejects a create without it (values are Base64-encoded).
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
    check = {
      "uddi.name"      = "{{random}}."
      "uddi.secret"    = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      "uddi.algorithm" = "hmac_sha256"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
    check = {
      "uddi.name" = "{{random}}."
    }
  }

  step {
    uddi {
      name   = "{{random2}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
    check = {
      "uddi.name" = "{{random2}}."
    }
  }

}

case "algorithm" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name      = "{{random}}."
      secret    = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      algorithm = "hmac_sha512"
    }
    check = {
      "uddi.algorithm" = "hmac_sha512"
    }
  }

  step {
    uddi {
      name      = "{{random}}."
      secret    = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      algorithm = "hmac_sha1"
    }
    check = {
      "uddi.algorithm" = "hmac_sha1"
    }
  }

}

case "secret" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
    check = {
      "uddi.secret" = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }

  step {
    uddi {
      name   = "{{random}}."
      secret = "FzpyuZuQAHxLmwZVGlYcwaPB7Ow9MSWqSyyJlNR1XUc="
    }
    check = {
      "uddi.secret" = "FzpyuZuQAHxLmwZVGlYcwaPB7Ow9MSWqSyyJlNR1XUc="
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}."
      secret  = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      comment = "key created through terraform"
    }
    check = {
      "uddi.comment" = "key created through terraform"
    }
  }

  step {
    uddi {
      name    = "{{random}}."
      secret  = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      comment = "key updated through terraform"
    }
    check = {
      "uddi.comment" = "key updated through terraform"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      tags   = { tag1 = "value1" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      tags   = { tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag2" = "value2"
    }
  }

}
