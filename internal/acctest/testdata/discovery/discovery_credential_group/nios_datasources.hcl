# Auto-generated datasource acceptance-test cases for DiscoveryCredentialgroup.
# Unfiltered read: NIOS reports no searchable fields on this object.

case "read_all" {
  backend = "nios"

  # Assert the created resource appears somewhere in the unfiltered results list.
  # Index-free: walks every `results.*.nios.name` looking for one that matches.
  match_result = ["nios.name"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}
