# Auto-generated datasource acceptance-test cases for CredentialGroup.
#
# discovery:credentialgroup has no searchable field: the WAPI schema reports
# searchable_by="" for name, its only field, and any filtered read comes back
# as "AdmConProtoError: Field is not searchable: name". An unfiltered read is
# therefore the only form this data source supports, so the case omits the
# filter block.

case "read_all" {
  backend = "nios"

  step {
    nios {
      name = "{{random}}"
    }
  }

}
