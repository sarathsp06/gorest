package main

import (
	"fmt"
	"io"
	"log"
	"os"

	config "github.com/sarathsp06/gorest/config"
)

//ConfigureLog Configures logging with the configuration
func ConfigureLog() {
	config := config.GetConfig()
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	if file, err := GetLogFile("app"); err == nil {
		writers = append(writers, file)
	} else {
		log.Println("Failed opening file for app log")
	}

	output := io.MultiWriter(writers...)
	log.SetOutput(output)
	log.SetPrefix(fmt.Sprintf("[%s]", config.ProcessName))
	log.SetFlags(log.LUTC | log.Lshortfile)
}

// GetLogFile returns an io.Writer for log
// if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogFile(typ string) (*os.File, error) {
	logFile := fmt.Sprintf("/var/log/%s/%s.log", config.GetConfig().ProcessName, typ)
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[LOGGER]Error opening file", err.Error())
		return nil, err
	}
	return file, err
}
