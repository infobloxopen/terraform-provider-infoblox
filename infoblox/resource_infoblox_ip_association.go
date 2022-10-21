package infoblox

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceIpAssociation() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: stateImporter,
		},

		Schema: map[string]*schema.Schema{
			"mac_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "MAC address of a cloud instance.",
			},
			"duid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DHCP unique identifier for IPv6.",
			},
			"enable_dhcp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "The flag which defines if the host record is to be used for IPAM purposes.",
			},
			"internal_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This value must point to the ID of the appropriate allocation resource. Required on resource creation.",
			},
		},
	}
}

// TODO: add validation of values (extra spaces, format, etc)
func resourceIpAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

	return resourceIpAssociationCreateUpdate(d, m)
}

func resourceIpAssociationRead(d *schema.ResourceData, m interface{}) error {
	var (
		err                                                          error
		hostRec                                                      *ibclient.HostRecord
		duid, macAddr                                                string
		duidActual, macAddrActual                                    string
		enableDhcp                                                   bool
		enableDhcpActual, enableDhcpActualIpv4, enableDhcpActualIpv6 bool
		importOp                                                     bool
	)

	hostRec, err = getAndRenewHostRecAltId(d, m, false)
	if err != nil {
		return err
	}

	if macAddrVal, ok := d.GetOk("mac_addr"); ok {
		macAddr = macAddrVal.(string)
	}

	if duidVal, ok := d.GetOk("duid"); ok {
		duid = duidVal.(string)
	}

	if hostRec.Ipv6Addrs != nil && len(hostRec.Ipv6Addrs) > 0 {
		if len(hostRec.Ipv6Addrs) > 1 {
			return fmt.Errorf("association with multiple IP addresses are not supported")
		}

		enableDhcpActualIpv6 = hostRec.Ipv6Addrs[0].EnableDhcp
		duidActual = hostRec.Ipv6Addrs[0].Duid
	}

	if hostRec.Ipv4Addrs != nil && len(hostRec.Ipv4Addrs) > 0 {
		if len(hostRec.Ipv4Addrs) > 1 {
			return fmt.Errorf("association with multiple IP addresses are not supported")
		}

		enableDhcpActualIpv4 = hostRec.Ipv4Addrs[0].EnableDhcp
		macAddrActual = hostRec.Ipv4Addrs[0].Mac
	}

	enableDhcpActual = enableDhcpActualIpv4 || enableDhcpActualIpv6

	if importOp {
		if duid != duidActual || macAddr != macAddrActual || enableDhcp != enableDhcpActual {
			return fmt.Errorf("one of the resource properties is not equal to appropriate NIOS object's value")
		}
	}
	d.Set("duid", duidActual)
	d.Set("mac_addr", macAddrActual)
	d.Set("enable_dhcp", enableDhcpActual)

	return nil
}

func resourceIpAssociationDelete(d *schema.ResourceData, m interface{}) error {
	if err := resourceIpAssociationCreateUpdateCommon(d, m, "00:00:00:00:00:00", ""); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceIpAssociationCreateUpdate(d *schema.ResourceData, m interface{}) error {
	var mac, duid string

	val, ok := d.GetOk("mac_addr")
	if ok {
		mac = val.(string)
	}
	val, ok = d.GetOk("duid")
	if ok {
		duid = val.(string)
	}

	return resourceIpAssociationCreateUpdateCommon(d, m, mac, duid)
}

func resourceIpAssociationCreateUpdateCommon(
	d *schema.ResourceData, m interface{}, mac string, duid string) (err error) {

	var (
		recIpV4Addr *ibclient.HostRecordIpv4Addr
		recIpV6Addr *ibclient.HostRecordIpv6Addr
		ipV4Addr    string
		ipV6Addr    string
		tenantId    string
	)

	objMgr := ibclient.NewObjectManager(m.(ibclient.IBConnector), "Terraform", "")

	internalIdStr := d.Get("internal_id").(string)
	if internalIdStr == "" {
		return fmt.Errorf("internal_id field must not be empty")
	}
	enableDhcp := d.Get("enable_dhcp").(bool)

	hostRec, err := objMgr.SearchHostRecordByAltId(internalIdStr, "", eaNameForInternalId)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return fmt.Errorf(
				"error getting the allocated host record with ID '%s': %s",
				d.Id(), err.Error())
		}
		log.Printf("resource with the ID '%s' has been lost, removing it", d.Id())
		d.SetId("")
		return nil
	}

	if len(hostRec.Ipv4Addrs) > 0 {
		recIpV4Addr = &hostRec.Ipv4Addrs[0]
		ipV4Addr = recIpV4Addr.Ipv4Addr
	}
	if len(hostRec.Ipv6Addrs) > 0 {
		recIpV6Addr = &hostRec.Ipv6Addrs[0]
		ipV6Addr = recIpV6Addr.Ipv6Addr
	}

	mac = strings.Replace(mac, "-", ":", -1)

	if hostRec.Ea != nil {
		if tempVal, found := hostRec.Ea["Tenant ID"]; found {
			if tempStrVal, ok := tempVal.(string); ok {
				tenantId = tempStrVal
			}
		}

	}
	objMgr = ibclient.NewObjectManager(
		m.(ibclient.IBConnector), "Terraform", tenantId)

	_, err = objMgr.UpdateHostRecord(
		hostRec.Ref,
		hostRec.EnableDns,
		enableDhcp,
		hostRec.Name,
		hostRec.NetworkView,
		"", "",
		ipV4Addr, ipV6Addr,
		mac, duid,
		hostRec.UseTtl, hostRec.Ttl,
		hostRec.Comment,
		hostRec.Ea, []string{})
	if err != nil {
		return fmt.Errorf(
			"failed to update the resource with ID '%s' (host record with internal ID '%s'): %s",
			d.Id(), internalIdStr, err.Error())
	}

	// Generate an ID for a newly created resource.
	if d.Id() == "" {
		d.SetId(generateInternalId().String())
	}

	return nil
}

func resourceIpAssociationInit() *schema.Resource {
	association := resourceIpAssociation()
	association.Create = resourceIpAssociationCreateUpdate
	association.Read = resourceIpAssociationRead
	association.Update = resourceIpAssociationUpdate
	association.Delete = resourceIpAssociationDelete

	return association
}
