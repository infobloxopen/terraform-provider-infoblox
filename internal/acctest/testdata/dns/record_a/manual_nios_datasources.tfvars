# Manual data source acceptance-test cases for RecordA (nios).
#
# This file is USER-OWNED and is never overwritten by codegen. Add your
# own custom scenarios here as `case "<name>" { ... }` blocks; they are
# loaded and run automatically next to the generated
# nios_datasources.tfvars cases (via acctest.RunDataSourceCases).
#
# Uncomment and adapt the skeleton below. The `step` block creates the
# resource; the `filter` block queries the data source for it.

# case "my_custom_ds" {
#   backend = "nios"
#
#   filter {
#     type = "filters"
#     values = {
#       # name = "nios.name"
#     }
#   }
#
#   step {
#     nios {
#       # field = "value"
#     }
#   }
# }
