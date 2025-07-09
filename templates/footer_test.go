package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFooter(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = Footer().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check footer exists
	if doc.Find("footer").Length() == 0 {
		t.Error("footer element not found")
	}

	// Check footer has some text
	footerText := doc.Find("footer").Text()
	if footerText == "" {
		t.Error("footer text is empty")
	}
}

func TestFooterMobile(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = FooterMobile().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check mobile footer has content
	footerText := doc.Find("p").Text()
	if footerText == "" {
		t.Error("mobile footer text is empty")
	}
}
