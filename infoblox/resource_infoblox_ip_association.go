package infoblox

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceIpAssociation() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: ipAssociationImporter,
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
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
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
		err                       error
		hostRec                   *ibclient.HostRecord
		duidActual, macAddrActual string
		enableDhcpActual          bool
		enableDhcpActualIpv4      bool
		enableDhcpActualIpv6      bool
	)

	hostRec, err = getOrFindHostRec(d, m)
	if err != nil {
		return err
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

	if err = d.Set("ref", hostRec.Ref); err != nil {
		return err
	}
	if err = d.Set("duid", duidActual); err != nil {
		return err
	}
	if err = d.Set("mac_addr", macAddrActual); err != nil {
		return err
	}
	if err = d.Set("enable_dhcp", enableDhcpActual); err != nil {
		return err
	}

	return nil
}

func resourceIpAssociationDelete(d *schema.ResourceData, m interface{}) error {
	// TODO: process carefully the case: the host record is already deleted
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
		hostRec     *ibclient.HostRecord
		recIpV4Addr *ibclient.HostRecordIpv4Addr
		recIpV6Addr *ibclient.HostRecordIpv6Addr
		ipV4Addr    string
		ipV6Addr    string
		tenantId    string
	)

	hostRec, err = getOrFindHostRec(d, m)
	if err != nil {
		return err
	}

	internalIdStr := d.Get("internal_id").(string)
	if internalIdStr == "" {
		return fmt.Errorf("internal_id field must not be empty")
	}
	enableDhcp := d.Get("enable_dhcp").(bool)

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
		if tempVal, found := hostRec.Ea[eaNameForTenantId]; found {
			if tempStrVal, ok := tempVal.(string); ok {
				tenantId = tempStrVal
			}
		}

	}
	objMgr := ibclient.NewObjectManager(
		m.(ibclient.IBConnector), "Terraform", tenantId)

	_, err = objMgr.UpdateHostRecord(
		hostRec.Ref,
		hostRec.EnableDns,
		enableDhcp,
		hostRec.Name,
		hostRec.NetworkView,
		hostRec.View,
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

	if err = d.Set("ref", hostRec.Ref); err != nil {
		return err
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

func ipAssociationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	internalId := newInternalResourceIdFromString(d.Id())
	if internalId == nil {
		return nil, fmt.Errorf("ID value provided is not in a proper format")
	}

	d.SetId(internalId.String())
	if err := d.Set("internal_id", internalId.String()); err != nil {
		return nil, err
	}
	if _, err := getOrFindHostRec(d, m); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
