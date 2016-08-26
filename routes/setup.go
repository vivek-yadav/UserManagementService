package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/routes/api"
)

func Setup(router *gin.RouterGroup) {

	r := router.Group("/api")
	api.Setup(r)
}
