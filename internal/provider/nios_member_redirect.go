package provider

import (
	"errors"
	"regexp"
)

// metaRefreshURL matches the redirect page a NIOS Grid member returns for WAPI calls:
// <META HTTP-EQUIV="REFRESH" CONTENT = "0; URL=https://<grid master>"/>
var metaRefreshURL = regexp.MustCompile(`(?i)http-equiv\s*=\s*"?refresh"?[^>]*url\s*=\s*([^"'\s>]+)`)

// niosMemberRedirect reports whether err comes from a Grid member's redirect page,
// returning the Grid Master URL it points to.
func niosMemberRedirect(err error) (string, bool) {
	var withBody interface{ Body() []byte }
	if !errors.As(err, &withBody) {
		return "", false
	}
	m := metaRefreshURL.FindSubmatch(withBody.Body())
	if m == nil {
		return "", false
	}
	return string(m[1]), true
}
