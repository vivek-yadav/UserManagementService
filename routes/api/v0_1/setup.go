package v0_1

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/routes/api/v0_1/user"
)

func Setup(router *gin.RouterGroup) {
	router.GET("/users", user.GetList)
	router.GET("/user", user.Get)
}
