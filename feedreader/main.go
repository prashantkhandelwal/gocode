package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type Feeds struct {
	FeedName string
	FeedURI  string
}

type FeedDetails struct {
	Author      string
	FeedAddress string
	Title       string
	Link        string
	Description string
	Content     string
	Categories  string
	Published   time.Time
}

func (f *Feeds) Download() *FeedDetails {
	fd := FeedDetails{}

	//fdl := make(map[string]fd)

	feedParser := gofeed.NewParser()
	feed, _ := feedParser.ParseURL(f.FeedURI)
	for _, m := range feed.Items {
		fd.Title = m.Title
		fd.Link = m.Link
		t, _ := time.Parse(time.RFC3339, m.Published)
		fd.Published = t

		fd.Author = m.Author.Name
		fd.Description = m.Description
		fd.Content = m.Content
		//fdl[f.FeedName] = fd
		//fmt.Printf("Title: %s | Link: %s\n", m.Title, m.Published)
	}

	return &fd

}

func main() {
	f := &Feeds{
		FeedName: "Midnightprogrammer",
		FeedURI:  "http://feeds.haacked.com/haacked",
	}

	d := f.Download()

	fmt.Println(d.Title)
	fmt.Println(d.Link)
	fmt.Println(d.Author)

}
