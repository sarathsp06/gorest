package main

import (
	"fmt"
	"time"
)

// about contains build details and uptime details about the service
// Version,MinVersion and BuildTime are set on build
type about struct {
	Version    string
	MinVersion string
	BuildTime  string
	StartedAt  time.Time
	Uptime     string
}

// serverDetails contains the build details of the current service
var serverDetails about

// Version details to be filled un build itself
var (
	Version    string
	MinVersion string
	BuildTime  string
)

func init() {
	serverDetails = about{Version: Version, MinVersion: MinVersion, BuildTime: BuildTime, StartedAt: time.Now()}
}

// Heartbeat returns details of the instance running
func Heartbeat() interface{} {
	uptime := time.Since(serverDetails.StartedAt)
	serverDetails.Uptime = fmt.Sprintf("%d days %s", uptime/(time.Hour*24), time.Time{}.Add(uptime).Format("15:04:05"))
	return serverDetails
}
