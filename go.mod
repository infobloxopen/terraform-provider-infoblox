module github.com/infobloxopen/terraform-provider-infoblox

go 1.15

replace github.com/infobloxopen/infoblox-go-client => ../infoblox-go-client

require (
	github.com/hashicorp/terraform v0.12.9
	github.com/infobloxopen/infoblox-go-client v1.1.0
)
