package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"github.com/vivek-yadav/UserManagementService/routes/api/v0_1"
	"net/http"
)

func Setup(router *gin.RouterGroup) {
	r := router.Group("/v0_1")
	v0_1.Setup(r)

	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.User{Name: "vivek"})
	})
}
