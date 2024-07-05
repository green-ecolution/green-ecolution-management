package info

import "time"

type AppEntity struct {
	Version   string
	GoVersion string
	BuildTime time.Time
	Git       GitEntity
	Server    ServerEntity
}
