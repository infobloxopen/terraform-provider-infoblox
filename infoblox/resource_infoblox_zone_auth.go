package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneAuthCreate,
		Read:   resourceZoneAuthGet,
		Update: resourceZoneAuthUpdate,
		Delete: resourceZoneAuthDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The FQDN of the Authoritative zone",
			},
			"ns_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Nameserver Group in Infoblox; will create NS records",
			},
			"restart_if_needed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Restart the member service if necessary for changes to take effect",
			},
			"comment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of this Authoritative Zone; max 256 characters",
			},
			"soa_default_ttl": {
				Type:        schema.TypeInt,
				Default:     3600,
				Optional:    true,
				Description: "Time To Live (TTL) of the SOA record, in seconds",
			},
			"soa_expire": {
				Type:        schema.TypeInt,
				Default:     2419200,
				Optional:    true,
				Description: "Time in seconds for secondary servers to stop answering about the zone because the data is stale (default 1 week)",
			},
			"soa_negative_ttl": {
				Type:        schema.TypeInt,
				Default:     900,
				Optional:    true,
				Description: "Time in seconds for secondary servers to cache data for 'Does Not Respond' responses",
			},
			"soa_refresh": {
				Type:        schema.TypeInt,
				Default:     10800,
				Optional:    true,
				Description: "Interval in seconds for secondary servers to check the primary server for fresh data about the zone",
			},
			"soa_retry": {
				Type:        schema.TypeInt,
				Default:     3600,
				Optional:    true,
				Description: "Interval in seconds for secondary servers to wait before recontacting primary server about the zone after failure",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes, as a map in JSON format",
			},
		},
	}
}

// CRUD FUNCTIONS

func resourceZoneAuthCreate(d *schema.ResourceData, m interface{}) error {
	fqdn := d.Get("fqdn").(string)
	nsGroup := d.Get("ns_group").(string)
	restartIfNeeded := d.Get("restart_if_needed").(bool)
	comment := d.Get("comment").(string)
	soaDefaultTtl := d.Get("soa_default_ttl").(int)
	soaExpire := d.Get("soa_expire").(int)
	soaNegativeTtl := d.Get("soa_negative_ttl").(int)
	soaRefresh := d.Get("soa_refresh").(int)
	soaRetry := d.Get("soa_retry").(int)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newZone, err := objMgr.CreateZoneAuth(
		fqdn, nsGroup, restartIfNeeded, comment, soaDefaultTtl, soaExpire, soaNegativeTtl, soaRefresh, soaRetry, extAttrs)
	if err != nil {
		return fmt.Errorf("error creating Zone Auth: %s", err.Error())
	}

	d.SetId(newZone.Ref)

	return nil
}

func resourceZoneAuthGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetZoneAuthByRef(d.Id())
	if err != nil {
		return fmt.Errorf("getting Zone Auth failed: %s", err.Error())
	}

	d.Set("fqdn", obj.Fqdn)
	d.Set("comment", obj.Comment)
	d.SetId(obj.Ref)

	return nil
}

func resourceZoneAuthUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZoneAuthDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteZoneAuth(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of Zone Auth failed : %s", err.Error())
	}
	d.SetId("")

	return nil
}
