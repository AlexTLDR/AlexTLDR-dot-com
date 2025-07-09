package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestIndex(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = Index("test").Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check basic HTML structure
	if doc.Find("html").Length() == 0 {
		t.Error("html element not found")
	}

	if doc.Find("head").Length() == 0 {
		t.Error("head element not found")
	}

	if doc.Find("body").Length() == 0 {
		t.Error("body element not found")
	}

	// Check title exists and is not empty
	title := doc.Find("title").Text()
	if title == "" {
		t.Error("title is empty")
	}

	// Check essential meta tags
	charset := doc.Find(`meta[charset="UTF-8"]`)
	if charset.Length() == 0 {
		t.Error("charset meta tag not found")
	}

	viewport := doc.Find(`meta[name="viewport"]`)
	if viewport.Length() == 0 {
		t.Error("viewport meta tag not found")
	}

	// Check favicon exists
	favicon := doc.Find(`link[rel="icon"]`)
	if favicon.Length() == 0 {
		t.Error("favicon link not found")
	}

	// Check HTMX script is included
	htmxScript := doc.Find(`script[src*="htmx.org"]`)
	if htmxScript.Length() == 0 {
		t.Error("HTMX script not found")
	}

	// Check Alpine.js script is included
	alpineScript := doc.Find(`script[src*="alpinejs"]`)
	if alpineScript.Length() == 0 {
		t.Error("Alpine.js script not found")
	}

	// Check main sections are rendered (they should be present as the template includes them)
	// We don't need to test their content here since they have their own tests
	sections := []string{"#about", "#portfolio", "#stuttgart-gophers", "#blog"}
	for _, section := range sections {
		if doc.Find(section).Length() == 0 {
			t.Errorf("section %s not found", section)
		}
	}
}
