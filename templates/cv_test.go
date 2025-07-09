package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestCV(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = CV().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check main heading exists and is not empty
	heading := doc.Find("h1").Text()
	if heading == "" {
		t.Error("main heading is empty")
	}

	// Check tech stack section exists
	techStackHeading := doc.Find("h2").Text()
	if techStackHeading == "" {
		t.Error("tech stack heading is empty")
	}

	// Check tech stack links exist (general count)
	techStackLinks := doc.Find("a[href*='golang.org'], a[href*='python.org'], a[href*='javascript.com']")
	if techStackLinks.Length() < 3 {
		t.Errorf("expected at least 3 tech stack links, got %d", techStackLinks.Length())
	}

	// Check certification links exist (general count)
	certificationLinks := doc.Find("a[href*='credly.com'], a[href*='boot.dev'], a[href*='neo4j.com']")
	if certificationLinks.Length() < 3 {
		t.Errorf("expected at least 3 certification links, got %d", certificationLinks.Length())
	}

	// Check specific LinkedIn certifications link that needs to be precise
	linkedinCertLink := "https://www.linkedin.com/in/alexandru-badragan/details/certifications/"
	if doc.Find(`a[href="`+linkedinCertLink+`"]`).Length() == 0 {
		t.Errorf("LinkedIn certifications link %s not found", linkedinCertLink)
	}

	// Check work experience section has company info
	companyLogos := doc.Find("img[alt*='Logo']")
	if companyLogos.Length() == 0 {
		t.Error("no company logos found")
	}

	// Check some paragraphs with content exist
	paragraphs := doc.Find("p")
	if paragraphs.Length() < 2 {
		t.Errorf("expected at least 2 paragraphs, got %d", paragraphs.Length())
	}
}
