# Auto-generated resource acceptance-test cases for ServicerestartGroup.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
    }
    check = {
      "nios.name"    = "{{random}}"
      "nios.service" = "DNS"
      "nios.mode"    = "SIMULTANEOUS"
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
      name    = "{{random}}"
      service = "DNS"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "Sample Comment"
      service = "DNS"
    }
    check = {
      "nios.comment" = "Sample Comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "Updated Comment"
      service = "DNS"
    }
    check = {
      "nios.comment" = "Updated Comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      service   = "DNS"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      service   = "DNS"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      service = "DHCP"
      members = ["{{grid_master_hostname}}"]
    }
    check = {
      "nios.members.0" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
      members = ["{{grid_member_hostname}}"]
    }
    check = {
      "nios.members.0" = "{{grid_member_hostname}}"
    }
  }

}

case "mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
      mode    = "SEQUENTIAL"
    }
    check = {
      "nios.mode" = "SEQUENTIAL"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
      mode    = "SIMULTANEOUS"
    }
    check = {
      "nios.mode" = "SIMULTANEOUS"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name    = "{{random2}}"
      service = "DNS"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "recurring_schedule" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      service            = "DNS"
      recurring_schedule = { services = ["DHCPV6", "DNS"], mode = "SIMULTANEOUS", force = false, schedule = { weekdays = ["TUESDAY", "WEDNESDAY", "MONDAY"], frequency = "WEEKLY", every = 15, minutes_past_hour = 6, disable = false, repeat = "RECUR", hour_of_day = 20 } }
    }
    check = {
      "nios.recurring_schedule.services.#"                 = "2"
      "nios.recurring_schedule.services.0"                 = "DHCPV6"
      "nios.recurring_schedule.services.1"                 = "DNS"
      "nios.recurring_schedule.mode"                       = "SIMULTANEOUS"
      "nios.recurring_schedule.force"                      = "false"
      "nios.recurring_schedule.schedule.weekdays.0"        = "TUESDAY"
      "nios.recurring_schedule.schedule.weekdays.1"        = "WEDNESDAY"
      "nios.recurring_schedule.schedule.weekdays.2"        = "MONDAY"
      "nios.recurring_schedule.schedule.frequency"         = "WEEKLY"
      "nios.recurring_schedule.schedule.every"             = "15"
      "nios.recurring_schedule.schedule.minutes_past_hour" = "6"
      "nios.recurring_schedule.schedule.disable"           = "false"
      "nios.recurring_schedule.schedule.repeat"            = "RECUR"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      service            = "DNS"
      recurring_schedule = { services = ["ALL"], mode = "SIMULTANEOUS", force = true, schedule = { minutes_past_hour = 6, repeat = "ONCE", day_of_month = 30, month = 1, year = 2050, hour_of_day = 20 } }
    }
    check = {
      "nios.recurring_schedule.services.#"                 = "1"
      "nios.recurring_schedule.services.0"                 = "ALL"
      "nios.recurring_schedule.mode"                       = "SIMULTANEOUS"
      "nios.recurring_schedule.force"                      = "true"
      "nios.recurring_schedule.schedule.minutes_past_hour" = "6"
      "nios.recurring_schedule.schedule.repeat"            = "ONCE"
      "nios.recurring_schedule.schedule.day_of_month"      = "30"
      "nios.recurring_schedule.schedule.month"             = "1"
      "nios.recurring_schedule.schedule.year"              = "2050"
      "nios.recurring_schedule.schedule.hour_of_day"       = "20"
    }
  }

}

case "service" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
    }
    check = {
      "nios.service" = "DNS"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      service = "DHCP"
    }
    check = {
      "nios.service" = "DHCP"
    }
  }

}
