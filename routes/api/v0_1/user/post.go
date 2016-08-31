package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/api"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"net/http"
)

func PostUser(c *gin.Context) {
	u := models.User{}
	//c.Bind(&u)
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.Create("users", c, u)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	fmt.Printf("%#v", uu)
	c.JSON(http.StatusOK, uu)
}

func PostUsers(c *gin.Context) {
	u := models.Users{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	uu, er := modelApi.CreateList("users", c, u)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	fmt.Printf("%#v", uu)
	c.JSON(http.StatusOK, uu)
}

func Login(c *gin.Context) {
	var er error
	u := models.User{}
	json.NewDecoder(c.Request.Body).Decode(&u)
	fmt.Println(u)
	if u.Username != "" && u.Password != "" {
		var e error
		u, e = u.FindLoginUserWithUsername()
		if e != nil {
			er = errors.New("Failed Login, Please retry with correct username and password. (\n\t" + e.Error() + "\n)")

		}
	} else if u.Email != "" && u.Password != "" {
		var e error
		u, e = u.FindLoginUserWithEmail()
		if e != nil {
			er = errors.New("Failed Login, Please retry with correct email and password. (\n\t" + e.Error() + "\n)")

		}
	} else {
		er = errors.New("Failed Login, Please retry with correct email and password.")
	}
	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": er.Error(),
		})
		return
	}
	fmt.Println(u)
	c.JSON(http.StatusOK, u)
}
