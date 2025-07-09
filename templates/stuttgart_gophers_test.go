package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestStuttgartGophers(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = StuttgartGophers().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check heading exists and is not empty
	heading := doc.Find("h2").Text()
	if heading == "" {
		t.Error("heading is empty")
	}

	// Check specific Stuttgart Gophers links that need to be precise
	stuttgartLinks := []string{
		"https://www.meetup.com/stuttgart-gophers",
		"https://stuttgart-gophers.de",
	}
	for _, link := range stuttgartLinks {
		if doc.Find(`a[href="`+link+`"]`).Length() == 0 {
			t.Errorf("Stuttgart Gophers link %s not found", link)
		}
	}

	// Check some content paragraphs exist
	paragraphs := doc.Find("p")
	if paragraphs.Length() < 2 {
		t.Errorf("expected at least 2 paragraphs, got %d", paragraphs.Length())
	}

	// Check dancing gopher images exist
	gopherImages := doc.Find(`img[alt="Dancing Gopher"]`)
	if gopherImages.Length() == 0 {
		t.Error("dancing gopher images not found")
	}
}
