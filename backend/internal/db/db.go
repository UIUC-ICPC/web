package db

// TODO: figure out how we're going to store and send links, i personally think we should
// store all their links in a single string seperated by commas, and then split
// in the frontend logic - AV Mishra

type Member struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	GradYear int    `json:"grad_year" db:"grad_year"`
	Alumni   bool   `json:"alumni" db:"alumni"`
	Links    string `json:"links" db:"links"`
	Verified bool   `json:"verified" db:"verified"`
}

// i imagine that image will be a string storing a link to an image, altho we prob need to
// iron those details out
type Post struct {
	ID        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Link      string `json:"link" db:"link"`
	Image     string `json:"image" db:"image"`
	Timestamp string `json:"timestamp" db:"created_at"`
	Author    string `json:"author" db:"author"`
}

// not touching auth until like next meeting!

// dummy data, replace as needed with actual db
var Members = []Member{
	{ID: 1, Name: "AV Mishra", GradYear: 2030, Alumni: false, Verified: true},
}

var Posts = []Post{
	{ID: 1, Title: "Welcome to SIGICPC", Author: "Exec Board"},
}
