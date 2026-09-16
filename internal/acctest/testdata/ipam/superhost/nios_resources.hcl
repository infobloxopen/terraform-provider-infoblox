# Auto-generated resource acceptance-test cases for Superhost.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                      = "{{random}}"
      "nios.delete_associated_objects" = "false"
      "nios.disabled"                  = "false"
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
      name = "{{random}}"
    }
  }

}

case "import" {
  backend  = "nios"
  parallel = true
  import   = true
  import_ignore = ["nios.delete_associated_objects", "nios.ext_attrs_all"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "Comment for the object"
    }
    check = {
      "nios.comment" = "Comment for the object"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "Updated comment for the object"
    }
    check = {
      "nios.comment" = "Updated comment for the object"
    }
  }

}

case "delete_associated_objects" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                      = "{{random}}"
      delete_associated_objects = true
    }
    check = {
      "nios.delete_associated_objects" = "true"
    }
  }

  step {
    nios {
      name                      = "{{random}}"
      delete_associated_objects = false
    }
    check = {
      "nios.delete_associated_objects" = "false"
    }
  }

}

# NOTE: When prerequisites object gets implemented, we can remove the hardcoded values
case "dhcp_associated_objects" {
  backend  = "nios"
  parallel = true
  #
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_network" "parent_network" {
  #   nios = {
  #     network = "22.0.0.0/24"
  #     network_view = "default"
  #     comment = "Parent network for DHCP fixed addresses"
  #   }
  # }
  # resource "infoblox_fixed_address" "fixed_address" {
  #   nios = {
  #     ipv4addr = "22.0.0.20"
  #     match_client = "CIRCUIT_ID"
  #     agent_circuit_id = 23
  #   }
  # }
  # resource "infoblox_fixed_address" "fixed_address2" {
  #   nios = {
  #     ipv4addr = "22.0.0.21"
  #     match_client = "CIRCUIT_ID"
  #     agent_circuit_id = 24
  #   }
  # }
  # PREREQ
  step {
    nios {
      name                    = "{{random}}"
      dhcp_associated_objects = ["fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIwLjAuLg:22.0.0.20/default"]
    }
    check = {
      "nios.dhcp_associated_objects.0" = "fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIwLjAuLg:22.0.0.20/default"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      dhcp_associated_objects = ["fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIxLjAuLg:22.0.0.21/default"]
    }
    check = {
      "nios.dhcp_associated_objects.0" = "fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIxLjAuLg:22.0.0.21/default"
    }
  }

}

case "disabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      disabled = true
    }
    check = {
      "nios.disabled" = "true"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      disabled = false
    }
    check = {
      "nios.disabled" = "false"
    }
  }

}

# NOTE: When prerequisites object gets implemented, we can remove the hardcoded values
case "dns_associated_objects" {
  backend  = "nios"
  parallel = true
# extracted prerequisites_hcl:
#  prerequisites_hcl = <<-PREREQ
#   resource "infoblox_record_a" "record_a" {
#     nios = {
#       name = "parent-record_a.$${nios_dns_zone_auth.parent_auth_zone.fqdn}"
#       ipv4addr = "10.0.0.20"
#       view = "default"
#     }
#   }
#   resource "infoblox_record_aaaa" "record_aaaa" {
#     nios = {
#       name = "parent-record_aaaa.$${nios_dns_zone_auth.parent_auth_zone.fqdn}"
#       ipv6addr = "2002:1111::1401"
#       view = "default"
#     }
#   }
#   resource "infoblox_record_ptr_unknown" "record_ptr" {
#     nios = {
#       name = "23.252.168.192.in-addr.arpa"
#       ptrdname = "test.example.com"
#       view = "default"
#     }
#   }
#   resource "infoblox_zone_auth" "parent_auth_zone" {
#     nios = {
#       fqdn = "{{random2}}.com"
#       view = "default"
#     }
#   }
#   resource "infoblox_zone_auth" "parent_reverse_zone" {
#     nios = {
#       fqdn = "192.168.252.0/24"
#       view = "default"
#       zone_format = "IPV4"
#     }
#   }
#   resource "infoblox_ip_allocation_unknown" "allocation" {
#     nios = {
#       name = "parent-record_host.$${nios_dns_zone_auth.parent_auth_zone.fqdn}"
#       view = "default"
#     }
#   }
#   resource "infoblox_ip_association_unknown" "association" {
#     nios = {
#       ref = infoblox_ip_allocation_unknown.allocation.nios.ref
#       mac = "12:00:43:fe:9a:8c"
#       configure_for_dhcp = true
#     }
#   }
#   PREREQ

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "parent_auth_zone" {
    nios = {
      fqdn = "{{random2}}.com"
      view = "default"
    }
  }
  resource "infoblox_record_a" "record_a" {
    nios = {
      name = "parent-record-a.$${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
      ipv4addr = "10.0.0.20"
      view = "default"
    }
  }
  resource "infoblox_record_aaaa" "record_aaaa" {
    nios = {
      name = "parent-record-aaaa.$${infoblox_zone_auth.parent_auth_zone.nios.fqdn}"
      ipv6addr = "2002:1111::1401"
      view = "default"
    }
  }
  resource "infoblox_zone_auth" "parent_reverse_zone" {
    nios = {
      fqdn = "192.168.252.0/24"
      view = "default"
      zone_format = "IPV4"
    }
  }
  resource "infoblox_record_ptr" "record_ptr" {
    nios = {
      name = "23.252.168.192.in-addr.arpa"
      ptrdname = "test.example.com"
      view = "default"
    }
    depends_on = [infoblox_zone_auth.parent_reverse_zone]
  }
  PREREQ

  step {
    nios {
      name                   = "{{random}}"
      dns_associated_objects = ["${infoblox_record_a.record_a.id}", "${infoblox_record_aaaa.record_aaaa.id}"]
    }
    check_pair = {
      "nios.dns_associated_objects.0" = infoblox_record_a.record_a.id
      "nios.dns_associated_objects.1" = infoblox_record_aaaa.record_aaaa.id
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      dns_associated_objects = ["${infoblox_record_ptr.record_ptr.id}", "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5zdXBlcmhvc3QtYWNjdGVzdC5wYXJlbnQtcmVjb3JkLWhvc3Q:parent-record-host.superhost-acctest.com/default"]
    }
    check_pair = {
      "nios.dns_associated_objects.0" = infoblox_record_ptr.record_ptr.id
    }
    check = {
      "nios.dns_associated_objects.1" = "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5zdXBlcmhvc3QtYWNjdGVzdC5wYXJlbnQtcmVjb3JkLWhvc3Q:parent-record-host.superhost-acctest.com/default"
    }
  }

}
case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
