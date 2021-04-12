package main

type Story map[string]Chapter

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

type Chapter struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []Option `json:"options,omitempty"`
}