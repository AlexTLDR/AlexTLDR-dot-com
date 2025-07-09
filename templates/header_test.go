package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHeader(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = Header().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check navigation links exist (general check)
	navLinks := doc.Find("a[href^='#']")
	if navLinks.Length() < 3 {
		t.Errorf("expected at least 3 navigation links, got %d", navLinks.Length())
	}

	// Check specific social links exist
	socialLinks := []string{
		"https://linkedin.com/in/alextldr",
		"https://github.com/AlexTLDR",
		"mailto:alex@alextldr.com",
		"https://blog.alextldr.com",
	}
	for _, link := range socialLinks {
		if doc.Find(`a[href="`+link+`"]`).Length() == 0 {
			t.Errorf("social link %s not found", link)
		}
	}

	// Check theme toggle exists
	if doc.Find("input[type='checkbox']").Length() == 0 {
		t.Error("theme toggle not found")
	}

	// Check mobile menu button exists
	if doc.Find(`[role="button"]`).Length() == 0 {
		t.Error("mobile menu button not found")
	}
}
