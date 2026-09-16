# Auto-generated resource acceptance-test cases for Ipv6sharednetwork.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
    check = {
      "nios.name"                        = "{{random}}"
      "nios.networks.#"                  = "2"
      "nios.ddns_generate_hostname"      = "false"
      "nios.ddns_server_always_updates"  = "true"
      "nios.ddns_ttl"                    = "0"
      "nios.ddns_use_option81"           = "false"
      "nios.disable"                     = "false"
      "nios.enable_ddns"                 = "false"
      "nios.preferred_lifetime"          = "27000"
      "nios.update_dns_on_lease_renewal" = "false"
      "nios.valid_lifetime"              = "43200"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
  }

}

case "import" {
  backend  = "nios"
  parallel = true
  import   = true
  import_ignore = ["nios.networks"]
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      comment  = "Comment for the object"
    }
    check = {
      "nios.comment" = "Comment for the object"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      comment  = "Updated comment for the object"
    }
    check = {
      "nios.comment" = "Updated comment for the object"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name            = "{{random}}"
      networks        = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_domainname = "example.com"
    }
    check = {
      "nios.ddns_domainname" = "example.com"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      networks        = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_domainname = "updated-example.com"
    }
    check = {
      "nios.ddns_domainname" = "updated-example.com"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                   = "{{random}}"
      networks               = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_generate_hostname = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      networks               = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_generate_hostname = false
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
    }
  }

}

case "ddns_server_always_updates" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                       = "{{random}}"
      networks                   = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_server_always_updates = true
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "true"
    }
  }

  step {
    nios {
      name                       = "{{random}}"
      networks                   = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_server_always_updates = false
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "false"
    }
  }

}

case "ddns_ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_ttl = 100
    }
    check = {
      "nios.ddns_ttl" = "100"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_ttl = 200
    }
    check = {
      "nios.ddns_ttl" = "200"
    }
  }

}

case "ddns_use_option81" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name              = "{{random}}"
      networks          = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_use_option81 = true
    }
    check = {
      "nios.ddns_use_option81" = "true"
    }
  }

  step {
    nios {
      name              = "{{random}}"
      networks          = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ddns_use_option81 = false
    }
    check = {
      "nios.ddns_use_option81" = "false"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "domain_name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random}}"
      networks    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      domain_name = "example.com"
    }
    check = {
      "nios.domain_name" = "example.com"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      networks    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      domain_name = "updated-example.com"
    }
    check = {
      "nios.domain_name" = "updated-example.com"
    }
  }

}

case "domain_name_servers" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      networks            = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      domain_name_servers = ["2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"]
    }
    check = {
      "nios.domain_name_servers.#" = "3"
      "nios.domain_name_servers.0" = "2001:4860:4860::8888"
      "nios.domain_name_servers.1" = "2001:4860:4860::9999"
      "nios.domain_name_servers.2" = "2001:4860:4860::8899"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      networks            = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      domain_name_servers = ["2001:4860:4860::8881", "2001:4860:4860::9991"]
    }
    check = {
      "nios.domain_name_servers.#" = "2"
      "nios.domain_name_servers.0" = "2001:4860:4860::8881"
      "nios.domain_name_servers.1" = "2001:4860:4860::9991"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random}}"
      networks    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      enable_ddns = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      networks    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      enable_ddns = false
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random}}"
      networks  = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      networks  = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random}}"
      networks           = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      logic_filter_rules = [{ filter = "ipv6_option_filter", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      networks           = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      logic_filter_rules = [{ filter = "ipv6_option_filter1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name     = "{{random2}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
    check = {
      "nios.network_view" = "default"
    }
  }

}

case "networks" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
    }
    check = {
      "nios.networks.#" = "2"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}"]
    }
    check = {
      "nios.networks.#" = "1"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      options  = [{ name = "domain-name", num = "15", value = "example.com" }, { num = "37", value = "remote-id", vendor_class = "DHCPv6" }, { name = "dhcp6.subscriber-id", value = "subscriber-id", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.options.#"              = "3"
      "nios.options.0.name"         = "domain-name"
      "nios.options.0.num"          = "15"
      "nios.options.0.value"        = "example.com"
      "nios.options.1.num"          = "37"
      "nios.options.1.value"        = "remote-id"
      "nios.options.1.vendor_class" = "DHCPv6"
      "nios.options.2.name"         = "dhcp6.subscriber-id"
      "nios.options.2.value"        = "subscriber-id"
      "nios.options.2.vendor_class" = "DHCPv6"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      options  = [{ name = "domain-name", num = "15", value = "example.org" }, { num = "37", value = "remote-id-updated", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.options.#"              = "2"
      "nios.options.0.name"         = "domain-name"
      "nios.options.0.num"          = "15"
      "nios.options.0.value"        = "example.org"
      "nios.options.1.num"          = "37"
      "nios.options.1.value"        = "remote-id-updated"
      "nios.options.1.vendor_class" = "DHCPv6"
    }
  }

}

case "preferred_lifetime" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random}}"
      networks           = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      preferred_lifetime = 200
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "200"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      networks           = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      preferred_lifetime = 400
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "400"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                        = "{{random}}"
      networks                    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      networks                    = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      update_dns_on_lease_renewal = false
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
    }
  }

}

case "valid_lifetime" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random}}"
      networks       = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      valid_lifetime = 30000
    }
    check = {
      "nios.valid_lifetime" = "30000"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      networks       = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      valid_lifetime = 40000
    }
    check = {
      "nios.valid_lifetime" = "40000"
    }
  }

}
