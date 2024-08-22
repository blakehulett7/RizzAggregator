package main

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Dc      string   `xml:"dc,attr"`
	Content string   `xml:"content,attr"`
	Channel struct {
		Text string `xml:",chardata"`
		Link struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Copyright   string `xml:"copyright"`
		PubDate     string `xml:"pubDate"`
		Generator   string `xml:"generator"`
		Docs        string `xml:"docs"`
		Image       struct {
			Text  string `xml:",chardata"`
			URL   string `xml:"url"`
			Link  string `xml:"link"`
			Title string `xml:"title"`
		} `xml:"image"`
		Items []Item `xml:"item"`
	} `xml:"channel"`
	Entries []Entry `xml:"entry"`
}

type Item struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Creator     string `xml:"creator"`
	Description string `xml:"description"`
	Encoded     string `xml:"encoded"`
	Guid        struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	PubDate   string `xml:"pubDate"`
	Enclosure struct {
		Text string `xml:",chardata"`
		URL  string `xml:"url,attr"`
	} `xml:"enclosure"`
}

type Entry struct {
	Text      string `xml:",chardata"`
	ID        string `xml:"id"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Link      struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	URL     string `xml:"url"`
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Content string `xml:"content"`
	Author  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"author"`
}

type post interface {
	GetTitle() string
}

func (item Item) GetTitle() string {
	return item.Title
}

func (entry Entry) GetTitle() string {
	return entry.Title
}
