case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.name"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: data source config helper 'testAccNsgroupDelegationDataSourceConfigExtAttrFilters' could not be parsed (no resource/data block found)
case "ext_attr_filters" {
  backend     = "nios"
  skip        = true
  skip_reason = "data source config helper 'testAccNsgroupDelegationDataSourceConfigExtAttrFilters' could not be parsed (no resource/data block found)"
}
