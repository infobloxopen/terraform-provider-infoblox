package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6networkcontainerDataSource(t *testing.T) {
	dsType := "infoblox_ipv6networkcontainer"
	resourceType := "infoblox_ipv6networkcontainer"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6networkcontainerExistsNIOS,
			Destroy: testAccCheckIpv6networkcontainerDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/ipv6networkcontainer/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
