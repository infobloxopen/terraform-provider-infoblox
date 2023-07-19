package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceZoneAuthCreate,
		ReadContext:   resourceZoneAuthRead,
		UpdateContext: resourceZoneAuthUpdate,
		DeleteContext: resourceZoneAuthDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the zone; maximum 256 characters.",
			},

			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of this DNS zone. For a reverse zone, this is in \"address/cidr\"\n" +
					"format. For other zones, this is in FQDN format. This value can be in\n" +
					"unicode format. Note that for a reverse zone, the corresponding zone_format\n" +
					" value should be set.",
			},

			"ns_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name server group that serves DNS for this zone.",
			},

			"restart_if_needed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restarts the member service.",
			},

			"soa_default_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The Time to Live (TTL) value of the SOA record of this zone. This value is\n" +
					"the number of seconds that data is cached.",
			},

			"soa_expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "This setting defines the amount of time, in seconds, after which the\n" +
					"secondary server stops giving out answers about the zone because the zone\n" +
					"data is too old to be useful. The default is one week.",
			},

			"soa_negative_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The negative Time to Live (TTL) value of the SOA of the zone indicates how\n" +
					"long a secondary server can cache data for \"Does Not Respond\" responses.",
			},

			"soa_refresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "This indicates the interval at which a secondary server sends a message to\n" +
					"the primary server for a zone to check that its data is current, and\n" +
					"retrieve fresh data if it is not.",
			},

			"soa_retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "This indicates how long a secondary server must wait before attempting to\n" +
					"recontact the primary server after a connection failure between the two\n" +
					"servers occurs.",
			},

			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the DNS view in which the zone resides. Example \"external\".",
			},

			"zone_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Determines the format of this zone.",
			},

			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Extensible attributes of the zone, as a map in JSON format",
			},
		},
	}
}

func resourceZoneAuthCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return diag.FromErr(fmt.Errorf("cannot process 'ext_attrs' field: %w", err))
		}
	}
	connector := m.(ibclient.IBConnector)

	zone := &ibclient.ZoneAuth{
		Fqdn:       d.Get("fqdn").(string),
		ZoneFormat: d.Get("zone_format").(string),
		Ea:         extAttrs,
	}

	if d.HasChange("comment") {
		zone.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("ns_group") {
		zone.NsGroup = utils.StringPtr(d.Get("ns_group").(string))
	}

	if d.HasChange("restart_if_needed") {
		zone.RestartIfNeeded = utils.BoolPtr(d.Get("restart_if_needed").(bool))
	}

	if d.HasChange("soa_default_ttl") {
		zone.SoaDefaultTtl = utils.Uint32Ptr(uint32(d.Get("soa_default_ttl").(int)))
	}

	if d.HasChange("soa_expire") {
		zone.SoaExpire = utils.Uint32Ptr(uint32(d.Get("soa_expire").(int)))
	}

	if d.HasChange("soa_negative_ttl") {
		zone.SoaNegativeTtl = utils.Uint32Ptr(uint32(d.Get("soa_negative_ttl").(int)))
	}

	if d.HasChange("soa_refresh") {
		zone.SoaRefresh = utils.Uint32Ptr(uint32(d.Get("soa_refresh").(int)))
	}

	if d.HasChange("soa_retry") {
		zone.SoaRetry = utils.Uint32Ptr(uint32(d.Get("soa_retry").(int)))
	}

	if d.HasChange("view") {
		zone.View = utils.StringPtr(d.Get("view").(string))
	}

	zoneRef, err := connector.CreateObject(zone)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create a zone: %w", err))
	}

	d.SetId(zoneRef)

	return resourceZoneAuthRead(ctx, d, m)
}

func resourceZoneAuthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)
	var diags diag.Diagnostics

	zoneRef := d.Id()

	zoneResult := ibclient.ZoneAuth{}

	zone := &ibclient.ZoneAuth{}
	zone.SetReturnFields(append(
		zoneResult.ReturnFields(),
		"comment",
		"ns_group",
		"soa_default_ttl",
		"soa_expire",
		"soa_negative_ttl",
		"soa_refresh",
		"soa_retry",
		"view",
		"zone_format",
		"extattrs",
	))

	err := connector.GetObject(zone, zoneRef, nil, &zoneResult)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read zone: %w", err))
	}

	err = d.Set("fqdn", zoneResult.Fqdn)
	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResult.Comment != nil {
		err = d.Set("comment", *zoneResult.Comment)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.NsGroup != nil {
		err = d.Set("ns_group", *zoneResult.NsGroup)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaDefaultTtl != nil {
		err = d.Set("soa_default_ttl", *zoneResult.SoaDefaultTtl)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaExpire != nil {
		err = d.Set("soa_expire", *zoneResult.SoaExpire)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaNegativeTtl != nil {
		err = d.Set("soa_negative_ttl", *zoneResult.SoaNegativeTtl)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaRefresh != nil {
		err = d.Set("soa_refresh", *zoneResult.SoaRefresh)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaRetry != nil {
		err = d.Set("soa_retry", *zoneResult.SoaRetry)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.View != nil {
		err = d.Set("view", *zoneResult.View)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err = d.Set("zone_format", zoneResult.ZoneFormat)
	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResult.Ea != nil && len(zoneResult.Ea) > 0 {
		// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
		//       (avoiding additional layer of keys ("value" key)
		eaMap := (map[string]interface{})(zoneResult.Ea)
		ea, err := json.Marshal(eaMap)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("ext_attrs", string(ea)); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(zoneResult.Ref)

	return diags
}
func resourceZoneAuthUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	if d.HasChange("fqdn") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: fqdn"))
	}

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return diag.FromErr(fmt.Errorf("cannot process 'ext_attrs' field: %w", err))
		}
	}

	connector := m.(ibclient.IBConnector)

	zone := &ibclient.ZoneAuth{
		Ea: extAttrs,
	}

	if d.HasChange("comment") {
		zone.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	zoneRef, err := connector.UpdateObject(zone, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create a zone: %w", err))
	}

	d.SetId(zoneRef)

	return resourceZoneAuthRead(ctx, d, m)
}

func resourceZoneAuthDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	zoneRef := d.Id()

	_, err := connector.DeleteObject(zoneRef)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
