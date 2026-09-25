package rir

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

var ripeEmailRegex = regexp.MustCompile(`^[^@]+@[^@]+\.com$`)
var ripeTechnicalContactRegex = regexp.MustCompile(`^[A-Za-z]{2,4}(?:[1-9][0-9]{0,5})?-[A-Za-z0-9]{1,9}$`)

// validRipeCountries is the set of valid RIPE country values.
var validRipeCountries = map[string]struct{}{
	"Afghanistan (AF)": {}, "Åland Islands (AX)": {}, "Albania (AL)": {}, "Algeria (DZ)": {}, "American Samoa (AS)": {}, "Andorra (AD)": {}, "Angola (AO)": {}, "Anguilla (AI)": {}, "Antarctica (AQ)": {}, "Antigua and Barbuda (AG)": {},
	"Argentina (AR)": {}, "Armenia (AM)": {}, "Aruba (AW)": {}, "Australia (AU)": {}, "Austria (AT)": {}, "Azerbaijan (AZ)": {}, "Bahamas (BS)": {}, "Bahrain (BH)": {}, "Bangladesh (BD)": {}, "Barbados (BB)": {},
	"Belarus (BY)": {}, "Belgium (BE)": {}, "Belize (BZ)": {}, "Benin (BJ)": {}, "Bermuda (BM)": {}, "Bhutan (BT)": {}, "Bolivia, Plurinational State of (BO)": {}, "Bonaire, Sint Eustatius and Saba (BQ)": {}, "Bosnia and Herzegovina (BA)": {}, "Botswana (BW)": {},
	"Bouvet Island (BV)": {}, "Brazil (BR)": {}, "British Indian Ocean Territory (IO)": {}, "Brunei Darussalam (BN)": {}, "Bulgaria (BG)": {}, "Burkina Faso (BF)": {}, "Burundi (BI)": {}, "Cambodia (KH)": {}, "Cameroon (CM)": {}, "Canada (CA)": {},
	"Cape Verde (CV)": {}, "Cayman Islands (KY)": {}, "Central African Republic (CF)": {}, "Chad (TD)": {}, "Chile (CL)": {}, "China (CN)": {}, "Christmas Island (CX)": {}, "Cocos (Keeling) Islands (CC)": {}, "Colombia (CO)": {}, "Comoros (KM)": {},
	"Congo (CG)": {}, "Congo, The Democratic Republic of the (CD)": {}, "Cook Islands (CK)": {}, "Costa Rica (CR)": {}, "Côte d'Ivoire (CI)": {}, "Croatia (HR)": {}, "Cuba (CU)": {}, "Curaçao (CW)": {}, "Cyprus (CY)": {}, "Czech Republic (CZ)": {},
	"Denmark (DK)": {}, "Djibouti (DJ)": {}, "Dominica (DM)": {}, "Dominican Republic (DO)": {}, "Ecuador (EC)": {}, "Egypt (EG)": {}, "El Salvador (SV)": {}, "Equatorial Guinea (GQ)": {}, "Eritrea (ER)": {}, "Estonia (EE)": {},
	"Ethiopia (ET)": {}, "Falkland Islands (Malvinas) (FK)": {}, "Faroe Islands (FO)": {}, "Fiji (FJ)": {}, "Finland (FI)": {}, "France (FR)": {}, "French Guiana (GF)": {}, "French Polynesia (PF)": {}, "French Southern Territories (TF)": {}, "Gabon (GA)": {},
	"Gambia (GM)": {}, "Georgia (GE)": {}, "Germany (DE)": {}, "Ghana (GH)": {}, "Gibraltar (GI)": {}, "Greece (GR)": {}, "Greenland (GL)": {}, "Grenada (GD)": {}, "Guadeloupe (GP)": {}, "Guam (GU)": {},
	"Guatemala (GT)": {}, "Guernsey (GG)": {}, "Guinea (GN)": {}, "Guinea-Bissau (GW)": {}, "Guyana (GY)": {}, "Haiti (HT)": {}, "Heard Island and McDonald Islands (HM)": {}, "Holy See (Vatican City State) (VA)": {}, "Honduras (HN)": {}, "Hong Kong (HK)": {},
	"Hungary (HU)": {}, "Iceland (IS)": {}, "India (IN)": {}, "Indonesia (ID)": {}, "Iran, Islamic Republic of (IR)": {}, "Iraq (IQ)": {}, "Ireland (IE)": {}, "Isle of Man (IM)": {}, "Israel (IL)": {}, "Italy (IT)": {},
	"Jamaica (JM)": {}, "Japan (JP)": {}, "Jersey (JE)": {}, "Jordan (JO)": {}, "Kazakhstan (KZ)": {}, "Kenya (KE)": {}, "Kiribati (KI)": {}, "Korea, Democratic People's Republic of (KP)": {}, "Korea, Republic of (KR)": {}, "Kuwait (KW)": {},
	"Kyrgyzstan (KG)": {}, "Lao People's Democratic Republic (LA)": {}, "Latvia (LV)": {}, "Lebanon (LB)": {}, "Lesotho (LS)": {}, "Liberia (LR)": {}, "Libya (LY)": {}, "Liechtenstein (LI)": {}, "Lithuania (LT)": {}, "Luxembourg (LU)": {},
	"Macao (MO)": {}, "Macedonia, The Former Yugoslav Republic of (MK)": {}, "Madagascar (MG)": {}, "Malawi (MW)": {}, "Malaysia (MY)": {}, "Maldives (MV)": {}, "Mali (ML)": {}, "Malta (MT)": {}, "Marshall Islands (MH)": {}, "Martinique (MQ)": {},
	"Mauritania (MR)": {}, "Mauritius (MU)": {}, "Mayotte (YT)": {}, "Mexico (MX)": {}, "Micronesia, Federated States of (FM)": {}, "Moldova, Republic of (MD)": {}, "Monaco (MC)": {}, "Mongolia (MN)": {}, "Montenegro (ME)": {}, "Montserrat (MS)": {},
	"Morocco (MA)": {}, "Mozambique (MZ)": {}, "Myanmar (MM)": {}, "Namibia (NA)": {}, "Nauru (NR)": {}, "Nepal (NP)": {}, "Netherlands (NL)": {}, "New Caledonia (NC)": {}, "New Zealand (NZ)": {}, "Nicaragua (NI)": {},
	"Niger (NE)": {}, "Nigeria (NG)": {}, "Niue (NU)": {}, "Norfolk Island (NF)": {}, "Northern Mariana Islands (MP)": {}, "Norway (NO)": {}, "Oman (OM)": {}, "Pakistan (PK)": {}, "Palau (PW)": {}, "Palestinian Territory, Occupied (PS)": {},
	"Panama (PA)": {}, "Papua New Guinea (PG)": {}, "Paraguay (PY)": {}, "Peru (PE)": {}, "Philippines (PH)": {}, "Pitcairn (PN)": {}, "Poland (PL)": {}, "Portugal (PT)": {}, "Puerto Rico (PR)": {}, "Qatar (QA)": {},
	"Réunion (RE)": {}, "Romania (RO)": {}, "Russian Federation (RU)": {}, "Rwanda (RW)": {}, "Saint Barthélemy (BL)": {}, "Saint Helena, Ascension and Tristan da Cunha (SH)": {}, "Saint Kitts and Nevis (KN)": {}, "Saint Lucia (LC)": {}, "Saint Martin (French part) (MF)": {}, "Saint Pierre and Miquelon (PM)": {},
	"Saint Vincent and the Grenadines (VC)": {}, "Samoa (WS)": {}, "San Marino (SM)": {}, "Sao Tome and Principe (ST)": {}, "Saudi Arabia (SA)": {}, "Senegal (SN)": {}, "Serbia (RS)": {}, "Seychelles (SC)": {}, "Sierra Leone (SL)": {}, "Singapore (SG)": {},
	"Sint Maarten (Dutch part) (SX)": {}, "Slovakia (SK)": {}, "Slovenia (SI)": {}, "Solomon Islands (SB)": {}, "Somalia (SO)": {}, "South Africa (ZA)": {}, "South Georgia and the South Sandwich Islands (GS)": {}, "South Sudan (SS)": {}, "Spain (ES)": {}, "Sri Lanka (LK)": {},
	"Sudan (SD)": {}, "Suriname (SR)": {}, "Svalbard and Jan Mayen (SJ)": {}, "Swaziland (SZ)": {}, "Sweden (SE)": {}, "Switzerland (CH)": {}, "Syrian Arab Republic (SY)": {}, "Taiwan, Province of China (TW)": {}, "Tajikistan (TJ)": {}, "Tanzania, United Republic of (TZ)": {},
	"Thailand (TH)": {}, "Timor-Leste (TL)": {}, "Togo (TG)": {}, "Tokelau (TK)": {}, "Tonga (TO)": {}, "Trinidad and Tobago (TT)": {}, "Tunisia (TN)": {}, "Turkey (TR)": {}, "Turkmenistan (TM)": {}, "Turks and Caicos Islands (TC)": {},
	"Tuvalu (TV)": {}, "Uganda (UG)": {}, "Ukraine (UA)": {}, "United Arab Emirates (AE)": {}, "United Kingdom (GB)": {}, "United States (US)": {}, "United States Minor Outlying Islands (UM)": {}, "Uruguay (UY)": {}, "Uzbekistan (UZ)": {}, "Vanuatu (VU)": {},
	"Venezuela, Bolivarian Republic of (VE)": {}, "Viet Nam (VN)": {}, "Virgin Islands, British (VG)": {}, "Virgin Islands, U.S. (VI)": {}, "Wallis and Futuna (WF)": {}, "Western Sahara (EH)": {}, "Yemen (YE)": {}, "Zambia (ZM)": {}, "Zimbabwe (ZW)": {},
}

// validRipeOrgTypes is the set of valid RIPE organization types.
var validRipeOrgTypes = map[string]struct{}{
	"IANA":              {},
	"RIR":               {},
	"NIR":               {},
	"LIR":               {},
	"WHITEPAGES":        {},
	"DIRECT_ASSIGNMENT": {},
	"OTHER":             {},
}

// ValidateRirOrganization validates the RirOrganization configuration.
func ValidateRirOrganization(ctx context.Context, data RirOrganizationModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRirOrganizationModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRirOrganizationNIOSConfig(ctx, nios, resp)
	}
}

// validateRirOrganizationNIOSConfig validates RIPE-specific ext_attrs values when rir == "RIPE".
// Mirrors the legacy nios provider's ValidateConfig logic.
func validateRirOrganizationNIOSConfig(ctx context.Context, m *NIOSRirOrganizationModel, resp *resource.ValidateConfigResponse) {
	if m.Rir.IsUnknown() || m.Rir.IsNull() || m.Rir.ValueString() != "RIPE" {
		return
	}

	if m.ExtAttrs.IsUnknown() || m.ExtAttrs.IsNull() {
		return
	}

	var extattrsMap map[string]string
	resp.Diagnostics.Append(m.ExtAttrs.ElementsAs(ctx, &extattrsMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for key, value := range extattrsMap {
		switch key {
		case "RIPE Country":
			if value == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("ext_attrs"),
					"Invalid RIPE Country",
					"RIPE Country cannot be empty.",
				)
			} else if _, ok := validRipeCountries[value]; !ok {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("ext_attrs"),
					"Invalid RIPE Country",
					fmt.Sprintf("RIPE Country '%s' is not a valid option.", value),
				)
			}
		case "RIPE Email":
			if !ripeEmailRegex.MatchString(value) {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("ext_attrs"),
					"Invalid RIPE Email",
					fmt.Sprintf("RIPE Email '%s' is not a valid .com email address.", value),
				)
			}
		case "RIPE Technical Contact":
			if !ripeTechnicalContactRegex.MatchString(value) {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("ext_attrs"),
					"Invalid RIPE Technical Contact",
					fmt.Sprintf("RIPE Technical Contact '%s' is not a valid value. Valid format is 'AB123-XYZ'.", value),
				)
			}
		case "RIPE Organization Type":
			if value != "" {
				if _, ok := validRipeOrgTypes[value]; !ok {
					resp.Diagnostics.AddAttributeError(
						path.Root("nios").AtName("ext_attrs"),
						"Invalid RIPE Organization Type",
						fmt.Sprintf("RIPE Organization Type '%s' is not a valid option.", value),
					)
				}
			}
		}
	}
}

// PostFlattenRirOrganizationNIOS copies write-only fields from the plan back to the flattened
// state, since the NIOS API never echoes password back in responses.
func PostFlattenRirOrganizationNIOS(ctx context.Context, planned, flattened *NIOSRirOrganizationModel, diags *diag.Diagnostics) {
	if planned != nil {
		flattened.Password = planned.Password
	}
}
