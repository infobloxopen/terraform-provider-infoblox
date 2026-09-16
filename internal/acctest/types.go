package acctest

import "github.com/hashicorp/terraform-plugin-testing/helper/resource"

// CaseConfig represents the parsed case file configuration.
type CaseConfig struct {
	Backend          string
	Parallel         bool
	Common           map[string]any
	NIOS             map[string]any
	UDDI             map[string]any
	PrerequisitesHCL string
	DSFilterField    string
}

// CheckFuncs contains resource-specific check functions.
type CheckFuncs struct {
	Exists     func(resourceName string) resource.TestCheckFunc
	Destroy    func(resourceType string) resource.TestCheckFunc
	Disappears func(resourceName string) resource.TestCheckFunc
}
