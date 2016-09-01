package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/app"
	"net/http"
)

func GetList(c *gin.Context) {
	u := models.Apps{}
	r, er := modelApi.FetchAll("apps", c, u.DbFetchAll)
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
	u := models.App{}
	r, er := modelApi.FetchOne("apps", c, u.DbFetchOne)
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
	u := models.App{}
	r, er := modelApi.FetchById("apps", c, u.DbFetchOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}
