package infoblox

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceZoneAuthCreate,
		ReadContext:   resourceZoneAuthRead,
		UpdateContext: resourceZoneAuthUpdate,
		DeleteContext: resourceZoneAuthDelete,
		Importer: &schema.ResourceImporter{
			State: resourceZoneAuthImport,
		},

		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of this DNS zone. For a reverse zone, this is in 'address/cidr' " +
					"format. For other zones, this is in FQDN format. This value can be in " +
					"unicode format. Note that for a reverse zone, the corresponding zone_format " +
					"value should be set.",
			},

			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the DNS view in which the zone resides. Example: 'external'",
			},

			"zone_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Determines the format of this zone. Valid values are: FORWARD, IPV4, IPV6; the default is FORWARD",
			},

			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the zone; maximum 256 characters.",
			},

			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Extensible attributes of the zone, as a map in JSON format",
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
				Default:  28800,
				Description: "The Time to Live (TTL) value of the SOA record of this zone. This value is " +
					"the number of seconds that data is cached.",
			},

			"soa_expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2419200,
				Description: "This setting defines the amount of time, in seconds, after which the " +
					"secondary server stops giving out answers about the zone because the zone " +
					"data is too old to be useful. The default is one week.",
			},

			"soa_negative_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
				Description: "The negative Time to Live (TTL) value of the SOA of the zone indicates how " +
					"long a secondary server can cache data for 'Does Not Respond' responses.",
			},

			"soa_refresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10800,
				Description: "This indicates the interval at which a secondary server sends a message to " +
					"the primary server for a zone to check that its data is current, and " +
					"retrieve fresh data if it is not.",
			},

			"soa_retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
				Description: "This indicates how long a secondary server must wait before attempting to " +
					"recontact the primary server after a connection failure between the two " +
					"servers occurs.",
			},
		},
	}
}

func checkZoneFormat(f string) diag.Diagnostics {
	for _, v := range []string{"FORWARD", "IPV4", "IPV6"} {
		if f == v {
			return nil
		}
	}
	return diag.FromErr(fmt.Errorf("zone's format must be either of FORWARD, IPV4, IPV6"))
}

func formZone(
	create bool, d *schema.ResourceData, m interface{}) (
	*ibclient.ZoneAuth, diag.Diagnostics) {

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	zone := &ibclient.ZoneAuth{
		Ea: extAttrs,
	}

	if create {
		zone.Fqdn = d.Get("fqdn").(string)

		zone.View = utils.StringPtr(d.Get("view").(string))
		if *zone.View == "" {
			zone.View = utils.StringPtr("default")
		}

		zf := d.Get("zone_format").(string)
		if errs := checkZoneFormat(zf); errs != nil {
			if zf == "" {
				zf = "FORWARD"
			} else {
				return nil, errs
			}
		}
		zone.ZoneFormat = zf
	}

	if d.HasChange("comment") {
		zone.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("ns_group") {
		nsGrp := d.Get("ns_group").(string)
		if nsGrp != "" {
			zone.NsGroup = utils.StringPtr(nsGrp)
		} else {
			zone.NsGroup = nil
		}
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

	return zone, nil
}

func resourceZoneAuthCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	zone, errs := formZone(true, d, m)
	if errs != nil {
		return errs
	}

	connector := m.(ibclient.IBConnector)
	zoneRef, err := connector.CreateObject(zone)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create a zone: %w", err))
	}

	d.SetId(zoneRef)

	return resourceZoneAuthRead(ctx, d, m)
}

func resourceZoneAuthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return diag.FromErr(err)
	}

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

	err = connector.GetObject(zone, zoneRef, nil, &zoneResult)
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

	omittedEAs := omitEAs(zoneResult.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
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

	if d.HasChange("zone_format") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: zone_format"))
	}

	if d.HasChange("view") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: view"))
	}

	zone, errs := formZone(false, d, m)
	if errs != nil {
		return errs
	}

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "terraform_test_tenant")
	zoneVal, err := objMgr.GetZoneAuthByRef(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read Zone Auth for update operation: %w", err))
	}

	zone.Ea, err = mergeEAs(zoneVal.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return diag.FromErr(err)
	}

	zoneRef, err := connector.UpdateObject(zone, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update a zone: %w", err))
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

func resourceZoneAuthImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)

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

	err = connector.GetObject(zone, zoneRef, nil, &zoneResult)
	if err != nil {
		return nil, fmt.Errorf("failed to read zone: %w", err)
	}

	err = d.Set("fqdn", zoneResult.Fqdn)
	if err != nil {
		return nil, err
	}

	if zoneResult.Comment != nil {
		err = d.Set("comment", *zoneResult.Comment)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.NsGroup != nil {
		err = d.Set("ns_group", *zoneResult.NsGroup)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaDefaultTtl != nil {
		err = d.Set("soa_default_ttl", *zoneResult.SoaDefaultTtl)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaExpire != nil {
		err = d.Set("soa_expire", *zoneResult.SoaExpire)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaNegativeTtl != nil {
		err = d.Set("soa_negative_ttl", *zoneResult.SoaNegativeTtl)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaRefresh != nil {
		err = d.Set("soa_refresh", *zoneResult.SoaRefresh)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaRetry != nil {
		err = d.Set("soa_retry", *zoneResult.SoaRetry)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.View != nil {
		err = d.Set("view", *zoneResult.View)
		if err != nil {
			return nil, err
		}
	}

	err = d.Set("zone_format", zoneResult.ZoneFormat)
	if err != nil {
		return nil, err
	}

	if zoneResult.Ea != nil && len(zoneResult.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(zoneResult.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	d.SetId(zoneResult.Ref)

	return []*schema.ResourceData{d}, nil
}
