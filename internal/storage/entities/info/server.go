package info

import (
	"net"
	"net/url"
	"time"
)

type ServerEntity struct {
	OS        string
	Arch      string
	Hostname  string
	Url       *url.URL
	IP        net.IP
	Port      int
	Interface string
	Uptime    time.Duration
}
