package notification_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNotificationRuleDataSource(t *testing.T) {
	dsType := "infoblox_notification_rule"
	resourceType := "infoblox_notification_rule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNotificationRuleExistsNIOS,
			Destroy: testAccCheckNotificationRuleDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "notification/notification_rule/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
