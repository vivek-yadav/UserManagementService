package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"net/http"
)

func DeleteOne(c *gin.Context) {
	er := modelApi.DeleteOne("apps", c)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func DeleteAll(c *gin.Context) {
	er := modelApi.DeleteAll("apps", c)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
