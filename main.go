package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/infobloxopen/terraform-provider-infoblox/infoblox"
	//"fmt"
	//ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: infoblox.Provider})
}

/*
func main(){
	hostConfig := ibclient.HostConfig{
		Host: "10.120.22.235",
		Version: "2.9",
		Port: "443",
		Username: "admin",
		Password: "infoblox",
	}
	transportConfig := ibclient.NewTransportConfig("false",20,10)
	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}
	conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder,requestor)
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Logout()
	objMgr := ibclient.NewObjectManager(conn, "myclient","ntripathy")
	arecord, err := objMgr.GetARecord("default","record1.test.com","10.0.0.1")
	fmt.Println(arecord.Ref)
}*/
