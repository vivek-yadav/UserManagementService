package ums

import (
	"errors"
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/vivek-yadav/UserManagementService/routes"
	"html/template"
	"net/http"
	"time"
)

type Service struct {
	Config     Config `json:"Config"`
	Engine     *gin.Engine
	RootRouter *gin.RouterGroup
}

// This returns a new Instance of User Management Service
func NewInstance() (*Service, error) {
	service := Service{}
	result, err := service.Config.setEnvArgs()
	if result == false && err != nil {
		return nil, errors.New("ERROR : Environment Variables were not proper ( " + err.Error() + " )")
	}
	return &service, nil
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
// If you send isblocking true then the system waits for the server to end first before return
// Else the call starts the server and returns, then it is up to you to hold the system to keep the
// service running.
func (this *Service) Start(isBlocking bool) {
	var paths gin.RoutesInfo
	this.RootRouter.GET("/_routes", func(c *gin.Context) {
		c.JSON(http.StatusOK, paths)
	})
	paths = this.Engine.Routes()
	this.Config.Show()
	if isBlocking {
		r := make(chan bool)
		go func(v chan bool) {
			serverPort := fmt.Sprintf(":%v", this.Config.WebServer.Port)
			manners.ListenAndServe(serverPort, this.Engine)
			v <- true
		}(r)
		<-r
	} else {
		go func() {
			serverPort := fmt.Sprintf(":%v", this.Config.WebServer.Port)
			manners.ListenAndServe(serverPort, this.Engine)
		}()
	}

}

func (this *Service) GetRootRouter() (*gin.RouterGroup, error) {
	if r, err := this.initService(); r == false && err != nil {
		return nil, err
	}
	if r, err := this.setupRootRouter(); r == false && err != nil {
		return nil, err
	}
	return this.RootRouter, nil
}

// This function sets up the root routing
func (this *Service) setupRootRouter() (bool, error) {
	this.RootRouter = this.Engine.Group("/")
	routes.Setup(this.RootRouter)
	return true, nil
}

func (this *Service) initService() (bool, error) {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(static.Serve("/", static.LocalFile(this.Config.FrontEnd.ViewsPath, true)))
	html, err := template.New("").Delims(this.Config.FrontEnd.TemplateDelimiterStart, this.Config.FrontEnd.TemplateDelimiterEnd).ParseGlob(this.Config.FrontEnd.TemplatesPath + "/**/*")
	if err != nil {
		return false, errors.New("ERROR : Failed to set Templates Path for Server : ( " + err.Error() + " )")
	}
	router.SetHTMLTemplate(html)

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	this.Engine = router

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
