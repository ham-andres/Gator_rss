package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/xml"
	"context"
	"io"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// fetching feed function
func fetchFeed(ctx context.Context, feedURL string) ( *RSSFeed, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot access the feedURL: %w", err)
	}
	req.Header.Set("User-Agent","gator")

	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response not read: %v", err)
	}

	rss := &RSSFeed{} // or we can do rss := new(RSSFeed) gives a pointer 

	err = xml.Unmarshal([]byte(body), rss)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal failed: %v", err)
	}
	
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for i, item := range rss.Channel.Item{
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rss.Channel.Item[i] = item
	}
	return rss, nil 
}
