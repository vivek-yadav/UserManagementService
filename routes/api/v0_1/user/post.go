package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"net/http"
)

func PostOne(c *gin.Context) {
	u := models.User{}
	//c.Bind(&u)
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.InsertOne("users", u.DbInsertOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func PostAll(c *gin.Context) {
	u := models.Users{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.InsertAll("users", u.DbInsertAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uu)
}

func Login(c *gin.Context) {
	var er error
	u := models.User{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	var result bool
	er = u.IsLogin(&u, &result)
	if er != nil {
		er = errors.New("Failed Login, Please retry with correct username and password.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, u)
}
