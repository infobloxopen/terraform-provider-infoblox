package notification_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNotificationRestEndpointList(t *testing.T) {
	resourceType := "infoblox_notification_rest_endpoint"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNotificationRestEndpointExistsNIOS,
			Destroy: testAccCheckNotificationRestEndpointDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "notification/notification_rest_endpoint/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
