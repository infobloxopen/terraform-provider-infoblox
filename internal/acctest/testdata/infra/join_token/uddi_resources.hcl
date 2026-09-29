# Auto-generated resource acceptance-test cases for JoinToken.
#
# join_token is returned once, on the create response, and never again. Cases
# assert only that it is set, never its value.
#
# name and description are RequiresReplaceIfConfigured, so a second step that
# changes either replaces the resource rather than updating it in place.

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
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
      name = "{{random}}"
    }
  }

}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      description = "Join token for site A"
    }
    check = {
      "uddi.description" = "Join token for site A"
    }
  }

  # A description change forces replacement, and BloxOne keeps the old name
  # reserved after the delete, so the replacement needs a fresh name.
  step {
    uddi {
      name        = "{{random2}}"
      description = "Join token for site B"
    }
    check = {
      "uddi.description" = "Join token for site B"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      tags = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "expires_at" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      expires_at = "{{future_rfc3339_24h}}"
    }
    check = {
      "uddi.expires_at" = "{{future_rfc3339_24h}}"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      expires_at = "{{future_rfc3339_48h}}"
    }
    check = {
      "uddi.expires_at" = "{{future_rfc3339_48h}}"
    }
  }

}
