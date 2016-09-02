package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/app"
	"net/http"
)

func UpdateOneById(c *gin.Context) {
	update := map[string]interface{}{}
	json.NewDecoder(c.Request.Body).Decode(&update)
	uu, er := modelApi.UpdateOneById("apps", c, update)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func UpdateOne(c *gin.Context) {
	u := models.App{}
	update := map[string]interface{}{}
	json.NewDecoder(c.Request.Body).Decode(&update)
	uu, er := modelApi.Update("apps", c, update, u.DbUpdateOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func UpdateAll(c *gin.Context) {
	u := models.Apps{}
	update := map[string]interface{}{}
	json.NewDecoder(c.Request.Body).Decode(&update)
	uu, er := modelApi.Update("apps", c, update, u.DbUpdateAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func ReplaceAll(c *gin.Context) {
	u := models.Apps{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.ReplaceAll("apps", u.DbReplaceAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func ReplaceOne(c *gin.Context) {
	u := models.App{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.ReplaceOne("apps", u.DbReplaceOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}
