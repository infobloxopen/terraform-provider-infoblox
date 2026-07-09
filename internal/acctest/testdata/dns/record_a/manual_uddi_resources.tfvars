# Manual resource acceptance-test cases for RecordA (uddi).
#
# This file is USER-OWNED and is never overwritten by codegen. Add your
# own custom scenarios here as `case "<name>" { ... }` blocks; they are
# loaded and run automatically next to the generated
# uddi_resources.tfvars cases (via acctest.RunResourceCases).
#
# Uncomment and adapt the skeleton below. Placeholders such as {{random}}
# are materialized per-subtest. Reference prerequisite resources declared
# in prerequisites_hcl using unquoted refs.

# case "my_custom_scenario" {
#   backend = "uddi"
#
#   # Optional: resources created before the step this case depends on.
#   # prerequisites_hcl = <<-PREREQ
#   #   resource "<resource_type>" "dep" {
#   #     uddi = { /* fields */ }
#   #   }
#   # PREREQ
#
#   step {
#     uddi {
#       # field = "value"
#     }
#     check = {
#       # "uddi.field" = "value"
#     }
#   }
#
#   # Optional second step for update testing:
#   # step {
#   #   uddi {
#   #     # field = "updated"
#   #   }
#   #   check = {
#   #     # "uddi.field" = "updated"
#   #   }
#   # }
# }
