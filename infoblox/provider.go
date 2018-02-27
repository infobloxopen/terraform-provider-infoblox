package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOST", nil),
				Description: "NIOS IP address. Server IP address",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("USERNAME", nil),
				Description: "User to authenticate ",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", nil),
				Description: "Password to authenticate ",
			},
			"wapi_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("WAPI_VERSION", "2.8"),
				Description: "which wapi version to be used",
			},
			"port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PORT", "443"),
				Description: "port number used for connection",
			},

			"sslmode": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SSLMODE", "false"),
				Description: "If set, Infoblox client will permit unverifiable SSL certificates",
			},
			"connect_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONNECT_TIMEOUT", "60"),
				Description: "Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.",
			},
			"pool_connections": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("POOL_CONNECTIONS", "10"),
				Description: "Maximum number of connections to establish to the database. Zero means unlimited.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"infoblox_network":       resourceNetwork(),
			"infoblox_network_view":  resourceNetworkView(),
			"infoblox_ip_allocation": resourceAllocation(),
		},
		ConfigureFunc: providerConfigure,
	}

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	hostConfig := ibclient.HostConfig{
		Host:     d.Get("server").(string),
		Port:     d.Get("port").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Version:  d.Get("wapi_version").(string),
	}

	transportConfig := ibclient.TransportConfig{
		SslVerify:           d.Get("sslmode").(bool),
		HttpRequestTimeout:  d.Get("connect_timeout").(int),
		HttpPoolConnections: d.Get("pool_connections").(int),
	}

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)

	return conn, err
}
