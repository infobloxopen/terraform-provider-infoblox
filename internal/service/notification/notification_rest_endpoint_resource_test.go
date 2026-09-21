package notification_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNotificationRestEndpointResource(t *testing.T) {
	resourceType := "infoblox_notification_rest_endpoint"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:     testAccCheckNotificationRestEndpointExistsNIOS,
			Destroy:    testAccCheckNotificationRestEndpointDestroyNIOS,
			Disappears: testAccCheckNotificationRestEndpointDisappearsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "notification/notification_rest_endpoint/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

func testAccCheckNotificationRestEndpointExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		conn := acctest.NIOSClient
		res, _, err := conn.NotificationAPI.NotificationRestEndpointAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to read NotificationRestEndpoint: %w", err)
		}
		if res == nil {
			return fmt.Errorf("NotificationRestEndpoint not found: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNotificationRestEndpointDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.NIOSClient
		for name, rs := range s.RootModule().Resources {
			// Skipping data source entries as this is already checked in the resource destroy call.
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			_, httpRes, err := conn.NotificationAPI.NotificationRestEndpointAPI.Read(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
			if err != nil {
				if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}
			return fmt.Errorf("NotificationRestEndpoint still exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckNotificationRestEndpointDisappearsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn := acctest.NIOSClient
		_, err := conn.NotificationAPI.NotificationRestEndpointAPI.Delete(context.Background(), acctest.ExtractNIOSRef(rs.Primary.ID)).Execute()
		if err != nil {
			return fmt.Errorf("failed to delete NotificationRestEndpoint: %w", err)
		}
		return nil
	}
}
