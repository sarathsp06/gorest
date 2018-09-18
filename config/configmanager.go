package configmanager

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// configFile stores file name for the config
var configFile *string

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration() error {
	if configFile == nil {
		return errors.New("config not initialized")
	}
	config := new(ApplicatonConfig)
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	if err := GetDecoder(JSON)(file, config); err != nil {
		return err
	}
	if err := config.Validate(); err != nil {
		return err
	}
	Config = config
	return nil
}

// GetConfig returns a copy of init-ed application config instance
// if not already initialized it is initialized with "." as config path
func GetConfig() ApplicatonConfig {
	if Config != nil {
		return *Config
	}
	err := InitConfig(".")
	if err != nil {
		panic(err)
	}
	return *Config
}

//setup routine to reload config on sighup
func setupReload() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			log.Printf("Reloading configurations....\n")
			if configFile == nil {
				panic("Config file path not set!")
			}
			if err := LoadConfiguration(); err != nil {
				log.Printf("Error on reloading configurations, using old configuration : Error : %s\n", err.Error())
			}
		}
	}()
}

// InitConfig will initialize app config with config file name
func InitConfig(configDirectory string) error {
	if configDirectory == "" {
		configDirectory = "."
	}
	configFilepath := strings.Join([]string{configDirectory, "config.json"}, string(os.PathSeparator))
	configFile = &configFilepath
	setupReload()
	return LoadConfiguration()
}
