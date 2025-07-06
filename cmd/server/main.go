package main

import (
	"context"
	"log"
	"net/http"

	"github.com/AlexTLDR/AlexTLDR-dot-com/services"
	"github.com/AlexTLDR/AlexTLDR-dot-com/templates"
)

func main() {
	// Create HTTP server
	mux := http.NewServeMux()

	// Serve static files (for any additional assets)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle root route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set content type to HTML
		w.Header().Set("Content-Type", "text/html")

		// Render the index template with a name
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}

		component := templates.Index(name)
		err := component.Render(context.Background(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// API endpoint for latest blog posts
	mux.HandleFunc("/api/blog/latest", func(w http.ResponseWriter, r *http.Request) {
		// Set content type to HTML (since we're returning HTMX content)
		w.Header().Set("Content-Type", "text/html")

		// Fetch latest blog posts
		posts, err := services.FetchLatestBlogPosts()
		if err != nil {
			log.Printf("Error fetching blog posts: %v", err)
			// Return empty state on error
			posts = nil
		}

		// Convert models.BlogPost to templates.BlogPost
		var templatePosts []templates.BlogPost
		for _, post := range posts {
			templatePosts = append(templatePosts, templates.BlogPost{
				Title:       post.Title,
				Description: post.Description,
				URL:         post.URL,
				Date:        post.Date,
				ReadTime:    post.ReadTime,
			})
		}

		component := templates.BlogPosts(templatePosts)
		err = component.Render(context.Background(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
