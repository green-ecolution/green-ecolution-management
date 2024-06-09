package info

import "net/url"

type Git struct {
	Branch     string
	Commit     string
	Repository *url.URL
}
