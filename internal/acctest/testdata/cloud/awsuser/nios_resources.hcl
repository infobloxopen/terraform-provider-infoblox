# Auto-generated resource acceptance-test cases for Awsuser.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - NIOS admin user : aws1
#   - NIOS admin user : aws2

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      access_key_id     = "AKIA{{random}}"
      account_id        = "337773173961"
      name              = "{{random2}}"
      secret_access_key = "S1JGWfwcZWEY+hSkfpyhxigL9A/uaJ96mY"
      govcloud_enabled  = false
    }
    check = {
      "nios.access_key_id"    = "AKIA{{random}}"
      "nios.account_id"       = "337773173961"
      "nios.name"             = "{{random2}}"
      "nios.govcloud_enabled" = "false"
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
      access_key_id     = "AKIA{{random}}"
      account_id        = "337773173961"
      name              = "{{random2}}"
      secret_access_key = "S1JGWfwcZWEY+jhSkfpyhxigL9A/ua6mY"
      govcloud_enabled  = false
    }
  }

}

case "access_key_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      access_key_id     = "AKIA{{random2}}"
      account_id        = "337773173961"
      name              = "{{random}}"
      secret_access_key = "S1JGWfwcZWEY+hSkfpyhxigL9A/ua96mY"
    }
    check = {
      "nios.account_id" = "337773173961"
      "nios.name"       = "{{random}}"
    }
  }

  step {
    nios {
      access_key_id     = "AKIA{{random3}}"
      account_id        = "337773173961"
      name              = "{{random}}"
      secret_access_key = "S1JGWfwcZWEY+hSkfpyhxigL9A/ua96mY"
    }
    check = {
      "nios.account_id" = "337773173961"
      "nios.name"       = "{{random}}"
    }
  }

}

case "account_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      account_id        = "33773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.account_id" = "33773173961"
    }
  }

  step {
    nios {
      account_id        = "12345689012"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.account_id" = "12345689012"
    }
  }

}

case "govcloud_enabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = true
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.govcloud_enabled" = "true"
    }
  }

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.govcloud_enabled" = "false"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}-updated"
      govcloud_enabled  = false
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.name" = "{{random2}}-updated"
    }
  }

}

case "nios_user_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      nios_user_name    = "aws1"
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.nios_user_name" = "aws1"
    }
  }

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      nios_user_name    = "aws2"
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
    check = {
      "nios.nios_user_name" = "aws2"
    }
  }

  # Unset path: field omitted on PUT (flex: empty_to_null), NIOS retains last value; computed: true prevents drift.
  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    }
  }

}

case "secret_access_key" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/J96mY"
    }
    check = {
      "nios.secret_access_key" = "S1JGWfwcZWEYhSkfpyhxigL9A/J96mY"
    }
  }

  step {
    nios {
      account_id        = "337773173961"
      access_key_id     = "AKIA{{random}}"
      name              = "{{random2}}"
      govcloud_enabled  = false
      secret_access_key = "K1JGWfwcZWEYYhSkfpyhxigL9A/J96mY"
    }
    check = {
      "nios.secret_access_key" = "K1JGWfwcZWEYYhSkfpyhxigL9A/J96mY"
    }
  }

}
