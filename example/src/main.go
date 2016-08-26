package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService"
	"net/http"
)

func main() {
	fmt.Println("Sample User Management Service Running....")
	service, err := ums.NewInstance()
	if err != nil {
		fmt.Printf("Error in creating instance of ums service : ( %v ) \n", err.Error())
		return
	}
	// Load Config
	configFilePath := "example/config/umsConfig.toml"
	if r, err := service.SetConfigFile(configFilePath); r == false && err != nil {
		configFilePath := "config/umsConfig.toml"
		r, err = service.SetConfigFile(configFilePath)
		if r == false && err != nil {
			fmt.Printf("Error in setting configurations from file of ums service : (%v)  : ( %v )", configFilePath, err.Error())
			return
		}
	}

	// Load Command Line Args
	if r, err := service.SetCmdArgs(); r == false && err != nil {
		fmt.Printf("Error in setting configurations from file of ums service : (%v)  : ( %v )", configFilePath, err.Error())
		return
	}

	router, err := service.GetRootRouter()
	if err != nil {
		fmt.Printf("Error in setting up Root Router :  ( %v )", err.Error())
		return
	}

	router.GET("/user/:name", HelloUser)

	service.Start(true)
}

func HelloUser(c *gin.Context) {
	name := c.Param("name")
	msg := "How have you been " + name + "?"
	//c.String(http.StatusOK, "Hello %s", name)
	c.HTML(http.StatusOK, "main/index.html", gin.H{
		"title":    "Welcome " + name,
		"userName": name,
		"msg":      msg,
	})
}
