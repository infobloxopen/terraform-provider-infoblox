package provider

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// metaRefreshURL matches the redirect page a NIOS Grid member returns for WAPI calls:
// <META HTTP-EQUIV="REFRESH" CONTENT = "0; URL=https://<grid master>"/>
var metaRefreshURL = regexp.MustCompile(`(?i)http-equiv\s*=\s*"?refresh"?[^>]*url\s*=\s*([^"'\s>]+)`)

// niosMemberRedirect reports whether err comes from a Grid member's redirect page,
// returning the Grid Master URL it points to. hostURL is the configured nios.host_url;
// a page that redirects to that same host is not treated as a member redirect.
func niosMemberRedirect(err error, hostURL string) (string, bool) {
	var withBody interface{ Body() []byte }
	if !errors.As(err, &withBody) {
		return "", false
	}
	m := metaRefreshURL.FindSubmatch(withBody.Body())
	if m == nil {
		return "", false
	}
	target := string(m[1])
	if sameHost(target, hostURL) {
		return "", false
	}
	return target, true
}

// probeNIOSMemberRedirect makes one WAPI call to detect a Grid member, for when the
// Terraform Internal ID check (which would otherwise surface it) is disabled.
// Any other outcome, including permission errors, is ignored.
func probeNIOSMemberRedirect(ctx context.Context, client *niosclient.APIClient, hostURL string) (string, bool) {
	_, _, err := client.GridAPI.ExtensibleattributedefAPI.
		List(ctx).
		Filters(map[string]any{"name": flex.TerraformInternalID}).
		ReturnAsObject(1).
		Execute()
	if err == nil {
		return "", false
	}
	return niosMemberRedirect(err, hostURL)
}

func addNIOSMemberRedirectError(resp *provider.ConfigureResponse, gridMaster string) {
	resp.Diagnostics.AddAttributeError(
		path.Root("nios").AtName("host_url"),
		"host_url Is Not the Grid Master",
		fmt.Sprintf("The NIOS host_url points to a Grid member, which redirects WAPI requests to the Grid Master at %s. "+
			"Set host_url to the Grid Master address (for an HA Grid Master, use its VIP).", gridMaster),
	)
}

func sameHost(a, b string) bool {
	ua, errA := url.Parse(a)
	ub, errB := url.Parse(b)
	if errA != nil || errB != nil || ua.Hostname() == "" {
		return false
	}
	return strings.EqualFold(ua.Hostname(), ub.Hostname())
}
