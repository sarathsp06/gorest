package metrics

import (
	"time"
)

// HandleDelay can be used as defer for function to deal with its delay
func HandleDelay(handlers ...func(delay time.Duration)) func() {
	start := time.Now()
	return func() {
		delay := time.Since(start)
		for _, handler := range handlers {
			handler(delay)
		}
	}
}

// CaptureDelay returns a function that can be used with defer to calculate delay and sed to metrics
// Usage
//	function TestFunction() {}
//		defer CaptureDelay("request-id","TestFunction")()
//		do-something-for along time....
//	}
func CaptureDelay(funcName string) func() {
	return HandleDelay(func(delay time.Duration) { Delay(funcName, delay) })
}
