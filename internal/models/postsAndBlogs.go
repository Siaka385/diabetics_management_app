package models

import "html/template"

// BlogPost model
type BlogPost struct {
	Title   string
	Author  string
	Date    string
	Content string
}

// Post model
type Post struct {
	ID      string
	Title   string
	Excerpt template.HTML
	Author  string
	Date    string
	Content template.HTML
	Image   string
	Abbrev  string
}
