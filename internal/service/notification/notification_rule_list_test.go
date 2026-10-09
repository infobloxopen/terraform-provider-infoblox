package notification_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNotificationRuleList(t *testing.T) {
	resourceType := "infoblox_notification_rule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNotificationRuleExistsNIOS,
			Destroy: testAccCheckNotificationRuleDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "notification/notification_rule/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
