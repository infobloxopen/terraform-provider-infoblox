package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ValidateScheduleConfig validates schedule configuration for various schedule types (blackout_schedule, recurring_schedule, etc.)
func ValidateScheduleConfig(settingObj types.Object, scheduleAttrName string, basePath path.Path, diagnostics *diag.Diagnostics) {
	if settingObj.IsNull() || settingObj.IsUnknown() {
		return
	}

	var scheduleObj types.Object
	var schedulePath path.Path

	if scheduleAttrName == "" {
		// Direct schedule at root level (e.g., vdiscoverytask.scheduled_run)
		scheduleObj = settingObj
		schedulePath = basePath
	} else {
		// Nested schedule (e.g., discovery_blackout_setting.blackout_schedule)
		attrs := settingObj.Attributes()
		scheduleAttr, exists := attrs[scheduleAttrName]

		if !exists || scheduleAttr.IsNull() || scheduleAttr.IsUnknown() {
			return
		}

		var ok bool
		scheduleObj, ok = scheduleAttr.(types.Object)
		if !ok {
			diagnostics.AddAttributeError(
				basePath.AtName(scheduleAttrName),
				"Invalid Schedule Attribute",
				fmt.Sprintf("Expected %s to be an object but got different type", scheduleAttrName),
			)
			return
		}
		schedulePath = basePath.AtName(scheduleAttrName)
	}

	if scheduleObj.IsNull() || scheduleObj.IsUnknown() {
		return
	}

	schedule := scheduleObj.Attributes()
	if schedule == nil {
		return
	}

	recurringTime := schedule["recurring_time"]
	repeat := schedule["repeat"]
	weekdays := schedule["weekdays"]
	frequency := schedule["frequency"]
	every := schedule["every"]
	minutesPastHour := schedule["minutes_past_hour"]
	month := schedule["month"]
	dayOfMonth := schedule["day_of_month"]
	hourOfDay := schedule["hour_of_day"]
	year := schedule["year"]

	// Validate recurring_time conflicts
	if !recurringTime.IsNull() && !recurringTime.IsUnknown() {
		if (!hourOfDay.IsNull() && !hourOfDay.IsUnknown()) ||
			(!year.IsNull() && !year.IsUnknown()) ||
			(!month.IsNull() && !month.IsUnknown()) ||
			(!dayOfMonth.IsNull() && !dayOfMonth.IsUnknown()) {
			diagnostics.AddAttributeError(
				schedulePath.AtName("recurring_time"),
				"Invalid Configuration for "+scheduleAttrName,
				"Cannot set recurring_time if any of hour_of_day, year, month, day_of_month is set",
			)
		}
	}

	// Validate repeat field logic
	if !repeat.IsNull() && !repeat.IsUnknown() {
		repeatStr, ok := repeat.(types.String)
		if !ok {
			diagnostics.AddAttributeError(
				schedulePath.AtName("repeat"),
				"Invalid Repeat Attribute",
				"Expected repeat to be a string but got different type",
			)
			return
		}

		repeatValue := repeatStr.ValueString()
		if repeatValue == "" {
			repeatValue = "ONCE"
		}

		switch repeatValue {
		case "ONCE":
			// For ONCE: cannot set weekdays, frequency, every
			if (!weekdays.IsNull() && !weekdays.IsUnknown()) ||
				(!frequency.IsNull() && !frequency.IsUnknown()) ||
				(!every.IsNull() && !every.IsUnknown()) {
				diagnostics.AddAttributeError(
					schedulePath.AtName("repeat"),
					"Invalid Configuration for Repeat",
					"Cannot set frequency, weekdays and every if repeat is set to ONCE",
				)
				return
			}

			// For ONCE: must set month, day_of_month, hour_of_day, minutes_past_hour
			if month.IsNull() || dayOfMonth.IsNull() || hourOfDay.IsNull() || minutesPastHour.IsNull() {
				diagnostics.AddAttributeError(
					schedulePath.AtName("repeat"),
					"Invalid Configuration for Schedule",
					"If repeat is set to ONCE, then month, day_of_month, hour_of_day and minutes_past_hour must be set",
				)
				return
			}

		case "RECUR":
			// For RECUR: cannot set month, day_of_month, year
			if (!month.IsNull() && !month.IsUnknown()) ||
				(!dayOfMonth.IsNull() && !dayOfMonth.IsUnknown()) ||
				(!year.IsNull() && !year.IsUnknown()) {
				diagnostics.AddAttributeError(
					schedulePath.AtName("repeat"),
					"Invalid Configuration for Repeat",
					"Cannot set month, day_of_month and year if repeat is set to RECUR",
				)
				return
			}

			// For RECUR: must set frequency, hour_of_day, minutes_past_hour
			if frequency.IsNull() ||
				hourOfDay.IsNull() ||
				minutesPastHour.IsNull() {
				diagnostics.AddAttributeError(
					schedulePath.AtName("repeat"),
					"Invalid Configuration for Schedule",
					"If repeat is set to RECUR, then frequency, hour_of_day and minutes_past_hour must be set",
				)
				return
			}

			// Handle weekdays validation based on frequency for RECUR only
			if !frequency.IsNull() && !frequency.IsUnknown() {
				freqStr, ok := frequency.(types.String)
				if ok && freqStr.ValueString() == "WEEKLY" {
					// WEEKLY requires weekdays
					if weekdays.IsNull() {
						diagnostics.AddAttributeError(
							schedulePath.AtName("weekdays"),
							"Invalid Configuration for Weekdays",
							"Weekdays must be set if frequency is set to WEEKLY",
						)
					}
				} else {
					// Non-WEEKLY cannot have weekdays
					if !weekdays.IsNull() && !weekdays.IsUnknown() {
						diagnostics.AddAttributeError(
							schedulePath.AtName("weekdays"),
							"Invalid Configuration for Weekdays",
							"Weekdays can only be set if frequency is set to WEEKLY",
						)
					}
				}
			}
		}
	}
}
