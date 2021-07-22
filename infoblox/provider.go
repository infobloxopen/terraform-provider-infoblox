package infoblox

import (
	"math"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// Common parameters
const (
	ttlUndef = math.MinInt32
)

//Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_SERVER", nil),
				Description: "Infoblox server IP address.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USERNAME", nil),
				Description: "User to authenticate with Infoblox server.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil),
				Description: "Password to authenticate with Infoblox server.",
			},
			"wapi_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("WAPI_VERSION", "2.7"),
				Description: "WAPI Version of Infoblox server defaults to v2.7.",
			},
			"port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PORT", "443"),
				Description: "Port number used for connection for Infoblox Server.",
			},

			"sslmode": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SSLMODE", "false"),
				Description: "If set, Infoblox client will permit unverifiable SSL certificates.",
			},
			"connect_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONNECT_TIMEOUT", 60),
				Description: "Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.",
			},
			"pool_connections": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("POOL_CONNECTIONS", "10"),
				Description: "Maximum number of connections to establish to the Infoblox server. Zero means unlimited.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"infoblox_network_view":           resourceNetworkView(),
			"infoblox_ipv4_network_container": resourceIPv4NetworkContainer(),
			"infoblox_ipv6_network_container": resourceIPv6NetworkContainer(),
			"infoblox_ipv4_network":           resourceIPv4Network(),
			"infoblox_ipv6_network":           resourceIPv6Network(),
			"infoblox_ipv4_allocation":        resourceIPv4Allocation(),
			"infoblox_ipv6_allocation":        resourceIPv6Allocation(),
			"infoblox_ipv4_association":       resourceIPv4Association(),
			"infoblox_ipv6_association":       resourceIPv6Association(),
			"infoblox_a_record":               resourceARecord(),
			"infoblox_aaaa_record":            resourceAAAARecord(),
			"infoblox_cname_record":           resourceCNAMERecord(),
			"infoblox_ptr_record":             resourcePTRRecord(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"infoblox_ipv4_network": dataSourceIPv4Network(),
			"infoblox_a_record":     dataSourceARecord(),
			"infoblox_cname_record": dataSourceCNameRecord(),
		},
		ConfigureFunc: providerConfigure,
	}

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	var seconds int64
	seconds = int64(d.Get("connect_timeout").(int))
	hostConfig := ibclient.HostConfig{
		Host:     d.Get("server").(string),
		Port:     d.Get("port").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Version:  d.Get("wapi_version").(string),
	}

	transportConfig := ibclient.TransportConfig{
		SslVerify:           d.Get("sslmode").(bool),
		HttpRequestTimeout:  time.Duration(seconds),
		HttpPoolConnections: d.Get("pool_connections").(int),
	}

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)
	if err != nil {
		return nil, err
	}
	return conn, err
}
