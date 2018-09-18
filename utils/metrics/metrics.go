package metrics

import (
	"log"
	"time"
)

// TODO: Replace the log.Println with proper metrics lib functions

// Delay sends metrics about delay on multiple
func Delay(reqID string, funcName string, delay time.Duration) {
	log.Println("METRIC", reqID, funcName, delay)
}
