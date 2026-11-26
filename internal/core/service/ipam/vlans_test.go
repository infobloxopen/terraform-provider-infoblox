package ipam

import (
	"errors"
	"fmt"
	"testing"
)

type bodyErr struct{ body string }

func (e bodyErr) Error() string { return "400 Bad Request" }
func (e bodyErr) Body() []byte  { return []byte(e.body) }

func TestIsVlansAssignedError(t *testing.T) {
	msg := `{ "Error": "Cannot delete the '10.0.0.0/24' network because it has VLANs assigned to it." }`
	cases := map[string]struct {
		err  error
		want bool
	}{
		"nil":             {nil, false},
		"in message":      {errors.New("400 Bad Request, '" + msg + "'"), true},
		"in wrapped body": {fmt.Errorf("delete: %w", bodyErr{msg}), true},
		"other error":     {bodyErr{`{ "Error": "Network not found" }`}, false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := isVlansAssignedError(tc.err); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
