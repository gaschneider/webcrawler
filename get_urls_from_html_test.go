package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no url",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
			</body>
		</html>
		`,
			expected: nil,
		},
		{
			name:      "even empty body wont generate error",
			inputURL:  "https://blog.boot.dev",
			inputBody: "",
			expected:  nil,
		},
		{
			name:          "invalid base url",
			inputURL:      ":\\invalidBaseURL",
			inputBody:     "",
			expected:      nil,
			errorContains: "couldn't parse base URL",
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href=":\\invalidBaseURL">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v URLs, actual: %v", i, tc.name, len(tc.expected), len(actual))
			}
		})
	}
}
