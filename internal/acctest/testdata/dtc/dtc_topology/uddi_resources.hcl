# Auto-generated resource acceptance-test cases for DtcTopology (UDDI backend).
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.name"                  = "topology-{{random}}"
      "uddi.sources.#"             = "1"
      "uddi.sources.0.name"        = "src-{{random}}"
      "uddi.sources.0.source"      = "subnet"
      "uddi.sources.0.subnets.#"   = "1"
      "uddi.sources.0.subnets.0"   = "10.0.0.0/8"
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
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      comment = "resource-comment"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.comment" = "resource-comment"
    }
  }

  step {
    uddi {
      name    = "topology-{{random}}"
      comment = "resource-comment-update"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.comment" = "resource-comment-update"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "topology-{{random}}"
      disabled = true
      sources  = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "topology-{{random}}"
      disabled = false
      sources  = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
      tags    = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
      tags    = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.name" = "topology-{{random}}"
    }
  }

  step {
    uddi {
      name    = "topology-{{random2}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
    check = {
      "uddi.name" = "topology-{{random2}}"
    }
  }

}

case "sources" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [
        { name = "src1-{{random}}", source = "subnet", subnets = ["10.0.0.0/8", "192.168.0.0/16"] }
      ]
    }
    check = {
      "uddi.sources.#"                 = "1"
      "uddi.sources.0.source"          = "subnet"
      "uddi.sources.0.subnets.#"       = "2"
    }
  }

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [
        { name = "src1-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] },
        { name = "src2-{{random}}", source = "subnet", subnets = ["172.16.0.0/12"] }
      ]
    }
    check = {
      "uddi.sources.#" = "2"
    }
  }

}

case "sources_tag_rule" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [
        {
          name      = "src-{{random}}"
          source    = "tag_rule"
          tag_rules = [{ key = "env", op = "EQUALS", value = "production" }]
        }
      ]
    }
    check = {
      "uddi.sources.#"                   = "1"
      "uddi.sources.0.source"            = "tag_rule"
      "uddi.sources.0.tag_rules.#"       = "1"
      "uddi.sources.0.tag_rules.0.key"   = "env"
      "uddi.sources.0.tag_rules.0.op"    = "EQUALS"
      "uddi.sources.0.tag_rules.0.value" = "production"
    }
  }

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [
        {
          name      = "src-{{random}}"
          source    = "tag_rule"
          tag_rules = [
            { key = "env", op = "EQUALS", value = "production" },
            { key = "region", op = "NOT_EQUALS", value = "us-east-1" }
          ]
        }
      ]
    }
    check = {
      "uddi.sources.0.tag_rules.#" = "2"
    }
  }

}
