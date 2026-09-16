# Auto-generated datasource acceptance-test cases for Bfdtemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.authentication_key_id", "nios.authentication_type", "nios.detection_multiplier", "nios.min_rx_interval", "nios.min_tx_interval", "nios.name"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}
