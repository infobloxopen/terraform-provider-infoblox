// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timetypes

import "github.com/hashicorp/terraform-plugin-framework/diag"

// rfc3339InvalidStringDiagnostic returns an error diagnostic intended to report
// when a string is not RFC3339 format.
func rfc3339InvalidStringDiagnostic(value string, err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Invalid RFC3339 String Value",
		"A string value was provided that is not valid RFC3339 string format.\n\n"+
			"Given Value: "+value+"\n"+
			"Error: "+err.Error(),
	)
}

// goDurationInvalidStringDiagnostic returns an error diagnostic intended to report
// when a string is not a time duration.
func goDurationInvalidStringDiagnostic(value string, err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Invalid Time Duration String Value",
		"A string value was provided that is not a valid Go Time Duration string format. "+
			`A duration string is a sequence of numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". `+
			`Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".\n\n`+
			"Given Value: "+value+"\n"+
			"Error: "+err.Error(),
	)
}
