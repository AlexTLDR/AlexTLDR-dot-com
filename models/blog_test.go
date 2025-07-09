package models

import (
	"testing"
)

func TestBlogPost_Creation(t *testing.T) {
	post := BlogPost{
		Title:       "Test Title",
		Description: "Test Description",
		URL:         "https://example.com/post",
		Date:        "Jan 1, 2024",
		ReadTime:    "5 min read",
	}

	if post.Title != "Test Title" {
		t.Errorf("Expected Title to be 'Test Title', got %q", post.Title)
	}

	if post.Description != "Test Description" {
		t.Errorf("Expected Description to be 'Test Description', got %q", post.Description)
	}

	if post.URL != "https://example.com/post" {
		t.Errorf("Expected URL to be 'https://example.com/post', got %q", post.URL)
	}

	if post.Date != "Jan 1, 2024" {
		t.Errorf("Expected Date to be 'Jan 1, 2024', got %q", post.Date)
	}

	if post.ReadTime != "5 min read" {
		t.Errorf("Expected ReadTime to be '5 min read', got %q", post.ReadTime)
	}
}

func TestBlogPost_PartialFields(t *testing.T) {
	post := BlogPost{
		Title: "Only Title Set",
		URL:   "https://example.com",
	}

	if post.Title != "Only Title Set" {
		t.Errorf("Expected Title to be 'Only Title Set', got %q", post.Title)
	}

	if post.URL != "https://example.com" {
		t.Errorf("Expected URL to be 'https://example.com', got %q", post.URL)
	}

	// Other fields should be empty
	if post.Description != "" {
		t.Errorf("Expected empty Description, got %q", post.Description)
	}

	if post.Date != "" {
		t.Errorf("Expected empty Date, got %q", post.Date)
	}

	if post.ReadTime != "" {
		t.Errorf("Expected empty ReadTime, got %q", post.ReadTime)
	}
}

func TestBlogPost_Slice(t *testing.T) {
	posts := []BlogPost{
		{
			Title:    "First Post",
			URL:      "https://example.com/first",
			ReadTime: "3 min read",
		},
		{
			Title:    "Second Post",
			URL:      "https://example.com/second",
			ReadTime: "7 min read",
		},
	}

	if len(posts) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(posts))
	}

	if posts[0].Title != "First Post" {
		t.Errorf("Expected first post title to be 'First Post', got %q", posts[0].Title)
	}

	if posts[1].Title != "Second Post" {
		t.Errorf("Expected second post title to be 'Second Post', got %q", posts[1].Title)
	}
}
