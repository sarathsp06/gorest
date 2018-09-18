package configmanager

import "errors"

//Google contains configuration for accessing google services
type Google struct {
	AppName string `json:"app_name"`
	Key     string `json:"key"`
}

// Validate validates the application config and fails if mandatory parameters are missing
func (cnf Google) Validate() error {
	if cnf.Key == "" {
		return errors.New("Key is mandatory for google configuration")
	}
	return nil
}
