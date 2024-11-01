package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name        string
		headers     map[string]string
		expectedKey string
		expectError bool
	}{
		{
			name:        "with valid authorizatio header",
			headers:     map[string]string{"Authorization": "ApiKey validAPIKey"},
			expectedKey: "validAPIKey",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headers:     map[string]string{},
			expectedKey: "",
			expectError: true,
		},
		// invalid authorization format
		// Authorization header with empty string
		// malformed authorization header
		// multiple authorization headers
		// non-string characters in authorization
		// very large authorization value
		// authorization header wth lowercase key
		{
			name:        "Authorization header with lowercase key",
			headers:     map[string]string{"authorization": "ApiKey validAPIKey"},
			expectedKey: "validAPIKey",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpHeaders := make(http.Header)
			for k, v := range tc.headers {
				httpHeaders.Set(k, v)
			}

			apiKey, err := GetAPIKey(httpHeaders)
			if (err != nil) != tc.expectError {
				t.Errorf("expected error: %v, got %v", tc.expectError, err)
			}
			if apiKey != tc.expectedKey {
				t.Errorf("expected key: %s, got %s", tc.expectedKey, apiKey)
			}
		})
	}
}
