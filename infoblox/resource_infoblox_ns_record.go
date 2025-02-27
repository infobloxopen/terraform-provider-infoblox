package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strings"
)

func resourceNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSRecordCreate,
		Read:   resourceNSRecordRead,
		Update: resourceNSRecordUpdate,
		Delete: resourceNSRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceNSRecordImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the NS record in FQDN format. This value can be in unicode format.",
				//write a validation function to check for leading or trailing white space .
				StateFunc: trimWhitespace,
			},
			"nameserver": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name of an authoritative server for the redirected zone.",
				//write a validation function to check for leading or trailing white space .
				StateFunc: trimWhitespace,
			},
			"addresses": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of zone name servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "he address of the Zone Name Server.",
						},
						"auto_create_ptr": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Flag to indicate if ptr records need to be auto created.",
						},
					},
				},
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if newValue == "0" {
						return false
					}
					oldList, newList := d.GetChange("addresses")
					return CompareSortedList(oldList, newList, "address", "auto_create_ptr")
				},
			},
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the DNS view in which the record resides.Example: “external”.",
				//validation function to check for leading or trailing zeros
				StateFunc: trimWhitespace,
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
		},
	}
}
func resourceNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	nameserver := d.Get("nameserver").(string)
	addressesInterface := d.Get("addresses").([]interface{})
	addresses := ConvertInterfaceToZoneNameServers(addressesInterface)
	view := d.Get("view").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	recordNS, err := objMgr.CreateNSRecord(name, nameserver, view, addresses)
	if err != nil {
		return err
	}
	d.SetId(recordNS.Ref)
	if err = d.Set("ref", recordNS.Ref); err != nil {
		return err
	}
	return nil
}

func resourceNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevNameServer, _ := d.GetChange("nameserver")
			prevView, _ := d.GetChange("view")
			prevAddresses, _ := d.GetChange("addresses")
			// TODO: move to the new Terraform plugin framework and
			// process all the errors instead of ignoring them here.
			_ = d.Set("name", prevName.(string))
			_ = d.Set("view", prevView.(string))
			_ = d.Set("nameserver", prevNameServer.(string))
			_ = d.Set("address", prevAddresses)
		}
	}()
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	if d.HasChange("name") {
		return fmt.Errorf("changing the value of 'name' field is not allowed")
	}
	if d.HasChange("view") {
		return fmt.Errorf("changing the value of 'view' field is not allowed")
	}

	nameserver := d.Get("nameserver").(string)
	addressesInterface := d.Get("addresses").([]interface{})
	addresses := ConvertInterfaceToZoneNameServers(addressesInterface)
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	rec, err := objMgr.UpdateNSRecord(
		d.Id(), "", nameserver, "", addresses)
	if err != nil {
		return fmt.Errorf("error updating MX-Record: %s", err)
	}
	updateSuccessful = true
	d.SetId(rec.Ref)
	if err = d.Set("ref", rec.Ref); err != nil {
		return err
	}
	return resourceNSRecordRead(d, m)
}
func resourceNSRecordRead(d *schema.ResourceData, m interface{}) error {
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	recordNS, err := objMgr.GetNSRecordByRef(d.Id())

	if err != nil {
		return fmt.Errorf("failed getting NS-record: %w", err)
	}

	if err = d.Set("name", recordNS.Name); err != nil {
		return err
	}
	if err = d.Set("nameserver", recordNS.Nameserver); err != nil {
		return err
	}
	if err = d.Set("view", recordNS.View); err != nil {
		return err
	}
	if recordNS.Addresses != nil {
		addressInterface := convertZoneNameServersToInterface(recordNS.Addresses)
		if err = d.Set("addresses", addressInterface); err != nil {
			return err
		}
	}
	if err = d.Set("ref", recordNS.Ref); err != nil {
		return err
	}
	return nil
}
func resourceNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	recordNS, err := objMgr.GetNSRecordByRef(d.Id())
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteARecord(recordNS.Ref)
	if err != nil {
		return fmt.Errorf("deletion of NS-record failed: %w", err)
	}
	d.SetId("")

	return nil
}
func resourceNSRecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	obj, err := objMgr.GetNSRecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting NS-record: %w", err)
	}

	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("nameserver", obj.Nameserver); err != nil {
		return nil, err
	}
	if err = d.Set("view", obj.View); err != nil {
		return nil, err
	}
	if obj.Addresses != nil {
		addressInterface := convertZoneNameServersToInterface(obj.Addresses)
		if err = d.Set("addresses", addressInterface); err != nil {
			return nil, err
		}
	}
	d.SetId(obj.Ref)

	// Resource NSRecord update Terraform Internal ID and Ref on NIOS side
	// After the record is imported, call the update function
	err = resourceNSRecordUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
func ConvertInterfaceToZoneNameServers(addressesInterface []interface{}) []*ibclient.ZoneNameServer {
	var zoneNameServers []*ibclient.ZoneNameServer
	for _, addressInterface := range addressesInterface {
		addressMap := addressInterface.(map[string]interface{})
		zoneNameServer := &ibclient.ZoneNameServer{
			Address:       addressMap["address"].(string),
			AutoCreatePtr: addressMap["auto_create_ptr"].(bool),
		}
		zoneNameServers = append(zoneNameServers, zoneNameServer)
	}
	return zoneNameServers
}
func trimWhitespace(val interface{}) string {
	return strings.TrimSpace(val.(string))
}
func convertZoneNameServersToInterface(zoneNameServers []*ibclient.ZoneNameServer) []map[string]interface{} {
	nsInterface := make([]map[string]interface{}, 0, len(zoneNameServers))
	for _, ns := range zoneNameServers {
		nsMap := make(map[string]interface{})
		nsMap["address"] = ns.Address
		nsMap["auto_create_ptr"] = ns.AutoCreatePtr
		nsInterface = append(nsInterface, nsMap)
	}
	return nsInterface
}
