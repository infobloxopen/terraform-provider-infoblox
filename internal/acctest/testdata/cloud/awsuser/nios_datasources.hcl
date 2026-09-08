# Awsuser — nios datasource test cases
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.access_key_id", "nios.account_id", "nios.govcloud_enabled", "nios.name", "nios.nios_user_name"]

  step {
    nios {
      name             = "{{random}}"
      access_key_id    = "AKIA{{random2}}"
      account_id       = "337773173961"
      govcloud_enabled = false
      # tostring() makes this a RawExpr so the framework skips the auto pair-check
      # for secret_access_key (writeOnly — NIOS never echoes it back in GET responses).
      secret_access_key = tostring("S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY")
    }
  }

}
