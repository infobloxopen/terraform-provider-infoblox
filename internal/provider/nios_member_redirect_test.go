package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/types"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosoption "github.com/infobloxopen/infoblox-nios-go-client/option"
)

const (
	memberPage = `<HTML><HEAD><META HTTP-EQUIV="REFRESH" CONTENT = "0; URL=https://10.0.0.1"/></HEAD><BODY>Click <A HREF="https://10.0.0.1">here</A></BODY></HTML>`
	gm401Page  = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN"><html><head><title>401 Unauthorized</title></head></html>`
)

// fakeWAPI returns a NIOS client whose every call gets the given reply.
func fakeWAPI(t *testing.T, status int, contentType, body string) *niosclient.APIClient {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return niosclient.NewAPIClient(
		niosoption.WithNIOSHostUrl(srv.URL),
		niosoption.WithNIOSUsername("u"),
		niosoption.WithNIOSPassword("p"),
	)
}

func TestNiosMemberRedirect(t *testing.T) {
	cases := map[string]struct {
		status      int
		contentType string
		body        string
		hostURL     string
		wantURL     string
		wantMatch   bool
	}{
		"grid member redirect page": {
			status: http.StatusOK, contentType: "text/html", body: memberPage,
			hostURL: "https://10.0.0.5", wantURL: "https://10.0.0.1", wantMatch: true,
		},
		"redirect to the configured host itself": {
			status: http.StatusOK, contentType: "text/html", body: memberPage,
			hostURL: "https://10.0.0.1",
		},
		"grid master 401 page": {
			status: http.StatusUnauthorized, contentType: "text/html; charset=iso-8859-1", body: gm401Page,
			hostURL: "https://10.0.0.5",
		},
		"wapi json error": {
			status: http.StatusBadRequest, contentType: "application/json", body: `{"Error": "AdmConProtoError: bad request"}`,
			hostURL: "https://10.0.0.5",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client := fakeWAPI(t, tc.status, tc.contentType, tc.body)
			err := checkAndCreatePreRequisitesForNIOS(context.Background(), client)
			if err == nil {
				t.Fatal("expected an error")
			}

			gotURL, ok := niosMemberRedirect(err, tc.hostURL)
			if ok != tc.wantMatch || gotURL != tc.wantURL {
				t.Fatalf("niosMemberRedirect() = (%q, %v), want (%q, %v); err: %v", gotURL, ok, tc.wantURL, tc.wantMatch, err)
			}
		})
	}
}

func TestEnsureNIOSPreRequisitesMemberRedirect(t *testing.T) {
	hostURLPath := path.Root("nios").AtName("host_url")
	optOut := types.BoolValue(false)

	cases := map[string]struct {
		status      int
		contentType string
		body        string
		hostURL     string
		manage      types.Bool
		wantOK      bool
		wantHostErr bool // exactly one error, on nios.host_url
	}{
		"member, default": {
			status: http.StatusOK, contentType: "text/html", body: memberPage,
			hostURL: "https://10.0.0.5", manage: types.BoolNull(), wantHostErr: true,
		},
		"member, manage_internal_id_ea = false": {
			status: http.StatusOK, contentType: "text/html", body: memberPage,
			hostURL: "https://10.0.0.5", manage: optOut, wantHostErr: true,
		},
		"member page, portal passthrough": {
			status: http.StatusOK, contentType: "text/html", body: memberPage,
			hostURL: "", manage: types.BoolNull(),
		},
		"grid master 401, manage_internal_id_ea = false": {
			status: http.StatusUnauthorized, contentType: "text/html; charset=iso-8859-1", body: gm401Page,
			hostURL: "https://10.0.0.5", manage: optOut, wantOK: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client := fakeWAPI(t, tc.status, tc.contentType, tc.body)
			resp := &provider.ConfigureResponse{}

			ok := ensureNIOSPreRequisites(context.Background(), client, tc.hostURL, tc.manage, resp)
			if ok != tc.wantOK {
				t.Fatalf("ensureNIOSPreRequisites() = %v, want %v; diagnostics: %v", ok, tc.wantOK, resp.Diagnostics)
			}

			errs := resp.Diagnostics.Errors()
			hostErr := len(errs) == 1 && errs[0].Summary() == "host_url Is Not the Grid Master"
			if hostErr {
				withPath, isPath := errs[0].(diag.DiagnosticWithPath)
				hostErr = isPath && withPath.Path().Equal(hostURLPath)
			}
			if hostErr != tc.wantHostErr {
				t.Fatalf("got host_url error = %v, want %v; errors: %v", hostErr, tc.wantHostErr, errs)
			}
			if tc.wantOK && len(errs) != 0 {
				t.Fatalf("expected no errors, got %v", errs)
			}
		})
	}
}
