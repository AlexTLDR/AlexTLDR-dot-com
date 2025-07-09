package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestAbout(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = About().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check main heading exists and is not empty
	heading := doc.Find("h1").Text()
	if heading == "" {
		t.Error("heading is empty")
	}

	// Check profile image exists
	img := doc.Find(`img[alt="Alex's Avatar"]`)
	if img.Length() == 0 {
		t.Error("profile image not found")
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

	// Check Stuttgart Gophers link exists
	if doc.Find(`a[href="https://stuttgart-gophers.de/"]`).Length() == 0 {
		t.Error("Stuttgart Gophers link not found")
	}

	// Check some content paragraphs exist
	paragraphs := doc.Find("p")
	if paragraphs.Length() < 3 {
		t.Errorf("expected at least 3 paragraphs, got %d", paragraphs.Length())
	}

	// Check divider exists and has text
	divider := doc.Find(".divider").Text()
	if divider == "" {
		t.Error("divider text is empty")
	}
}
