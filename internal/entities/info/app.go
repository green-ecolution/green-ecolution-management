package info

import "time"

type App struct {
	Version   string
	GoVersion string
	BuildTime time.Time
	Git       Git
	Server    Server
}
