package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strings"
)

func convertDtcServerMonitorsToInterface(monitors []*ibclient.DtcServerMonitor, connector ibclient.IBConnector) []map[string]interface{} {
	monitorsInterface := make([]map[string]interface{}, 0, len(monitors))
	for _, monitor := range monitors {
		monitorMap := make(map[string]interface{})
		var monitorResult ibclient.DtcMonitorHttp
		err := connector.GetObject(&ibclient.DtcMonitorHttp{}, monitor.Monitor, nil, &monitorResult)
		if err != nil {
			return nil
		}
		referenceParts := strings.Split(monitor.Monitor, ":")
		monitorType := strings.Split(referenceParts[2], "/")[0]
		monitorMap["monitor_name"] = monitorResult.Name
		monitorMap["monitor_type"] = monitorType
		monitorMap["host"] = monitor.Host
		monitorsInterface = append(monitorsInterface, monitorMap)
	}
	return monitorsInterface
}

func resourceDtcServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDtcServerCreate,
		Read:   resourceDtcServerGet,
		Update: resourceDtcServerUpdate,
		Delete: resourceDtcServerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDtcServerImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			return nil
		},
		Schema: map[string]*schema.Schema{
			"auto_create_host_record": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes.\n\n",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the Dtc server.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines if the zone is disabled or not.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the  Dtc Server to be added/updated, as a map in JSON format",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address or FQDN of the server.",
			},
			"monitors": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of IP/FQDN and monitor pairs to be used for additional monitoring.\n\n",
				Elem: &schema.Resource{
					//check the required part once
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address or FQDN of the server used for monitoring.",
						},
						"monitor_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The monitor name related to server.",
						},
						"monitor_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The monitor type related to server.",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The DTC Server display name.",
			},
			"sni_hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hostname for Server Name Indication (SNI) in FQDN format.",
			},
			"use_sni_hostname": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Use flag for: sni_hostname",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
		},
	}
}

func resourceDtcServerCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	host := d.Get("host").(string)
	AutoCreateHostRecord := d.Get("auto_create_host_record").(bool)
	Disable := d.Get("disable").(bool)
	sniHostname := d.Get("sni_hostname").(string)
	useSniHostname := d.Get("use_sni_hostname").(bool)
	extAttrJSON := d.Get("ext_attrs").(string)
	monitors := d.Get("monitors").([]interface{})
	dtcServerMonitor := convertInterfaceToList(monitors)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	if tempVal, found := extAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newDtcServer, err := objMgr.CreateDtcServer(comment, name, host, AutoCreateHostRecord, Disable, extAttrs, dtcServerMonitor, sniHostname, useSniHostname)
	if err != nil {
		return err
	}
	d.SetId(newDtcServer.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newDtcServer.Ref); err != nil {
		return err
	}
	return resourceDtcServerGet(d, m)
}

func resourceDtcServerGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)
	rec, err := searchObjectByRefOrInternalId("DtcServer", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}
	var dtcServer *ibclient.DtcServer
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC Server : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcServer)
	if err != nil {
		return fmt.Errorf("failed getting DTC Server : %s", err.Error())
	}
	delete(dtcServer.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(dtcServer.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("name", dtcServer.Name); err != nil {
		return err
	}
	if err = d.Set("comment", dtcServer.Comment); err != nil {
		return err
	}
	if err = d.Set("disable", dtcServer.Disable); err != nil {
		return err
	}
	if err = d.Set("host", dtcServer.Host); err != nil {
		return err
	}
	monitorInterface := convertDtcServerMonitorsToInterface(dtcServer.Monitors, connector)
	if err = d.Set("monitors", monitorInterface); err != nil {
		return err
	}
	if err = d.Set("auto_create_host_record", dtcServer.AutoCreateHostRecord); err != nil {
		return err
	}
	if err = d.Set("sni_hostname", dtcServer.SniHostname); err != nil {
		return err
	}
	if err = d.Set("use_sni_hostname", dtcServer.UseSniHostname); err != nil {
		return err
	}
	if err = d.Set("ref", dtcServer.Ref); err != nil {
		return err
	}
	d.SetId(dtcServer.Ref)
	return nil
}

func resourceDtcServerUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevEa, _ := d.GetChange("ext_attrs")
			prevMonitors, _ := d.GetChange("monitors")
			prevSniHostname, _ := d.GetChange("sni_hostname")
			prevUseSniHostname, _ := d.GetChange("use_sni_hostname")
			PrevAutoCreateHostRecord, _ := d.GetChange("auto_create_host_record")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("name", prevName.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("ext_attrs", prevEa.(string))
			_ = d.Set("monitors", prevMonitors)
			_ = d.Set("sni_hostname", prevSniHostname.(string))
			_ = d.Set("use_sni_hostname", prevUseSniHostname.(bool))
			_ = d.Set("auto_create_host_record", PrevAutoCreateHostRecord.(bool))

		}
	}()
	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	host := d.Get("host").(string)
	AutoCreateHostRecord := d.Get("auto_create_host_record").(bool)
	Disable := d.Get("disable").(bool)
	sniHostname := d.Get("sni_hostname").(string)
	useSniHostname := d.Get("use_sni_hostname").(bool)
	monitors := d.Get("monitors").([]interface{})
	dtcServerMonitor := convertInterfaceToList(monitors)
	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return err
	}
	var tenantID string
	if tempVal, found := newExtAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	var dtcServer *ibclient.DtcServer

	rec, err := searchObjectByRefOrInternalId("DtcServer", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal Dtc Server : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcServer)
	if err != nil {
		return fmt.Errorf("failed getting Dtc Server : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(dtcServer.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}
	dtcServer, err = objMgr.UpdateDtcServer(d.Id(), comment, name, host, AutoCreateHostRecord, Disable, newExtAttrs, dtcServerMonitor, sniHostname, useSniHostname)
	if err != nil {
		return fmt.Errorf("error updating dtc-server: %w", err)
	}
	updateSuccessful = true
	d.SetId(dtcServer.Ref)
	if err = d.Set("ref", dtcServer.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	return resourceDtcServerGet(d, m)
}

func resourceDtcServerDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := searchObjectByRefOrInternalId("DtcServer", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var DtcServer *ibclient.DtcServer
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &DtcServer)

	_, err = objMgr.DeleteDtcServer(DtcServer.Ref)
	if err != nil {
		return fmt.Errorf("deletion of Dtc Server failed: %w", err)
	}
	d.SetId("")

	return nil
}

func resourceDtcServerImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	obj, err := objMgr.GetDtcServerByRef(d.Id())

	if err != nil {
		return nil, fmt.Errorf("getting DtcServer with ID: %s failed: %w", d.Id(), err)
	}

	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}
	if err = d.Set("disable", obj.Disable); err != nil {
		return nil, err
	}
	if err = d.Set("host", obj.Host); err != nil {
		return nil, err
	}
	monitorInterface := convertDtcServerMonitorsToInterface(obj.Monitors, connector)
	if err = d.Set("monitors", monitorInterface); err != nil {
		return nil, err
	}
	if err = d.Set("auto_create_host_record", obj.AutoCreateHostRecord); err != nil {
		return nil, err
	}
	if err = d.Set("sni_hostname", obj.SniHostname); err != nil {
		return nil, err
	}
	if err = d.Set("use_sni_hostname", obj.UseSniHostname); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)
	err = resourceDtcServerUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
