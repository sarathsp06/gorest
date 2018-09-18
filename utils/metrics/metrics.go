package metrics

import (
	"fmt"
	"time"
)

// TODO: Replace the fmt.Println with proper metrics lib functions

// Delay sends metrics about delay on multiple
func Delay(reqID string, funcName string, delay time.Duration) {
	fmt.Println("METRIC", reqID, funcName, delay)
}
