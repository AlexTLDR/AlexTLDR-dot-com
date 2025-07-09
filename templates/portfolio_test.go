package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestPortfolio(t *testing.T) {
	// Render the template
	r, w := io.Pipe()
	go func() {
		_ = Portfolio().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Check heading exists and is not empty
	heading := doc.Find("h1").Text()
	if heading == "" {
		t.Error("heading is empty")
	}

	// Check project cards exist
	cards := doc.Find(".card")
	if cards.Length() < 5 {
		t.Errorf("expected at least 5 project cards, got %d", cards.Length())
	}

	// Check specific project links that need to be precise
	projectLinks := []string{
		"https://pizzeria.alextldr.com/",
		"https://interpretor.alextldr.com/",
		"https://blog.alextldr.com/",
		"https://stuttgart-gophers.de/",
		"https://fotohive.alextldr.com/",
	}
	for _, link := range projectLinks {
		if doc.Find(`a[href="`+link+`"]`).Length() == 0 {
			t.Errorf("project link %s not found", link)
		}
	}

	// Check specific GitHub links that need to be precise
	githubLinks := []string{
		"https://github.com/AlexTLDR/pizzeria",
		"https://github.com/AlexTLDR/interpretor",
		"https://github.com/AlexTLDR/blog.alextldr.com",
		"https://github.com/Stuttgart-Gophers/StuttgartGophers",
		"https://github.com/AlexTLDR/WebDev",
		"https://github.com/AlexTLDR/raffle",
		"https://github.com/AlexTLDR/imdb",
	}
	for _, link := range githubLinks {
		if doc.Find(`a[href="`+link+`"]`).Length() == 0 {
			t.Errorf("GitHub link %s not found", link)
		}
	}

	// Check that each project card has title and description
	cards.Each(func(i int, card *goquery.Selection) {
		title := card.Find(".card-title").Text()
		if title == "" {
			t.Errorf("project card %d title is empty", i)
		}

		description := card.Find("p").Text()
		if description == "" {
			t.Errorf("project card %d description is empty", i)
		}
	})
}
