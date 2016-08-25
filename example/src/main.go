package main

import (
	"fmt"
	"github.com/vivek-yadav/UserManagementService"
)

func main() {
	fmt.Println("Sample User Management Service Running....")
	service, err := ums.NewInstance()
	if err != nil {
		fmt.Printf("Error in creating instance of ums service : ( %v ) \n", err.Error())
		return
	}
	configFilePath := "example/config/umsConfig.toml"
	r, err := service.SetConfigFile(configFilePath)
	if r == false && err != nil {
		configFilePath := "config/umsConfig.toml"
		r, err = service.SetConfigFile(configFilePath)
		if r == false && err != nil {
			fmt.Printf("Error in setting configurations from file of ums service : (%v)  : ( %v )", configFilePath, err.Error())
			return
		}
	}
	r, err = service.SetCmdArgs()
	if r == false && err != nil {
		fmt.Printf("Error in setting configurations from file of ums service : (%v)  : ( %v )", configFilePath, err.Error())
		return
	}

	r, err = service.Start()
	if r == false && err != nil {
		fmt.Printf("Error in Running ums service : (%v)  : ( %v )", configFilePath, err.Error())
		return
	}
}
