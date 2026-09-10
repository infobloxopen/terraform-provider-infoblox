package dns_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIPAssociationResource(t *testing.T) {
	resourceType := "infoblox_ip_association"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIPAssociationExistsNIOS,
			Destroy: testAccCheckIPAssociationDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunResourceCases(t, resourceType, "dns/ip_association/"+backend+"_resources.hcl", checksByBackend)
		})
	}
}

// testAccCheckIPAssociationExistsNIOS checks the association's settings actually
// landed on the host record, rather than trusting state. There is no Disappears
// counterpart: the association has no object of its own to delete out of band.
func testAccCheckIPAssociationExistsNIOS(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		ref := rs.Primary.Attributes["nios.record_host_id"]
		if ref == "" {
			return fmt.Errorf("nios.record_host_id is not set on %s", resourceName)
		}

		got, err := associatedDHCPIdentifiers(ref)
		if err != nil {
			return err
		}
		// Only the identifier a case configured is asserted; the other family is
		// absent from that record.
		for attr, actual := range got {
			if want := rs.Primary.Attributes[attr]; want != "" && want != actual {
				return fmt.Errorf("host record %s has %s %q, want %q", ref, attr, actual, want)
			}
		}
		return nil
	}
}

// testAccCheckIPAssociationDestroyNIOS asserts the inverse of the usual destroy
// check: destroying this resource must clear the DHCP settings and leave the host
// record standing, because the record's lifecycle belongs to infoblox_record_host.
func testAccCheckIPAssociationDestroyNIOS(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != resourceType || strings.HasPrefix(name, "data.") {
				continue
			}
			ref := rs.Primary.Attributes["nios.record_host_id"]
			if ref == "" {
				continue
			}

			got, err := associatedDHCPIdentifiers(ref)
			if err != nil {
				// The host record went too, which is fine: the case tears both down.
				return nil
			}
			for attr, actual := range got {
				if actual != "" {
					return fmt.Errorf("host record %s still carries %s %q after the association was destroyed", ref, attr, actual)
				}
			}
		}
		return nil
	}
}

// associatedDHCPIdentifiers reads back the identifiers the association writes,
// keyed by the state attribute each one corresponds to.
func associatedDHCPIdentifiers(ref string) (map[string]string, error) {
	res, _, err := acctest.NIOSClient.DNSAPI.RecordHostAPI.
		Read(context.Background(), acctest.ExtractNIOSRef(ref)).
		ReturnFieldsPlus("ipv4addrs,ipv6addrs").
		ReturnAsObject(1).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to read host record %s: %w", ref, err)
	}

	host := res.GetRecordHostResponseObjectAsResult.GetResult()
	got := map[string]string{"nios.mac": "", "nios.duid": ""}

	if len(host.Ipv4addrs) > 0 && host.Ipv4addrs[0].Mac != nil {
		got["nios.mac"] = *host.Ipv4addrs[0].Mac
	}
	if len(host.Ipv6addrs) > 0 {
		v6 := host.Ipv6addrs[0]
		if v6.Duid != nil {
			got["nios.duid"] = *v6.Duid
		}
		if got["nios.mac"] == "" && v6.Mac != nil {
			got["nios.mac"] = *v6.Mac
		}
	}
	return got, nil
}
