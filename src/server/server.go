package ums

import "errors"

type Service struct {
	Config Config `json:"Config"`
}

// This returns a new Instance of User Management Service
func NewInstance() (Service, error) {
	service := Service{}
	result, err := service.Config.setEnvArgs()
	if result == false && err != nil {
		return nil, errors.New("ERROR : Environment Variables were not proper ( " + err.Error() + " )")
	}
	return service, nil
}

// This SetConfig function takes filePath of the config file
// and loads the User Management Service Instance with specified settings
// if some error occurs it throws error.
// if no file is sent in filePath param then default settings are loaded
func (this *Service) SetConfigFile(filePath string) (bool, error) {
	return this.Config.setFromFile(filePath)
}

// This sets configuration from command line arguments.
// Use this when you think your users might want to give command line arguments.
// Call this after SetConfig if you want it to have more priority.
func (this *Service) SetCmdArgs() (bool, error) {
	return this.Config.setFromCmdArgs()
}

// This function is used to start-up the service with given settings or default settings
func (this *Service) Start() (bool, error) {
	this.Config.Show()
	return true, nil
}

// This function is used to stop the service
func (this *Service) Stop() (bool, error) {
	return true, nil
}

// This function is used to Re-Start the service
func (this *Service) ReStart() (bool, error) {
	return true, nil
}
