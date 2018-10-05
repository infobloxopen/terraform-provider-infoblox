package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"strings"
	"time"
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
				DefaultFunc: schema.EnvDefaultFunc("WAPI_VERSION", "2.8"),
				Description: "WAPI Version of Infoblox server defaults to v2.8.",
				ValidateFunc: StringInSlice([]string{"2.1", "2.1.1",
					"2.1.2", "2.2", "2.2.1", "2.2.2", "2.3", "2.3.1", "2.4", "2.5", "2.6", "2.6.1",
					"2.7", "2.7.1", "2.8", "2.9"}, false),
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
			"infoblox_network":        resourceNetwork(),
			"infoblox_network_view":   resourceNetworkView(),
			"infoblox_ip_allocation":  resourceIPAllocation(),
			"infoblox_ip_association": resourceIPAssociation(),
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
	objMgr := ibclient.NewObjectManager(conn, "infoblox", "terraform-provider-infoblox")
	err = CheckCloudLicense(objMgr, "cloud")
	if err != nil {
		return nil, err
	}
	return conn, err
}

//StringInSlice returns a ValidateFunc which tests if the provided version
// is of type string and matches the value of element in the slice
func StringInSlice(valid []string, ignoreCase bool) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(string)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be string", k))
			return
		}

		for _, str := range valid {
			if v == str || (ignoreCase && strings.ToLower(v) == strings.ToLower(str)) {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %s", k, valid, v))
		return
	}
}

//CheckCloudLicense checks whether the user has applied License to
//Infoblox Server.
func CheckCloudLicense(objMgr *ibclient.ObjectManager, licenseType string) (err error) {
	license, err := objMgr.GetLicense()

	if err != nil {
		return

	}
	for _, v := range license {
		if strings.ToLower(v.Licensetype) == licenseType {
			if v.ExpirationStatus != "DELETED" && v.ExpirationStatus != "EXPIRED" {
				return
			}

		}
	}
	err = fmt.Errorf("%s license is not applied/deleted for the grid. Apply the license and try again", licenseType)
	return
}
