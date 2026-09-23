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
      tags               = { tag1 = "{{random}}", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random2}}"
      service            = "DNS"
      tags               = { tag2 = "{{random}}", tag3 = "{{random3}}" }
    }
    check = {
      "uddi.tags.tag2" = "{{random}}"
      "uddi.tags.tag3" = "{{random3}}"
    }
  }

}
