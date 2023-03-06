package uptime

import "time"

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func GetUptime() string {

	uptime := time.Since(startTime)

	return uptime.String()
}
