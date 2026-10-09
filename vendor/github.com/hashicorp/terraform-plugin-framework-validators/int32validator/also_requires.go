// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package int32validator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/internal/schemavalidator"
)

// AlsoRequires checks that a set of path.Expression has a non-null value,
// if the current attribute also has a non-null value.
//
// This implements the validation logic declaratively within the schema.
// Refer to [datasourcevalidator.RequiredTogether],
// [providervalidator.RequiredTogether], or [resourcevalidator.RequiredTogether]
// for declaring this type of validation outside the schema definition.
//
// Relative path.Expression will be resolved using the attribute being
// validated.
func AlsoRequires(expressions ...path.Expression) validator.Int32 {
	return schemavalidator.AlsoRequiresValidator{
		PathExpressions: expressions,
	}
}
