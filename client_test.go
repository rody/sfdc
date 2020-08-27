package sfdc

import (
	"net/http"
	"testing"
)

func testNewClientWithValidURL(t *testing.T) {
	if _, err := NewClient(http.DefaultClient, "https://myorg.salesforce.com"); err != nil {
		t.Fatalf("expected a new client, got this error: %s", err)
	}
}
