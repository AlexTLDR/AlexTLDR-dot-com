package services

import (
	"fmt"
	"net/url"
	"testing"
)

func TestTruncateDescription(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxLength int
		expected  string
	}{
		{
			name:      "simple text within limit",
			input:     "This is a simple description",
			maxLength: 100,
			expected:  "This is a simple description",
		},
		{
			name:      "text exceeding limit",
			input:     "This is a very long description that should be truncated",
			maxLength: 20,
			expected:  "This is a very lo...",
		},
		{
			name:      "HTML tags removed",
			input:     "This is <b>bold</b> and <i>italic</i> text",
			maxLength: 100,
			expected:  "This is bold and italic text",
		},
		{
			name:      "HTML entities decoded",
			input:     "This &amp; that &lt;test&gt; &quot;quoted&quot;",
			maxLength: 100,
			expected:  "This & that <test> \"quoted\"",
		},
		{
			name:      "complex HTML with entities",
			input:     "<p>Hello &amp; welcome to <strong>my blog</strong>!</p>",
			maxLength: 100,
			expected:  "Hello & welcome to my blog !",
		},
		{
			name:      "empty string",
			input:     "",
			maxLength: 50,
			expected:  "",
		},
		{
			name:      "only HTML tags",
			input:     "<div><p></p></div>",
			maxLength: 50,
			expected:  "",
		},
		{
			name:      "multiple spaces cleaned up",
			input:     "This    has     multiple   spaces",
			maxLength: 100,
			expected:  "This has multiple spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncateDescription(tt.input, tt.maxLength)
			if result != tt.expected {
				t.Errorf("truncateDescription() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name      string
		urlStr    string
		shouldErr bool
	}{
		{
			name:      "valid HTTPS blog URL",
			urlStr:    "https://blog.alextldr.com/post/example",
			shouldErr: false,
		},
		{
			name:      "valid HTTP blog URL",
			urlStr:    "http://blog.alextldr.com/post/example",
			shouldErr: false,
		},
		{
			name:      "invalid scheme",
			urlStr:    "ftp://blog.alextldr.com/post/example",
			shouldErr: true,
		},
		{
			name:      "file scheme",
			urlStr:    "file:///etc/passwd",
			shouldErr: true,
		},
		{
			name:      "wrong domain",
			urlStr:    "https://malicious.com/post/example",
			shouldErr: true,
		},
		{
			name:      "invalid URL",
			urlStr:    "not-a-url",
			shouldErr: true,
		},
		{
			name:      "empty URL",
			urlStr:    "",
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBlogURL(tt.urlStr)
			if tt.shouldErr && err == nil {
				t.Errorf("validateBlogURL() should have returned error for %q", tt.urlStr)
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("validateBlogURL() should not have returned error for %q, got: %v", tt.urlStr, err)
			}
		})
	}
}

// Helper function to test URL validation logic
func validateBlogURL(urlStr string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Only allow HTTPS/HTTP schemes
	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return fmt.Errorf("invalid scheme: %s", parsedURL.Scheme)
	}

	// Only allow requests to blog.alextldr.com domain
	if parsedURL.Host != "blog.alextldr.com" {
		return fmt.Errorf("invalid domain: %s", parsedURL.Host)
	}

	return nil
}

func TestFetchReadingTimeFromPost_URLValidation(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "invalid URL returns fallback",
			url:      "not-a-url",
			expected: "< 1 min read",
		},
		{
			name:     "wrong domain returns fallback",
			url:      "https://malicious.com/post",
			expected: "< 1 min read",
		},
		{
			name:     "invalid scheme returns fallback",
			url:      "ftp://blog.alextldr.com/post",
			expected: "< 1 min read",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fetchReadingTimeFromPost(tt.url)
			if result != tt.expected {
				t.Errorf("fetchReadingTimeFromPost() = %q, want %q", result, tt.expected)
			}
		})
	}
}
