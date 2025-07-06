package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/AlexTLDR/AlexTLDR-dot-com/models"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func FetchLatestBlogPosts() ([]models.BlogPost, error) {
	// Fetch RSS feed from your Hugo blog
	resp, err := http.Get("https://blog.alextldr.com/index.xml")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RSS feed returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read RSS response: %w", err)
	}

	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, fmt.Errorf("failed to parse RSS: %w", err)
	}

	// Convert to BlogPost and get latest 3 (excluding about page)
	var posts []models.BlogPost
	count := 0

	for _, item := range rss.Channel.Items {
		// Skip about page and other non-post pages
		if strings.Contains(item.Link, "/about/") || strings.Contains(item.Link, "/tags/") || strings.Contains(item.Link, "/categories/") {
			continue
		}

		if count >= 3 {
			break
		}

		// Parse date
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		var dateStr string
		if err != nil {
			// Try alternative date format
			pubDate, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
			if err != nil {
				dateStr = item.PubDate // Use raw date if parsing fails
			} else {
				dateStr = pubDate.Format("Jan 2, 2006")
			}
		} else {
			dateStr = pubDate.Format("Jan 2, 2006")
		}

		// Estimate read time (very basic)
		readTime := estimateReadTime(item.Description)

		posts = append(posts, models.BlogPost{
			Title:       item.Title,
			Description: truncateDescription(item.Description, 150),
			URL:         item.Link,
			Date:        dateStr,
			ReadTime:    readTime,
		})

		count++
	}

	return posts, nil
}

func truncateDescription(desc string, maxLength int) string {
	// Strip all HTML tags completely
	for strings.Contains(desc, "<") && strings.Contains(desc, ">") {
		start := strings.Index(desc, "<")
		end := strings.Index(desc[start:], ">")
		if end != -1 {
			desc = desc[:start] + " " + desc[start+end+1:]
		} else {
			break
		}
	}

	// Decode HTML entities
	desc = strings.ReplaceAll(desc, "&lt;", "<")
	desc = strings.ReplaceAll(desc, "&gt;", ">")
	desc = strings.ReplaceAll(desc, "&amp;", "&")
	desc = strings.ReplaceAll(desc, "&quot;", "\"")
	desc = strings.ReplaceAll(desc, "&#39;", "'")
	desc = strings.ReplaceAll(desc, "&nbsp;", " ")
	desc = strings.ReplaceAll(desc, "&rsquo;", "'")
	desc = strings.ReplaceAll(desc, "&lsquo;", "'")
	desc = strings.ReplaceAll(desc, "&rdquo;", "\"")
	desc = strings.ReplaceAll(desc, "&ldquo;", "\"")

	// Clean up extra spaces
	desc = strings.Join(strings.Fields(desc), " ")

	if len(desc) <= maxLength {
		return desc
	}

	return desc[:maxLength-3] + "..."
}

func estimateReadTime(content string) string {
	// Strip HTML tags first for better word count
	cleanContent := content
	cleanContent = strings.ReplaceAll(cleanContent, "<", " <")
	cleanContent = strings.ReplaceAll(cleanContent, ">", "> ")

	// Remove all HTML tags
	for strings.Contains(cleanContent, "<") && strings.Contains(cleanContent, ">") {
		start := strings.Index(cleanContent, "<")
		end := strings.Index(cleanContent[start:], ">")
		if end != -1 {
			cleanContent = cleanContent[:start] + cleanContent[start+end+1:]
		} else {
			break
		}
	}

	// Clean up extra spaces and count words
	words := len(strings.Fields(cleanContent))

	// Average reading speed: ~200 words per minute
	minutes := words / 200
	if minutes < 1 {
		return "< 1 min read"
	}
	return fmt.Sprintf("%d min read", minutes)
}
