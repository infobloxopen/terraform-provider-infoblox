# Manual resource acceptance-test cases for RecordA (nios).
#
# This file is USER-OWNED and is never overwritten by codegen. Add your
# own custom scenarios here as `case "<name>" { ... }` blocks; they are
# loaded and run automatically next to the generated
# nios_resources.tfvars cases (via acctest.RunResourceCases).
#
# Uncomment and adapt the skeleton below. Placeholders such as {{random}}
# are materialized per-subtest. Reference prerequisite resources declared
# in prerequisites_hcl using unquoted refs.

# case "my_custom_scenario" {
#   backend = "nios"
#
#   # Optional: resources created before the step this case depends on.
#   # prerequisites_hcl = <<-PREREQ
#   #   resource "<resource_type>" "dep" {
#   #     nios = { /* fields */ }
#   #   }
#   # PREREQ
#
#   step {
#     nios {
#       # field = "value"
#     }
#     check = {
#       # "nios.field" = "value"
#     }
#   }
#
#   # Optional second step for update testing:
#   # step {
#   #   nios {
#   #     # field = "updated"
#   #   }
#   #   check = {
#   #     # "nios.field" = "updated"
#   #   }
#   # }
# }
