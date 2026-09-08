package notification_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNotificationRestEndpointDataSource(t *testing.T) {
	dsType := "infoblox_notification_rest_endpoint"
	resourceType := "infoblox_notification_rest_endpoint"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNotificationRestEndpointExistsNIOS,
			Destroy: testAccCheckNotificationRestEndpointDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "notification/notification_rest_endpoint/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
