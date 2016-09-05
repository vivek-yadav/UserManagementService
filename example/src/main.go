package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService"
	"github.com/vivek-yadav/UserManagementService/microServices/authService"
	"github.com/vivek-yadav/UserManagementService/microServices/loginService"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Sample User Management Service Running....")
	service, err := ums.GetInstance()
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
	var stop chan bool
	go func() {
		service.Start(false)
		stop <- true
	}()
	go func() {
		loginService.StartService(":7001")
		time.Sleep(5 * time.Millisecond)
		authService.StartService(":7002")
		time.Sleep(50 * time.Millisecond)

		stop <- true
	}()
	for i := 0; i < 2; i++ {
		<-stop
	}
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
