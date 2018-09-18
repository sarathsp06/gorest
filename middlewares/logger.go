package middlewares

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	config "github.com/sarathsp06/gorest/config"
)

// GetLogFile returns an io.Writer for log
// if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogFile(typ string) (*os.File, error) {
	logFile := fmt.Sprintf("/var/log/%s/%s.log", config.Config.ProcessName, typ)
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[LOGGER]Error opening file", err.Error())
		return nil, err
	}
	return file, err
}

// FileLogger logs access logs to a file named `/var/log/{processname}/access.log`
// if failed to get access log file it logs to standard out
func FileLogger() echo.MiddlewareFunc {
	var writer io.Writer
	writer, err := GetLogFile("access")
	if err != nil {
		log.Println("Failed to open file for write,Err:", err.Error())
		log.Println("Writing to std out", err.Error())
		writer = os.Stdout
	}

	const format = `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out}}` + "\n"

	return middleware.LoggerWithConfig(
		middleware.LoggerConfig{Output: writer, Format: format},
	)
}
