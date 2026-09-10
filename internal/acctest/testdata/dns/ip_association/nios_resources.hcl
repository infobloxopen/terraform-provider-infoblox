# Hand-authored resource acceptance-test cases for IPAssociation.
case "basic" {
  backend           = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }

  resource "infoblox_record_host" "test" {
    nios = {
      name      = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ipv4addrs = [{ ipv4addr = "{{random_ip}}" }]
    }
  }
  PREREQ

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac}}"
      configure_for_dhcp = false
    }
    check = {
      "nios.mac"                = "{{random_mac}}"
      "nios.configure_for_dhcp" = "false"
    }
  }

}

case "configure_for_dhcp" {
  backend           = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }

  resource "infoblox_record_host" "test" {
    nios = {
      name      = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ipv4addrs = [{ ipv4addr = "{{random_ip}}" }]
    }
  }
  PREREQ

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac}}"
      configure_for_dhcp = true
    }
    check = {
      "nios.configure_for_dhcp" = "true"
    }
  }

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac}}"
      configure_for_dhcp = false
    }
    check = {
      "nios.configure_for_dhcp" = "false"
    }
  }

}

case "mac" {
  backend           = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }

  resource "infoblox_record_host" "test" {
    nios = {
      name      = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ipv4addrs = [{ ipv4addr = "{{random_ip}}" }]
    }
  }
  PREREQ

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac}}"
      configure_for_dhcp = true
    }
    check = {
      "nios.mac" = "{{random_mac}}"
    }
  }

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac2}}"
      configure_for_dhcp = true
    }
    check = {
      "nios.mac" = "{{random_mac2}}"
    }
  }

}

case "duid" {
  backend           = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }

  resource "infoblox_record_host" "test" {
    nios = {
      name      = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ipv6addrs = [{ ipv6addr = "2002:1f93::12:31" }]
    }
  }
  PREREQ

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      duid               = "{{random_duid}}"
      match_client       = "DUID"
      configure_for_dhcp = true
    }
    check = {
      "nios.duid"               = "{{random_duid}}"
      "nios.match_client"       = "DUID"
      "nios.configure_for_dhcp" = "true"
    }
  }

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      duid               = "{{random_duid2}}"
      match_client       = "DUID"
      configure_for_dhcp = true
    }
    check = {
      "nios.duid" = "{{random_duid2}}"
    }
  }

}

case "match_client" {
  backend           = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }

  resource "infoblox_record_host" "test" {
    nios = {
      name      = "{{random2}}.$${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ipv6addrs = [{ ipv6addr = "2002:1f93::12:32" }]
    }
  }
  PREREQ

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      duid               = "{{random_duid}}"
      match_client       = "DUID"
      configure_for_dhcp = true
    }
    check = {
      "nios.match_client" = "DUID"
    }
  }

  step {
    nios {
      record_host_id     = infoblox_record_host.test.id
      mac                = "{{random_mac}}"
      match_client       = "MAC_ADDRESS"
      configure_for_dhcp = true
    }
    check = {
      "nios.match_client" = "MAC_ADDRESS"
      "nios.mac"          = "{{random_mac}}"
    }
  }

}
