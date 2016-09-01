package v0_1

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/routes/api/v0_1/app"
	"github.com/vivek-yadav/UserManagementService/routes/api/v0_1/user"
)

func Setup(router *gin.RouterGroup) {
	// Users
	router.GET("/users", user.GetList)
	router.GET("/user", user.Get)
	router.GET("/user/:id", user.GetById)

	router.POST("/users", user.PostUsers)
	router.POST("/user", user.PostUser)

	router.POST("/user/login", user.Login)

	router.PUT("/user/update", user.UpdateOne)
	router.PUT("/users/update", user.UpdateAll)

	router.PUT("/user", user.ReplaceOneById)
	//router.PUT("/user/search", user.ReplaceOne)
	router.PUT("/users", user.ReplaceAll)

	// Apps
	router.POST("/apps", app.Posts)
	router.POST("/app", app.Post)

	router.GET("/apps", app.GetList)
	router.GET("/app", app.Get)
	router.GET("/app/:id", app.GetById)

}
