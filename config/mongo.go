package configmanager

import "errors"

// Mongo configuration
type Mongo struct {
	Host              string
	Port              int
	DB                string
	ConnectionTimeout int `json:"connection_timeout"`
}

// Validate validates the application config and fails if mandatory parameters are missing
func (cnf Mongo) Validate() error {
	if cnf.Host == "" || cnf.Port == 0 {
		return errors.New("Port and Host  are mandatory")
	}
	return nil
}
