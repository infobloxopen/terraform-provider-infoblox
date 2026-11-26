package ipam

import (
	"errors"
	"strings"
)

// vlansAssignedMessage is the WAPI error NIOS returns when deleting an IPv4 or IPv6 network
// that still has VLANs assigned to it.
const vlansAssignedMessage = "has VLANs assigned to it"

// isVlansAssignedError reports whether err is the NIOS refusal to delete a network with VLANs.
func isVlansAssignedError(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), vlansAssignedMessage) {
		return true
	}
	var withBody interface{ Body() []byte }
	if errors.As(err, &withBody) {
		return strings.Contains(string(withBody.Body()), vlansAssignedMessage)
	}
	return false
}
