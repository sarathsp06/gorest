package metrics

import (
	"fmt"
	"time"
)

// TODO: Replace the log.Println with proper metrics lib functions
// Add request id

// Delay sends metrics about delay on multiple
func Delay(funcName string, delay time.Duration) {
	fmt.Println("METRIC", funcName, delay)
}
