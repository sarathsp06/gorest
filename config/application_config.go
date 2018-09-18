package configmanager

import "errors"

// ApplicatonConfig is the configuration for the ApplicatonConfig
// it does not have interface specific configuration
type ApplicatonConfig struct {
	ProcessName string `json:"process_name,omitempty"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Google      Google `json:"google"`
	Mongo       Mongo  `json:"mongo"`
}

// Validate validates the application config and fails if mandatory parameters are missing
func (cnf ApplicatonConfig) Validate() error {
	if cnf.ProcessName == "" || cnf.Host == "" || cnf.Port == 0 {
		return errors.New("invalid Config:  ProcessName,Port and Host  are mandatory")
	}
	if err := cnf.Google.Validate(); err != nil {
		return errors.New("invalid Config [google] " + err.Error())
	}
	if err := cnf.Mongo.Validate(); err != nil {
		return errors.New("invalid Config [Mongo] " + err.Error())
	}
	return nil
}

// Config stores the configuration
var Config *ApplicatonConfig
