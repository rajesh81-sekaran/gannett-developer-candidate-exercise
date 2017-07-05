package common

import (
	"html/template"
	"sync"
	"time"
)

// ProfileIDStruct structure to fetch the profile ID
type ProfileIDStruct struct {
	ProfileID int `json:profileId`
}

// AllArticles structure array of Articles
type AllArticles struct {
	Articles  []Article `json:"articles"`
	ProfileID int       `json:"profileID"`
	Theme     string    `json:"theme"`
}

// Article individual article
type Article struct {
	Href    string `json:"href"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
}

// DefaultProfileURL to get the profile ID
const DefaultProfileURL = "https://peaceful-springs-7920.herokuapp.com/profile/"

// DefaultContentURL to get the Content ID
const DefaultContentURL = "https://peaceful-springs-7920.herokuapp.com/content/"

// ProfileURL profile url that will be used in the code
var ProfileURL string

// ContentURL content url that will be used in the code
var ContentURL string

// DefaultPort port on which server will listen
const DefaultPort = "9090"

// Port port that will be used in code
var Port string

// SingleUser user details
type SingleUser struct {
	profileID      int
	timeRegistered time.Time
	timeExpiration time.Time
}

// AllUsers map of registered users
var AllUsers map[string]SingleUser

// UsersLock lock for synchronizing user registration
var UsersLock sync.RWMutex

// Template function
var fm = template.FuncMap{
	"fontColor": fontColor,
}

// Function to set the font color
func fontColor(s string) string {
	if s == "rare" {
		return "#DC143C"
	} else if s == "well" {
		return "#8B4513"
	} else {
		return "black"
	}
}
