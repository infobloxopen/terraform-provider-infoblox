case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      service            = "NTP"
      anycast_ip_address = "{{random_ip}}"
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
      name               = "{{random}}"
      service            = "NTP"
      anycast_ip_address = "{{random_ip}}"
    }
  }

}

case "anycast_ip_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "NTP"
    }
    check = {
      "uddi.name"               = "{{random}}"
      "uddi.anycast_ip_address" = "{{random_ip}}"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip2}}"
      name               = "{{random}}"
      service            = "NTP"
    }
    check = {
      "uddi.name"               = "{{random}}"
      "uddi.anycast_ip_address" = "{{random_ip2}}"
    }
  }

}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "DNS"
      description        = "Anycast comment"
    }
    check = {
      "uddi.description" = "Anycast comment"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "DNS"
      description        = "Anycast comment updated"
    }
    check = {
      "uddi.description" = "Anycast comment updated"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "DNS"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random2}}"
      service            = "DNS"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "onprem_hosts" {
  backend  = "uddi"
  parallel = true
  skip_if_env_empty = ["UDDI_INFRA_HOST_DISPLAY_NAME_1","UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_DISPLAY_NAME_1 and UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"


  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "NTP"
      onprem_hosts = [
        {
          id   = "{{uddi_infra_host_legacy_id_1}}"
          name = "{{uddi_infra_host_display_name_1}}",
        }
      ]
    }
    check = {
      "uddi.onprem_hosts.#"      = "1"
      "uddi.onprem_hosts.0.id"   = "{{uddi_infra_host_legacy_id_1}}"
      "uddi.onprem_hosts.0.name" = "{{uddi_infra_host_display_name_1}}",
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "NTP"

    }
    check = {
      "uddi.onprem_hosts.#" = "0"
    }
  }
}

case "service" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "NTP"
    }
    check = {
      "uddi.service" = "NTP"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "DNS"
    }
    check = {
      "uddi.service" = "DNS"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random2}}"
      service            = "DNS"
      tags               = { tag1 = "{{random}}", tag2 = "{{random2}}" }
    }
    check = {
      "uddi.tags.tag1" = "{{random}}"
      "uddi.tags.tag2" = "{{random2}}"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random3}}"
      service            = "DNS"
      tags               = { tag2 = "{{random}}", tag3 = "{{random4}}" }
    }
    check = {
      "uddi.tags.tag2" = "{{random}}"
      "uddi.tags.tag3" = "{{random4}}"
    }
  }

}
