package components

import (
	"strings"
	"time"
)

var (
	StartTime time.Time
)

func SetStartTime() {
	StartTime = time.Now()
}

func GetUptime() string {

	diff := time.Since(StartTime)
	uptime := diff.String()
	uptime = strings.Split(uptime, ".")[0] + "s"

	return uptime
}
