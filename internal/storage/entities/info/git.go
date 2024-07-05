package info

import "net/url"

type GitEntity struct {
	Branch     string
	Commit     string
	Repository *url.URL
}
