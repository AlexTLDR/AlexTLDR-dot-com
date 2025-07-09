package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestBlog(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = Blog().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check heading exists
	heading := doc.Find("h2").Text()
	if heading == "" {
		t.Error("heading is empty")
	}

	// Check specific blog link exists
	if doc.Find(`a[href="https://blog.alextldr.com"]`).Length() == 0 {
		t.Error("blog link https://blog.alextldr.com not found")
	}

	// Check HTMX container exists
	htmxContainer := doc.Find("[hx-get]")
	if htmxContainer.Length() == 0 {
		t.Error("HTMX container not found")
	}

	// Check loading indicator exists
	if doc.Find("#blog-loading").Length() == 0 {
		t.Error("loading indicator not found")
	}
}

func TestBlogPosts_WithPosts(t *testing.T) {
	// Create test blog posts
	posts := []BlogPost{
		{
			Title:       "Test Post 1",
			Description: "This is a test description",
			URL:         "https://blog.alextldr.com/post1",
			Date:        "Jan 1, 2024",
			ReadTime:    "5 min read",
		},
		{
			Title:       "Test Post 2",
			Description: "Another test description",
			URL:         "https://blog.alextldr.com/post2",
			Date:        "Jan 2, 2024",
			ReadTime:    "3 min read",
		},
	}

	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = BlogPosts(posts).Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check correct number of cards
	cards := doc.Find(".card")
	if cards.Length() != len(posts) {
		t.Errorf("expected %d cards, got %d", len(posts), cards.Length())
	}

	// Check each post has title, description, and link
	for i := range posts {
		card := cards.Eq(i)

		// Check title exists
		title := card.Find(".card-title").Text()
		if title == "" {
			t.Errorf("post %d title is empty", i)
		}

		// Check description exists
		description := card.Find("p").Text()
		if description == "" {
			t.Errorf("post %d description is empty", i)
		}

		// Check link exists
		link := card.Find("a[href*='blog.alextldr.com']")
		if link.Length() == 0 {
			t.Errorf("post %d link not found", i)
		}
	}
}

func TestBlogPosts_NoPosts(t *testing.T) {
	// Test with empty posts slice
	posts := []BlogPost{}

	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = BlogPosts(posts).Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check no posts message exists
	alert := doc.Find(".alert")
	if alert.Length() == 0 {
		t.Error("no posts alert not found")
	}

	// Check alert has text
	alertText := alert.Text()
	if alertText == "" {
		t.Error("alert text is empty")
	}

	// Check no cards rendered
	cards := doc.Find(".card")
	if cards.Length() != 0 {
		t.Errorf("expected no cards, got %d", cards.Length())
	}
}
