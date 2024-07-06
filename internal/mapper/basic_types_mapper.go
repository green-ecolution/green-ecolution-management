package mapper

import (
	"net"
	"net/url"
	"time"
)

func TimeToTime(t time.Time) time.Time {
	return t
}

func UrlToUrl(u *url.URL) *url.URL {
  return u
}

func TimeDurationToTimeDuration(t time.Duration) time.Duration {
  return t
}

func StringToTime(s string) time.Time {
  t, _ := time.Parse(time.RFC3339, s)
  return t
}

func StringToUrl(s string) *url.URL {
  u, _ := url.Parse(s)
  return u
}

func StringToNetIP(s string) net.IP {
  ip := net.ParseIP(s)
  return ip
}

func NetIPToString(ip net.IP) string {
  return ip.String()
}

func StringToDuration(s string) time.Duration {
  d, _ := time.ParseDuration(s)
  return d
}

func TimeToString(t time.Time) string {
  return t.Format(time.RFC3339)
}

func NetUrlToString(u *url.URL) string {
  return u.String()
}

func TimeDurationToString(t time.Duration) string {
  return t.String()
}
