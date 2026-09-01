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

# NOTE: infoblox_fixed_address is not implemented yet, so the fixed-address prerequisites
# cannot be created by Terraform. infoblox_network IS implemented, but the parent network
# deliberately stays out of prerequisites_hcl too, because the two are coupled:
#   * NIOS cascades a network delete to its child fixed addresses (verified: deleting
#     23.0.0.0/24 silently removed fixed address 23.0.0.5), so a Terraform-managed network
#     would destroy the manually created fixed addresses at the end of every run — the case
#     would pass once and fail from then on.
#   * The network already exists, so a Terraform create would fail outright with
#     "The network 22.0.0.0/24 already exists. Select another network."
# So the network and both fixed addresses were created on the grid out of band (WAPI) and
# their refs are inlined below, so this case runs today:
#   network/ZG5zLm5ldHdvcmskMjIuMC4wLjAvMjQvMA:22.0.0.0/24/default
#   fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIwLjAuLg:22.0.0.20/default
#   fixedaddress/ZG5zLmZpeGVkX2FkZHJlc3MkMjIuMC4wLjIxLjAuLg:22.0.0.21/default
# When infoblox_fixed_address lands, delete the two literal refs and the checks below,
# restore prerequisites_hcl (kept verbatim in the commented block) and switch the step
# values back to ${infoblox_fixed_address.fixed_address[2].id}.

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

# NOTE: hybrid prerequisites. Step 1 references infoblox_record_a and infoblox_record_aaaa,
# which ARE implemented, so their parent zone and both records are created by Terraform via
# prerequisites_hcl below and asserted with check_pair.
# Step 2 references a PTR record and a host record (legacy nios_ip_allocation +
# nios_ip_association); infoblox_record_ptr / infoblox_ip_allocation / infoblox_ip_association
# are NOT implemented yet, so those two objects — and their parents, since NIOS cascades a
# zone delete to the records inside it — were created on the grid out of band (WAPI) and
# their refs are inlined in step 2:
#   zone_auth/…:192.168.252.0%2F24/default        (reverse zone, parent of the PTR)
#   zone_auth/…:superhost-acctest.com/default     (forward zone, parent of the host record)
#   network/…:192.168.1.0/24/default              (needed for the host record's DHCP config)
#   record:ptr/ZG5zLmJpbmRfcHRyJC5fZGVmYXVsdC5hcnBhLmluLWFkZHIuMTkyLjE2OC4yNTIuMjMudGVzdC5leGFtcGxlLmNvbQ:23.252.168.192.in-addr.arpa/default
#   record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5zdXBlcmhvc3QtYWNjdGVzdC5wYXJlbnQtcmVjb3JkLWhvc3Q:parent-record-host.superhost-acctest.com/default
# All carry comment "acctest prereq: superhost dns_associated_objects". When the three
# resources land, move these into prerequisites_hcl and swap step 2 to check_pair
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
      dns_associated_objects = ["record:ptr/ZG5zLmJpbmRfcHRyJC5fZGVmYXVsdC5hcnBhLmluLWFkZHIuMTkyLjE2OC4yNTIuMjMudGVzdC5leGFtcGxlLmNvbQ:23.252.168.192.in-addr.arpa/default", "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5zdXBlcmhvc3QtYWNjdGVzdC5wYXJlbnQtcmVjb3JkLWhvc3Q:parent-record-host.superhost-acctest.com/default"]
    }
    check = {
      "nios.dns_associated_objects.0" = "record:ptr/ZG5zLmJpbmRfcHRyJC5fZGVmYXVsdC5hcnBhLmluLWFkZHIuMTkyLjE2OC4yNTIuMjMudGVzdC5leGFtcGxlLmNvbQ:23.252.168.192.in-addr.arpa/default"
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
