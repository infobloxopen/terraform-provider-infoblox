# CredentialGroup — nios list cases
#
# discovery:credentialgroup has no searchable field (see nios_datasources.hcl),
# so there is no filters case: NIOS rejects any filtered query outright.

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "include_resource" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
  }

}
