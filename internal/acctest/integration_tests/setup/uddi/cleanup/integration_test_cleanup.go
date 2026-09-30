// Package main provides a standalone cleanup utility that removes resources
// created by the UDDI integration-test setup program.
//
// Objects deleted by this utility:
//
// Anycast Service:
//   - Service whose ID is stored in UDDI_ANYCAST_SERVICE_ID_1
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"
)

func cleanupAnyCastService(ctx context.Context, client *uddiclient.APIClient) {
	serviceID := strings.TrimSpace(os.Getenv("UDDI_ANYCAST_SERVICE_ID_1"))
	if serviceID == "" {
		fmt.Println("cleanup: UDDI_ANYCAST_SERVICE_ID_1 is not set, skipping anycast service deletion")
		return
	}

	_, err := client.InfraManagementAPI.ServicesAPI.Delete(ctx, serviceID).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to delete anycast service (ID=%q): %v\n", serviceID, err)
		return
	}

	fmt.Printf("cleanup: deleted anycast service (ID=%q)\n", serviceID)
}

func main() {
	cspURL := strings.TrimSpace(os.Getenv("INFOBLOX_PORTAL_URL"))
	apiKey := strings.TrimSpace(os.Getenv("INFOBLOX_PORTAL_KEY"))

	if cspURL == "" || apiKey == "" {
		fmt.Println("Missing required UDDI configuration.")
		fmt.Println("Supported env vars: INFOBLOX_PORTAL_URL, INFOBLOX_PORTAL_KEY")
		os.Exit(1)
	}

	client := uddiclient.NewAPIClient(
		uddioption.WithClientName("terraform-integration-test-cleanup"),
		uddioption.WithCSPUrl(cspURL),
		uddioption.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	fmt.Println("--- Cleaning up Anycast Service (UDDI_ANYCAST_SERVICE_ID_1) ---")
	cleanupAnyCastService(ctx, client)

	fmt.Println("Cleanup complete.")
}
