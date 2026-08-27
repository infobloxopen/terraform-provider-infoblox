package misc

import (
	niosmisc "github.com/infobloxopen/infoblox-nios-go-client/misc"
)

// Infoblox Ruleset model
type Ruleset struct {
	Id   *string
	NIOS *NIOSRulesetExt
}

// NIOSRulesetExt - NIOS specific fields for Ruleset
type NIOSRulesetExt struct {
	Comment       *string
	Disabled      *bool
	Name          *string
	NxdomainRules []niosmisc.RulesetNxdomainRules
	Type          *string
}
