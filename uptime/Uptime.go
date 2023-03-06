//Somewhat inspired by: https://stackoverflow.com/questions/37992660/golang-retrieve-application-uptime

package uptime

import "time"

// Global variable, the start-time of the service
var startTime time.Time

// Initialize the start-time variable
func Init() {
	startTime = time.Now()
}

// Return the uptime in seconds as a string
func GetUptime() string {

	uptime := time.Since(startTime)

	return uptime.String()
}
