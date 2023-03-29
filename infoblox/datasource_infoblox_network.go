package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceIPv4Network() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPv4NetworkRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_view": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  defaultNetView,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A string describing the network",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes for network datasource, as a map in JSON format",
			},
		},
	}
}

func dataSourceIPv4NetworkRead(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")
	obj, err := objMgr.GetNetwork(networkView, cidr, false, nil)
	if err != nil {
		return fmt.Errorf("getting network '%s' failed: %s", cidr, err)
	}
	if obj == nil {
		return fmt.Errorf("API returns a nil/empty ID for the network '%s'", cidr)
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	var eaMap map[string]interface{}
	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaMap = (map[string]interface{})(obj.Ea)
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return err
	}
	if err = d.Set("ext_attrs", string(ea)); err != nil {
		return err
	}

	if err := d.Set("comment", obj.Comment); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
