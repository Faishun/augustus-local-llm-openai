package openaicompat

import (
	"net/url"
	"strings"
)

// IsLocalBaseURL reports whether baseURL targets a local/loopback host.
// This is used to allow empty API keys for local testing servers.
func IsLocalBaseURL(baseURL string) bool {
	trimmed := strings.TrimSpace(baseURL)
	if trimmed == "" {
		return false
	}

	normalized := trimmed
	if !strings.Contains(normalized, "://") {
		normalized = "http://" + normalized
	}

	parsed, err := url.Parse(normalized)
	if err != nil {
		return false
	}

	host := parsed.Hostname()
	if host == "" {
		return false
	}

	if host == "localhost" || host == "127.0.0.1" || host == "0.0.0.0" || host == "::1" {
		return true
	}

	return strings.HasPrefix(host, "127.")
}
