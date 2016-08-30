package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"net/http"
)

func PostUser(c *gin.Context) {
	u := models.User{}
	//c.Bind(&u)
	json.NewDecoder(c.Request.Body).Decode(&u)
	u, _ = u.CreateUser()
	fmt.Printf("%#v", u)
	c.JSON(http.StatusOK, u)
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
