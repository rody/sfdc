package sfdc

import "golang.org/x/oauth2"

var (
	// Endpoint is the OAuth2 endpoint for production orgs
	Endpoint = oauth2.Endpoint{
		AuthURL:  "https://login.salesforce.com/services/oauth2/authorize",
		TokenURL: "https://login.salesforce.com/services/oauth2/token",
	}
	// TestEndpoint is the OAuth2 endpoint for sandboxes
	TestEndpoint = oauth2.Endpoint{
		AuthURL:  "https://test.salesforce.com/services/oauth2/authorize",
		TokenURL: "https://test.salesforce.com/services/oauth2/token",
	}
)
