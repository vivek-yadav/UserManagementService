package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"net/http"
)

func GetList(c *gin.Context) {
	u := models.Users{}
	r, er := modelApi.GetList("users", c, u)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func Get(c *gin.Context) {
	u := models.User{}
	r, er := modelApi.Get("users", c, u)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func GetById(c *gin.Context) {
	u := models.User{}
	r, er := modelApi.GetById("users", c, u)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}
