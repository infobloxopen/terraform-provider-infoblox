# Auto-generated resource acceptance-test cases for HaGroup.
#  TODO: Objects to be present in the grid for testing
#  DHCP Hosts

case "basic" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
    check = {
      "uddi.hosts.#"      = "2"
      "uddi.hosts.0.role" = "active"
      "uddi.hosts.1.role" = "active"
      "uddi.name"         = "{{random}}"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
  }

}

case "comment" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name    = "{{random}}"
      mode    = "active-active"
      comment = "HA Group created with Terraform"
    }
    check = {
      "uddi.comment" = "HA Group created with Terraform"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name    = "{{random}}"
      mode    = "active-active"
      comment = "HA Group was created with Terraform"
    }
    check = {
      "uddi.comment" = "HA Group was created with Terraform"
    }
  }

}

case "hosts" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
    }
    check = {
      "uddi.hosts.#"      = "2"
      "uddi.hosts.0.role" = "active"
      "uddi.hosts.1.role" = "passive"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_1}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
    }
    check = {
      "uddi.hosts.#"      = "2"
      "uddi.hosts.0.role" = "active"
      "uddi.hosts.1.role" = "passive"
    }
  }

}

case "mode" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
    check = {
      "uddi.mode" = "active-active"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
    }
    check = {
      "uddi.mode" = "active-passive"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "advanced-active-passive"
    }
    check = {
      "uddi.mode" = "advanced-active-passive"
    }
  }

}

case "split_ranges" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "split-ranges"
    }
    check = {
      "uddi.mode" = "split-ranges"
    }
  }

}

case "name" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random}}"
      mode = "active-active"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "active" }
      ]
      name = "{{random2}}"
      mode = "active-active"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "tags" {
  backend = "uddi"

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      hosts = [
        { host = "{{uddi_dhcp_host_id_1}}", role = "active" },
        { host = "{{uddi_dhcp_host_id_2}}", role = "passive" }
      ]
      name = "{{random}}"
      mode = "active-passive"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
