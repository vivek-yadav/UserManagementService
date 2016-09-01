package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/app"
	"net/http"
)

func Post(c *gin.Context) {
	a := models.App{}
	//c.Bind(&u)
	json.NewDecoder(c.Request.Body).Decode(&a)
	aa, er := modelApi.InsertOne("apps", a.DbInsertOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aa)
}

func Posts(c *gin.Context) {
	a := models.Apps{}
	json.NewDecoder(c.Request.Body).Decode(&a)
	aa, er := modelApi.InsertAll("apps", a.DbInsertAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aa)
}
