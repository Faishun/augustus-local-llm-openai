package openaicompat

import "testing"

func TestIsLocalBaseURL(t *testing.T) {
	cases := []struct {
		baseURL string
		expect  bool
	}{
		{"http://localhost:8000/v1", true},
		{"localhost:8000", true},
		{"http://127.0.0.1:8000", true},
		{"http://127.1.2.3:8000", true},
		{"http://0.0.0.0:8000", true},
		{"http://[::1]:8000", true},
		{"https://api.openai.com/v1", false},
		{"", false},
	}

	for _, tc := range cases {
		if got := IsLocalBaseURL(tc.baseURL); got != tc.expect {
			t.Errorf("IsLocalBaseURL(%q) = %v, want %v", tc.baseURL, got, tc.expect)
		}
	}
}
