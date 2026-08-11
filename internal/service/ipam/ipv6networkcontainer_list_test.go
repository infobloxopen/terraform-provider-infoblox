package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6networkcontainerList(t *testing.T) {
	resourceType := "infoblox_ipv6networkcontainer"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6networkcontainerExistsNIOS,
			Destroy: testAccCheckIpv6networkcontainerDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/ipv6networkcontainer/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
