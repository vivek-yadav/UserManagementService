package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetList(c *gin.Context) {
	u := models.Users{}
	u, _ = u.GetList()
	fmt.Printf("%#v", u)
	c.JSON(http.StatusOK, u)
}

func Get(c *gin.Context) {
	id := c.Query("id")
	var u models.User
	if bson.IsObjectIdHex(id) {
		u = models.User{Id: bson.ObjectIdHex(id)}
	} else {
		u = models.User{}
	}
	u, _ = u.GetById()
	c.JSON(http.StatusOK, u)
}
