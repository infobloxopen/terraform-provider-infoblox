case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
    check = {
      "nios.name"                              = "{{random}}"
      "nios.auth_method"                       = "KEYPAIR"
      "nios.auth_type"                         = "LOCAL"
      "nios.disable"                           = "false"
      "nios.enable_certificate_authentication" = "false"
      "nios.time_zone"                         = "UTC"
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
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
  }

}

case "admin_groups" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
    check = {
      "nios.admin_groups.0" = "admin-group"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["opa-group"]
    }
    check = {
      "nios.admin_groups.0" = "opa-group"
    }
  }

}

case "auth_method" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_method  = "KEYPAIR"
    }
    check = {
      "nios.auth_method" = "KEYPAIR"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_method  = "KEYPAIR_PASSWORD"
      ssh_keys     = [{ key_name = "sample-key", key_type = "RSA", key_value = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4duV7U2EWRn/Auxa+iX2nx/ffrH249xzDw2fUvP8aeDskP+M1gDJnNZG21Ty2HH1kMZ1qRpN3DY0YBRYopA5P3EFfkq4La8OdG6IH47CINdqCHQTKGZ40YCdCqqz9hd796rrQI1e9yY08MLob79M7jPZLsafDoB7sa4ZBvXP88WLLduOBP5unFdbDtgkT2tAOLmFFiPK2QuaJiTo7iqGfZzGwWOEyyizOFgHJa4LJGdfwDeb0bKOlscJgFfZi7Fl0Bh8LmuyLR43DqBfF53Ys6TjWDZu644takdqRncDiT6tnlo3zU/xyRh2VaZOhV3ZUzlEzDXLTSqR0exZRvRiF apattar@IB-QQCQWG9YF3" }]
    }
    check = {
      "nios.auth_method" = "KEYPAIR_PASSWORD"
    }
  }

}

case "auth_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_type    = "LOCAL"
    }
    check = {
      "nios.auth_type" = "LOCAL"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_type    = "REMOTE"
    }
    check = {
      "nios.auth_type" = "REMOTE"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_type    = "SAML"
    }
    check = {
      "nios.auth_type" = "SAML"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      auth_type    = "SAML_LOCAL"
    }
    check = {
      "nios.auth_type" = "SAML_LOCAL"
    }
  }

}

case "ca_certificate_issuer" {
  backend            = "nios"
  skip_if_env_empty  = ["NIOS_CA_CERT1_REF"]
  skip_reason        = "NIOS_CA_CERT1_REF environment variable must be set (run integration test setup to upload CA certs)"

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = true
      ca_certificate_issuer             = "{{nios_ca_cert1_ref}}"
      client_certificate_serial_number  = "{{nios_ca_cert1_serial}}"
    }
    check = {
      "nios.ca_certificate_issuer"            = "{{nios_ca_cert1_ref}}"
      "nios.enable_certificate_authentication" = "true"
    }
  }

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = true
      ca_certificate_issuer             = "{{nios_ca_cert2_ref}}"
      client_certificate_serial_number  = "{{nios_ca_cert2_serial}}"
    }
    check = {
      "nios.ca_certificate_issuer" = "{{nios_ca_cert2_ref}}"
    }
  }

}

case "client_certificate_serial_number" {
  backend           = "nios"
  parallel          = true
  skip_if_env_empty = ["NIOS_CA_CERT1_REF"]
  skip_reason       = "NIOS_CA_CERT1_REF environment variable must be set (run integration test setup to upload CA certs)"

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = true
      ca_certificate_issuer             = "{{nios_ca_cert1_ref}}"
      client_certificate_serial_number  = "{{nios_ca_cert1_serial}}"
    }
    check = {
      "nios.client_certificate_serial_number" = "{{nios_ca_cert1_serial}}"
    }
  }

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = true
      ca_certificate_issuer             = "{{nios_ca_cert2_ref}}"
      client_certificate_serial_number  = "{{nios_ca_cert2_serial}}"
    }
    check = {
      "nios.client_certificate_serial_number" = "{{nios_ca_cert2_serial}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      comment      = "example admin user"
    }
    check = {
      "nios.comment" = "example admin user"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      comment      = "example admin user updated"
    }
    check = {
      "nios.comment" = "example admin user updated"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      disable      = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      disable      = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "email" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      email        = "abc@example.com"
    }
    check = {
      "nios.email" = "abc@example.com"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      email        = "xyz@sample.com"
    }
    check = {
      "nios.email" = "xyz@sample.com"
    }
  }

}

case "enable_certificate_authentication" {
  backend           = "nios"
  skip_if_env_empty = ["NIOS_CA_CERT1_REF"]
  skip_reason       = "NIOS_CA_CERT1_REF environment variable must be set (run integration test setup to upload CA certs)"

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = true
      ca_certificate_issuer             = "{{nios_ca_cert1_ref}}"
      client_certificate_serial_number  = "{{nios_ca_cert1_serial}}"
    }
    check = {
      "nios.enable_certificate_authentication" = "true"
    }
  }

  step {
    nios {
      name                              = "{{random}}"
      password                          = "Example-Admin123!"
      admin_groups                      = ["admin-group"]
      enable_certificate_authentication = false
    }
    check = {
      "nios.enable_certificate_authentication" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      ext_attrs    = { Site = "{{random3}}" }
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
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name         = "{{random2}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "password" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-password123!"
      admin_groups = ["admin-group"]
    }
  }

}

case "ssh_keys" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      ssh_keys     = [{ key_name = "sample-key", key_type = "RSA", key_value = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4duV7U2EWRn/Auxa+iX2nx/ffrH249xzDw2fUvP8aeDskP+M1gDJnNZG21Ty2HH1kMZ1qRpN3DY0YBRYopA5P3EFfkq4La8OdG6IH47CINdqCHQTKGZ40YCdCqqz9hd796rrQI1e9yY08MLob79M7jPZLsafDoB7sa4ZBvXP88WLLduOBP5unFdbDtgkT2tAOLmFFiPK2QuaJiTo7iqGfZzGwWOEyyizOFgHJa4LJGdfwDeb0bKOlscJgFfZi7Fl0Bh8LmuyLR43DqBfF53Ys6TjWDZu644takdqRncDiT6tnlo3zU/xyRh2VaZOhV3ZUzlEzDXLTSqR0exZRvRiF apattar@IB-QQCQWG9YF3" }]
    }
    check = {
      "nios.ssh_keys.0.key_name" = "sample-key"
      "nios.ssh_keys.0.key_type" = "RSA"
    }
  }

}

case "time_zone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      time_zone    = "UTC"
    }
    check = {
      "nios.time_zone" = "UTC"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
      time_zone    = "Singapore"
    }
    check = {
      "nios.time_zone" = "Singapore"
    }
  }

}
